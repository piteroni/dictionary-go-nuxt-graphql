package model

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model

	// identify
	NationalNo int

	// profile
	Name         string `gorm:"unique"`
	ImageURL     string
	Descriptions []Description

	// attribute
	Species     string
	Types       []Type `gorm:"many2many:type_possessed"`
	EvolutionID *uint  `gorm:"unique"`
	Evolution   *Pokemon

	// physical
	Height  string
	Weight  string
	Genders []Gender `gorm:"many2many:gender_possessed"`

	// ability status, see https://pokemondb.net/pokedex
	HeartPoint          int
	AttackPoint         int
	DefensePoint        int
	SpecialAttackPoint  int
	SpecialDefensePoint int
	SpeedPoint          int
	Characteristics     []Characteristic `gorm:"many2many:characteristic_possessed"`
}
