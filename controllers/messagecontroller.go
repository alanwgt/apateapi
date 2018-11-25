package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/alanwgt/apateapi/cache"

	"github.com/golang/protobuf/proto"

	"github.com/jinzhu/gorm"

	"github.com/alanwgt/apateapi/protos"

	"github.com/alanwgt/apateapi/models"

	db "github.com/alanwgt/apateapi/database"
	"github.com/alanwgt/apateapi/messages"
	"github.com/gorilla/mux"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	m, uc, err := OpenRequestBox(w, r)

	if err != nil {
		return
	}

	un, ok := mux.Vars(r)["user"]

	if !ok {
		log.Println("Missing 'user' in request")
		messages.BadRequest(w)
		return
	}

	ud, err := cache.GetUser(un)

	if err != nil {
		log.Println(err)
		messages.BadRequest(w)
		return
	}

	dm := &protos.MessageBody{}
	// b64dm, err := base64.StdEncoding.En

	// if err != nil {
	// 	log.Println(err)
	// 	messages.BadRequest(w)
	// 	return
	// }

	if err := proto.Unmarshal(m, dm); err != nil {
		log.Println(err)
		messages.BadRequest(w)
		return
	}

	c := db.GetOpenConnection()

	nm := &models.Message{
		UserID:      uc.Model.ID,
		RecipientID: ud.Model.ID,
	}

	c.Create(nm)

	nmc := &models.MessageContent{
		Body:      dm.Body,
		MessageID: nm.ID,
		Nonce:     dm.Nonce,
		Type:      protos.MessageBody_Type_value[dm.Type.String()],
	}

	c.Create(nmc)

	// md := map[string]string{
	// 	"sender":    ud.Model.Username,
	// 	"timestamp": strconv.FormatInt(nm.CreatedAt.Unix(), 10),
	// 	"id":        strconv.FormatInt(nm.ID, 10),
	// }

	// FIXME: CHANGE THE FCM TOKEN
	// services.SendFCMMessage(
	// 	uc.Model.FcmToken,
	// 	"New message",
	// 	"from: "+uc.Model.Username,
	// 	md,
	// )
	messages.RequestOK(w, "message sent")
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {

}

// LoadMessages returns a proto MessageBody array and deletes them from database
func LoadMessages(w http.ResponseWriter, r *http.Request) {
	_, _, err := OpenRequestBox(w, r)

	if err != nil {
		return
	}

	mids, ok := mux.Vars(r)["id"]

	if !ok {
		log.Println("Missing 'id' in request")
		messages.BadRequest(w)
		return
	}

	midsarr := strings.Split(mids, ",")
	c := db.GetOpenConnection()
	res := &protos.MessagesContainer{}
	var msgs []*protos.MessageBody

	for _, mid := range midsarr {
		midc, err := strconv.ParseInt(mid, 10, 64)
		if err != nil {
			messages.ServerError(w, err)
			return
		}
		lm, err := loadMessage(c, midc)
		if err != nil {
			messages.ServerError(w, err)
			return
		}
		msgs = append(msgs, lm)
	}

	res.Messages = msgs
	resB, err := proto.Marshal(res)

	if err != nil {
		messages.ServerError(w, err)
		return
	}

	messages.CustomProto(w, resB)
}

func loadMessage(c *gorm.DB, id int64) (*protos.MessageBody, error) {
	m := &models.Message{}

	if c.Preload("Body").First(m, &models.Message{ID: id}).RecordNotFound() {
		log.Println("No message found for id:", id)
		return nil, errors.New("Record not found for message id: " + string(id))
	}

	var pType protos.MessageBody_Type

	if protos.MessageBody_Type_name[m.Body.Type] == "Text" {
		pType = protos.MessageBody_Text
	} else {
		pType = protos.MessageBody_Image
	}

	mp := &protos.MessageBody{
		MessageId: m.ID,
		Body:      m.Body.Body,
		Nonce:     m.Body.Nonce,
		Type:      pType,
	}

	// c.Model(&m).Update("opened_at", time.Now())
	// c.Delete(&m.Body)

	return mp, nil
}
