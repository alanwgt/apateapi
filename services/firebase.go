package services

import (
	"context"
	"fmt"
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
func SendFCMMessage(fcmID, t, b, ck, mt, from, mid string) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)

	// TODO: handle badge count

	if err != nil {
		// TODO: handle the error properly
		log.Fatal(err)
	}

	m := &messaging.Message{
		Token: fcmID,
		Notification: &messaging.Notification{
			Title: t,
			Body:  b,
		},
		Data: map[string]string{
			"title":   t,
			"body":    b,
			"mtype":   mt,
			"user":    from,
			"payload": mid,
		},
	}

	res, err := client.Send(ctx, m)

	if err != nil {
		// TODO: handle the error properly
		log.Println("Couldn't send fcm message, reason:")
		log.Println(err)
		return
	}

	fmt.Println("Message sent:", res)
	// message successfully sent
}
