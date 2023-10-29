package model

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	EventTypeID uint
	Description string
	Timestamp   string
}
