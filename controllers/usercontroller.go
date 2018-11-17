package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alanwgt/apateapi/crypto"

	"github.com/golang/protobuf/proto"

	db "github.com/alanwgt/apateapi/database"
	"github.com/alanwgt/apateapi/messages"
	"github.com/alanwgt/apateapi/models"
	"github.com/alanwgt/apateapi/protos"
)

// CreateAccount creates an user account if all the requirements are satisfied
// The username MUST be unique and the fcm_id cannot be a duplicate
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	decoded, _ := ioutil.ReadAll(r.Body)

	ar := &protos.AccountSignUp{}
	if err := proto.Unmarshal(decoded, ar); err != nil {
		// wrong proto! discard request
		log.Println(err)
		messages.ErrorWithMessage(w, http.StatusBadRequest, "The wrong proto was used.")
		return
	}

	// check if the username is available
	u := models.User{}
	c := db.GetOpenConnection()
	if !c.First(&u, &models.User{Username: ar.Username}).RecordNotFound() {
		// found a record, send error
		log.Println("User not created! Duplicated entry for username")
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
	decoded, _ := ioutil.ReadAll(r.Body)
	dr := &protos.DeviceRequest{}

	if err := proto.Unmarshal(decoded, dr); err != nil {
		// wrong proto! discard the request
		log.Println(err)
		messages.ErrorWithMessage(w, http.StatusBadRequest, "The wrong proto was used.")
		return
	}

	_, err := crypto.OpenUserBox(dr)

	if err != nil {
		log.Println("Couldn't authenticate the user!")
		log.Println(err)
		messages.ErrorWithMessage(w, http.StatusForbidden, err.Error())
		return
	}

	messages.RequestOK(w, "handshake:"+dr.Username)
}
