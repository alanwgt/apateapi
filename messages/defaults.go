package messages

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/alanwgt/apateapi/protos"
)

// Error sends a default error proto with a default message and specified status code
func Error(w http.ResponseWriter, statusCode int) {
	ErrorWithMessage(w, statusCode, "Whoops! It seems that the developer doesn't know what he's doing :(")
}

// ErrorWithMessage sends a default error proto with the specified message and status code
func ErrorWithMessage(w http.ResponseWriter, statusCode int, m string) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, buildB64Proto(
		m,
		statusCode,
		protos.ServerResponse_ERROR,
	))
}

// RequestOK simply sends a 200 status code with a message to the user
func RequestOK(w http.ResponseWriter, m string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, buildB64Proto(
		m,
		http.StatusOK,
		protos.ServerResponse_Ok,
	))
}

// BadRequest sends a default bad request response
func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

// RawJSON sends a JSON to the request
func RawJSON(w http.ResponseWriter, m interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, err := json.Marshal(m)

	if err != nil {
		log.Panic(err)
	}

	w.Write(d)
}

// CustomProto sends an OK response with a binary protobuf
// encoded in base64
func CustomProto(w http.ResponseWriter, p []byte) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, bToB64(p))
}

// returns a b64 server response proto message safe to transport via HTTP
func buildB64Proto(m string, statusCode int, status protos.ServerResponse_Status) string {
	p := &protos.ServerResponse{
		Message:    m,
		StatusCode: int32(statusCode),
		Status:     status,
	}

	out, err := proto.Marshal(p)

	if err != nil {
		log.Println(err)
		// TODO: return an appropriate error message
		return ""
	}

	return bToB64(out)
}

// Encodes an array of bytes to b64. This is safer than transporting binary data
func bToB64(p []byte) string {
	return base64.StdEncoding.EncodeToString(p)
}
