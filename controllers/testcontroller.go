package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func TestMe(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	fmt.Println(vals)
	w.Write([]byte("Gorilla!\n"))
}
