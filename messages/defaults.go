package messages

import (
	"net/http"
)

// Error sends a default error proto with a default message and specified status code
func Error(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

// ErrorWithMessage sends a default error proto with the specified message and status code
func ErrorWithMessage(w http.ResponseWriter, statusCode int, m string) {
	w.WriteHeader(statusCode)
}
