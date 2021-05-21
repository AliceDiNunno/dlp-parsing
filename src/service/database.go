package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/persistence"
	"log"
)

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDB() *gorm.DB {
	//TODO: move this to configuration
	db, err := gorm.Open(sqlite.Open("./DLP.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to open database: " + err.Error() + "\n")
	}

	db.AutoMigrate(&persistence.Availability{})

	return db
}
