package persistence

import "gorm.io/gorm"

type Availability struct {
	gorm.Model
	Date string
	Availability bool
}