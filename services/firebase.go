package services

import (
	"context"
	"log"
	"time"

	"firebase.google.com/go/messaging"

	"firebase.google.com/go"

	"google.golang.org/api/option"
)

var app *firebase.App
var oneHour = time.Duration(1) * time.Hour

func init() {
	opt := option.WithCredentialsFile("./apate-cb3e6-firebase-adminsdk-ezgjz-d912eb5ef8.json")
	a, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	app = a
}

// SendFCMMessage sends a message to a device through Firebase Cloud Messaging
func SendFCMMessage(fcmID string, t, b string, data map[string]string) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)

	// TODO: handle badge count

	if err != nil {
		// TODO: handle the error properly
		log.Fatal(err)
	}

	m := &messaging.Message{
		Token: fcmID,
		Data:  data,
		Notification: &messaging.Notification{
			Title: t,
			Body:  b,
		},
		Android: &messaging.AndroidConfig{
			TTL:          &oneHour,
			Notification: &messaging.AndroidNotification{
				// Icon:  "",
				// Color: "",
			},
		},
	}

	res, err := client.Send(ctx, m)

	if err != nil {
		// TODO: handle the error properly
		log.Fatal(err)
	}

	log.Println(res)
	// message successfully sent
}
