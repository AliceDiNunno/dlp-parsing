package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/persistence"
	"fmt"
	"github.com/tbalthazar/onesignal-go"
	"gorm.io/gorm"
	"log"
	"time"
)


var config infra.Config
var db *gorm.DB
var osConfig infra.OneSignal
var notificationClient *onesignal.Client

func LoadService(conf infra.Config, onesignalConfig infra.OneSignal, database *gorm.DB) {
	config = conf
	db = database
	osConfig = onesignalConfig

	notificationClient = onesignal.NewClient(nil)
	notificationClient.AppKey = onesignalConfig.AppKey
}

func messageFromDate(date string) string {
	input := date
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
	return t.Format("Monday 02-Jan-2006")
}

func Tick() {
	availabilities, err := GetAvailabilities(config)
	if err != nil {
		log.Println("Unable to fetch availabilities")
		log.Println(err.Error())
	}

	log.Println(fmt.Sprintf("Found %d availabilities", len(availabilities)))

	for _, availability := range(availabilities) {
		var persistedAvailability persistence.Availability

		req := db.Where("date = ?", availability.Date).First(&persistedAvailability)
		isAvailable := availability.Availability != "none"

		if (req.Error != nil) {
			println(fmt.Sprintf("New availability: %s", availability.Date))

			persistedAvailability = persistence.Availability{}
			persistedAvailability.Date = availability.Date
			persistedAvailability.Availability = isAvailable

			db.Model(&persistedAvailability).Create(&persistedAvailability)

			CreateNotifications(notificationClient, messageFromDate(availability.Date))
		} else {
			if persistedAvailability.Availability != isAvailable && isAvailable {
				println(fmt.Sprintf("Date is now available: %s", persistedAvailability.Date))
				CreateNotifications(notificationClient, messageFromDate(availability.Date))
			}
			db.Model(&persistedAvailability).Where("date = ?", persistedAvailability.Date).Update("availability", isAvailable)
		}


	}

//	sendDeviceMetrics(influxClient, clientsMap)
}
