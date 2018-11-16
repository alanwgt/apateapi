package vendor

import (
	"context"
	"log"

	"firebase.google.com/go"

	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	opt := option.WithCredentialsFile("../apate-49951-firebase-adminsdk-ijp4t-97140ca736.json")
	a, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	app = a
}
