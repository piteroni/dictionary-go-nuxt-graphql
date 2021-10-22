package model

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model

	Types []Type `gorm:"many2many:type_possessed"`

	Characteristics []Characteristic `gorm:"many2many:characteristic_possessed"`

	// @TODO: `gorm:"unique"`
	EvolutionID *uint
	Evolution   *Pokemon

	// identify
	NationalNo int `gorm:"unique"`

	// category
	Species string

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
