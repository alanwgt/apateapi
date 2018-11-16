package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/alanwgt/apateapi/database"
	"github.com/alanwgt/apateapi/messages"
	"github.com/alanwgt/apateapi/models"
	"github.com/alanwgt/apateapi/protos"
)

// Creates an account // TODO
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

	nu := models.User{
		Username: ar.Username,
		FcmToken: ar.FcmToken,
		PubKey:   ar.PubK,
	}

	database.Create(nu)
	fmt.Fprintf(w, "Ok!")
}
