package model

import (
	"github.com/jinzhu/gorm"
)

type EventTypeProp struct {
	gorm.Model
	EventTypeID      uint
	AdditionalPropID uint
	Value            string
}
