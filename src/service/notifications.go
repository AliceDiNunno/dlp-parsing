package service

import (
	"fmt"
	"github.com/tbalthazar/onesignal-go"
	"log"
)

func CreateNotifications(client *onesignal.Client, message string) string {
	if (client == nil) {
		println("FAILED")
		return ""
	}

	notificationReq := &onesignal.NotificationRequest{
		AppID:            osConfig.AppId,
		Headings: 		  map[string]string{"en": "Nouvelle disponibilitée"},
		Contents:         map[string]string{"en": message},
		IsIOS:            true,
		IncludedSegments: []string{"All"},
	}

	createRes, res, err := client.Notifications.Create(notificationReq)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Println()

	return createRes.ID
}
