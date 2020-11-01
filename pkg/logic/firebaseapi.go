package logic

//
// https://firebase.google.com/docs/admin/setup#go
// https://firebase.google.com/docs/cloud-messaging/send-message
//

import (
	"context"
	"fmt"
	"log"

	"github.com/kcsanad/FirebaseSendPushOpenFaas/pkg/commons"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func initializeAppWithServiceAccount() (*firebase.App, error) {
	// [START initialize_app_service_account_golang]
	//opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	//config := &firebase.Config{ProjectID: "my-project-id"}
	//app, err := firebase.NewApp(context.Background(), config, opt)
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}
	// [END initialize_app_service_account_golang]

	return app, nil
}

func sendToTokenWithCategory(app *firebase.App, category *string, request *commons.GeneralRequestPushNotifStruct) error {
	// [START send_to_token_golang]
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: request.Title,
			Body:  request.Message,
		},
		Data: request.Data,
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Category: *category,
				},
			},
		},
		Token: request.Token,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	// [END send_to_token_golang]

	return nil
}
