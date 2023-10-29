package model

import (
	"github.com/jinzhu/gorm"
)

type EventType struct {
	gorm.Model
	UserID    uint
	Name      string
	Color     string
	Timestamp string
	Duration  int // Duration in minutes
}
