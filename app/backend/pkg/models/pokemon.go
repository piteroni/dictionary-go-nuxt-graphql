package models

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	NationalNo int `gorm:"unique"`
	Species    string

	Types []Type `gorm:"many2many:type_possessed"`

	Characteristics []Characteristic `gorm:"many2many:characteristic_possessed"`

	EvolutionID *uint
	Evolution   *Pokemon

	// profile
	Name         string `gorm:"unique"`
	ImageURL     string
	Descriptions []Description

	// physical
	Height  string
	Weight  string
	Genders []Gender `gorm:"many2many:gender_possessed"`

	// ability status, see https://pokemondb.net/pokedex
	HeartPoint          int
	AttackPoint         int
	DefensePoint        int
	SpecialAttachPoint  int
	SpecialDefensePoint int
	SpeedPoint          int
}
