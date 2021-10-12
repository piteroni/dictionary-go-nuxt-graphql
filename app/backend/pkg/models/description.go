package models

import "gorm.io/gorm"

type Description struct {
	gorm.Model
	PokemonID uint
	Text      string
	Series    string
}
