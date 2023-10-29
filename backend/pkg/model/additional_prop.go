package model

import (
	"github.com/jinzhu/gorm"
)

type AdditionalProp struct {
	gorm.Model
	Name string
}
