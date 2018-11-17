package controllers

import (
	"fmt"
	"net/http"

	"github.com/alanwgt/apateapi/crypto"
)

// GetServerPubK sends the server's current public key encoded in base64
func GetServerPubK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, crypto.GetServerB64PubK())
}
