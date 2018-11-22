package messages

import (
	"log"
	"net/http"

	"github.com/alanwgt/apateapi/protos"
)

// CryptoError sends a default message when a user's box can not be opened
func CryptoError(w http.ResponseWriter, dr *protos.DeviceRequest, err error) {
	log.Println("Couldn't authenticate the user:", dr.Username)
	log.Println(dr)
	log.Println(err)
	ErrorWithMessage(w, http.StatusForbidden, err.Error())
}

// WrongProto sends a default message when the wrong proto is used in a device request
func WrongProto(w http.ResponseWriter, err error) {
	// wrong proto! discard the request
	log.Println(err)
	ErrorWithMessage(w, http.StatusBadRequest, "The wrong proto was used.")
}

// ServerError sends a default message with a 500 status code
func ServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	ErrorWithMessage(w, http.StatusInternalServerError, err.Error())
}
