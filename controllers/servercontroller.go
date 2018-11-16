package controllers

import (
	"fmt"
	"net/http"

	"github.com/alanwgt/apateapi/crypto"
)

func GetServerPubK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, crypto.GetB64PubK())
}
