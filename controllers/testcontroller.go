package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("Gorilla!\n"))
}
