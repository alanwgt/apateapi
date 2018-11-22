package controllers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alanwgt/apateapi/cache"

	"github.com/gorilla/mux"

	"github.com/alanwgt/apateapi/crypto"

	"github.com/golang/protobuf/proto"

	db "github.com/alanwgt/apateapi/database"
	"github.com/alanwgt/apateapi/messages"
	"github.com/alanwgt/apateapi/models"
	"github.com/alanwgt/apateapi/protos"
	"github.com/alanwgt/apateapi/protoutil"
)

// CreateAccount creates an user account if all the requirements are satisfied
// The username MUST be unique and the fcm_id cannot be a duplicate
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	decoded, _ := ioutil.ReadAll(r.Body)

	ar := &protos.AccountSignUp{}
	if err := proto.Unmarshal(decoded, ar); err != nil {
		messages.WrongProto(w, err)
		return
	}

	// check if the username is available
	u := models.User{}
	c := db.GetOpenConnection()
	if !c.First(&u, &models.User{Username: ar.Username}).RecordNotFound() {
		// found a record, send error
		log.Println("User not created! Duplicated entry for username:", u.Username)
		messages.ErrorWithMessage(w, http.StatusConflict, fmt.Sprintf("The username '%s' is already taken!", ar.Username))
		return
	}

	// check if there is a duplicated entry for the fcm_id
	if !c.First(&u, &models.User{FcmToken: ar.FcmToken}).RecordNotFound() {
		// the device already has an account registered to it!
		// this cannot happen to the end user
		log.Println("User not created! The device already has an associated account")
		messages.ErrorWithMessage(w, http.StatusConflict, "This device already has an account registered to it!")
		return
	}

	u = models.User{
		Username: ar.Username,
		FcmToken: ar.FcmToken,
		PubKey:   ar.PubK,
	}

	db.Create(&u)
	log.Printf("User '%s' created!\n", u.Username)
	messages.RequestOK(w, "Created!")
}

// Handshake exchanges an encrypted message to ensure that the user is authenticated
func Handshake(w http.ResponseWriter, r *http.Request) {
	_, uc, err := openRequestBox(w, r)

	if err != nil {
		return
	}

	// TODO: send all the data that the device will use (messages, contacts,...)
	c := db.GetOpenConnection()
	var frs []models.FriendRequest
	var cs []models.User
	var rms, sms []models.Message
	var utf []int64

	// get all the active contacts
	// TODO: maybe we need to send the users that this user has blocked
	rows, err := c.Raw(`
SELECT fr.user_id as uid, fr.request_to as reqto
FROM apate."user" u
		INNER JOIN apate.friend_request fr ON u.id IN (fr.user_id, fr.request_to)
WHERE fr.deleted_at IS NULL
	AND u.id = ?
	AND NOT EXISTS(SELECT NULL
					FROM apate.blocked b
					WHERE fr.request_to IN (b.user_id, b.blocked_id)
					AND fr.user_id IN (b.user_id, b.blocked_id));
	`, uc.Model.ID).Rows()

	if err != nil {
		messages.ServerError(w, err)
		return
	}

	for rows.Next() {
		var uid, reqto *int64
		var fu int64
		rows.Scan(&uid, &reqto)
		// the user initiated the friend request approval
		if *uid == uc.Model.ID {
			fu = *reqto
		} else {
			fu = *uid
		}
		utf = append(utf, fu)
	}

	rows.Close()

	for _, ucID := range utf {
		cm, err := cache.GetUserByID(ucID)
		if err != nil {
			messages.ServerError(w, err)
			return
		}
		cs = append(cs, *cm.Model)
	}

	c.
		Preload("Requester").
		Preload("RequestedTo").
		Where(&models.FriendRequest{
			RequestTo:  uc.Model.ID,
			DeletedAt:  nil,
			AcceptedAt: nil,
		}).Find(&frs)

	c.
		Preload("Sender").
		Preload("Receiver").
		Where(&models.Message{
			RecipientID: uc.Model.ID,
			OpenedAt:    nil,
			DeletedAt:   nil,
		}).Find(&rms)

	c.
		Preload("Sender").
		Preload("Receiver").
		Where(&models.Message{
			UserID:    uc.Model.ID,
			OpenedAt:  nil,
			DeletedAt: nil,
		}).Find(&sms)

	ac := &protos.AccountHandshake{
		Contacts:         protoutil.UserModelToProto(cs...),
		SentMessages:     protoutil.MessageModelToProto(sms...),
		UnopenedMessaged: protoutil.MessageModelToProto(rms...),
		FriendRequests:   protoutil.FriendRequestToProto(frs...),
	}

	out, err := proto.Marshal(ac)

	if err != nil {
		messages.ServerError(w, err)
		return
	}

	messages.CustomProto(w, out)
	// messages.RequestOK(w, "handshake:"+uc.Model.Username)
}

