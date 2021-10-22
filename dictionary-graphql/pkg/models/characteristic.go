package models

import "gorm.io/gorm"

type Characteristic struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
}
