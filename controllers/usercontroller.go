package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	db "github.com/alanwgt/apateapi/database"
	"github.com/alanwgt/apateapi/messages"
	"github.com/alanwgt/apateapi/models"
	"github.com/alanwgt/apateapi/protos"
)

// CreateAccount creates an user account if all the requirements are satisfied
// The username MUST be unique and the fcm_id cannot be a duplicate
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	decoded, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		messages.ErrorWithMessage(w, http.StatusBadRequest, "An error occurred while reading bytes from body.")
		return
	}

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
