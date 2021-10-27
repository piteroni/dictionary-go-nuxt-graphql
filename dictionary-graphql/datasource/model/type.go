package model

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Name    string `gorm:"unique"`
	IconURL string
}
