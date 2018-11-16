package messages

import (
	"encoding/base64"
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
