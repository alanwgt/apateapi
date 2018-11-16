package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/alanwgt/apateapi/database"
	"github.com/alanwgt/apateapi/routes"
	"github.com/alanwgt/apateapi/util"
	"github.com/rs/cors"
)

func main() {
	// opens the file for log output
	f, err := os.OpenFile(util.Conf.Server.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Printf("Couldn't open log file '%s'\n", util.Conf.Server.LogFile)
		log.Fatalln(err)
	}

	defer f.Close()
	log.SetOutput(f)

	r := routes.BuildRouter()
	r.Use(setDefaultHeaders)
	r.Use(decodeBase64Requests)

	corsOpts := genDefaultCorsOpts()

	log.Fatal(http.ListenAndServe(":"+util.Conf.Server.Port, corsOpts.Handler(r)))
}

// Sets the default headers for every request to the server. This is necessary due to CORS
func setDefaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		r.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

// This function will decode every request that uses protobuf and are encoded in base64
// I had to use b64 because I didn't find a way to send a binary protobuf in a request without having some issues
func decodeBase64Requests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// every POST MUST be encoded in base64!
		if r.Method == "POST" {
			d := base64.NewDecoder(base64.StdEncoding, r.Body)
			r.Body = ioutil.NopCloser(d)
		}
		next.ServeHTTP(w, r)
	})
}

// Generates the default CORS options
func genDefaultCorsOpts() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodOptions,
			http.MethodPost,
			http.MethodHead,
		},
		AllowedHeaders: []string{"*"},
	})
}