// QueryUsers will send all the users that matches the query
func QueryUsers(w http.ResponseWriter, r *http.Request) {
	decoded, _ := ioutil.ReadAll(r.Body)
	dr := &protos.DeviceRequest{}

	if err := proto.Unmarshal(decoded, dr); err != nil {
		messages.WrongProto(w, err)
		return
	}

	vars := mux.Vars(r)
	qu, ok := vars["username"]

	if !ok || len(qu) < 4 {
		log.Println("expecting username to be bigger than 3, received:", qu)
		messages.ErrorWithMessage(w, http.StatusBadRequest, "Expecting username to be bigger than 3 chars")
		return
	}

	c := db.GetOpenConnection()
	var users []models.User

	c.Where("username ILIKE ?", qu+"%").Select("id, username, pub_key, created_at").Limit(5).Order("username desc").Find(&users)

	// messages.RequestOK(w, "users...")
	messages.RawJSON(w, users)
}

// DeleteAccount expects a DELETE request on /user, and if the credentials are valid,
// the user is deleted from the database
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	d, ok := r.URL.Query()["d"]

	if !ok || len(d) != 1 {
		// missing data in request, or we got too much data (?)
		log.Println("Missing ?d in request")
		messages.ErrorWithMessage(w, http.StatusBadRequest, "Missing data.")
		return
	}

	log.Println(d)
	rd, _ := base64.StdEncoding.DecodeString(d[0])

	dr := &protos.DeviceRequest{}
	if err := proto.Unmarshal(rd, dr); err != nil {
		messages.WrongProto(w, err)
		return
	}

	if _, _, err := crypto.OpenUserBox(dr); err != nil {
		messages.CryptoError(w, dr, err)
		return
	}

	c := db.GetOpenConnection()
	u := &models.User{}

	c.First(&u, "username = ?", dr.Username)

	if u.Username == "" {
		log.Printf("User '%s' not found in DB", dr.Username)
		messages.ErrorWithMessage(w, http.StatusForbidden, "Forbidden.")
		return
	}

	// Unscoped will grant that the entry is hard deleted.
	c.Unscoped().Delete(&u)
	cache.RemoveUser(u)

	log.Printf("User '%s' successfully deleted.\n", u.Username)
	// delete from database and remove from cache!
	messages.RequestOK(w, "deleted")
}

// AddContact will create a friend request if everything is satisfied
func AddContact(w http.ResponseWriter, r *http.Request) {
	_, u, e := openRequestBox(w, r)

	if e != nil {
		return
	}

	un, ok := mux.Vars(r)["username"]

	if !ok {
		log.Println("Missing 'username' from request")
		messages.BadRequest(w)
		return
	}

	uc, err := cache.GetUser(un)

	if err != nil {
		log.Println(err)
		messages.BadRequest(w)
		return
	}

	c := db.GetOpenConnection()
	nc := &models.FriendRequest{}

	nc.UserID = u.Model.ID
	nc.RequestTo = uc.Model.ID

	log.Printf("User '%s' requested contact approval to '%s'.\n", u.Model.Username, uc.Model.Username)

	c.Create(&nc)
	messages.RequestOK(w, "requested")
}

func RemoveContact(w http.ResponseWriter, r *http.Request) {

}

func AcceptContact(w http.ResponseWriter, r *http.Request) {

}

// openRequestBox automatically opens the user's box, returning the
// plain content with the user model from cache. If an error occurs
// during the process, an error will be sent
func openRequestBox(w http.ResponseWriter, r *http.Request) (un string, u *cache.UserCache, e error) {
	var buf []byte

	if r.Method == "GET" || r.Method == "DELETE" {
		d, ok := r.URL.Query()["d"]
		if (!ok) || len(d) != 1 {
			if len(d) > 1 {
				e = errors.New("Received too much data in request, expecting one")
			} else {
				e = errors.New("Missing data 'd' in querystring")
			}
			messages.ErrorWithMessage(w, http.StatusBadRequest, e.Error())
			return
		}
		buf, _ = base64.StdEncoding.DecodeString(d[0])
	} else if r.Method == "POST" {
		buf, _ = ioutil.ReadAll(r.Body)
	}

	dr := &protos.DeviceRequest{}

	if e = proto.Unmarshal(buf, dr); e != nil {
		messages.WrongProto(w, e)
		return
	}

	if un, u, e = crypto.OpenUserBox(dr); e != nil {
		messages.CryptoError(w, dr, e)
	}

	return
}
