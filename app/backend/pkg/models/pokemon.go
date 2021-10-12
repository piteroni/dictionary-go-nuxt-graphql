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
	ImageName    string
	Descriptions []Description

	// physical
	Height  string
	Weight  string
	Genders []Gender `gorm:"many2many:gender_possessed"`

	// ability status
	HeartPoint          int
	AttackPoint         int
	DefensePoint        int
	SpecialAttachPoint  int
	SpecialDefensePoint int
	SpeedPoint          int
}

func (p *Pokemon) Schema(db *gorm.DB) *PokemonSchema {
	return &PokemonSchema{
		db:      db,
		Pokemon: p,
	}
}

type PokemonSchema struct {
	db      *gorm.DB
	Pokemon *Pokemon
}

func (s *PokemonSchema) ScanTypes() error {
	return s.db.Model(s.Pokemon).Association("Types").Find(&s.Pokemon.Types)
}

func (s *PokemonSchema) AddType(t *Type) error {
	return s.db.Model(s.Pokemon).Association("Types").Append(t)
}

func (s *PokemonSchema) ScanGenders() error {
	return s.db.Model(s.Pokemon).Association("Genders").Find(&s.Pokemon.Genders)
}

func (s *PokemonSchema) AddGender(g *Gender) error {
	return s.db.Model(s.Pokemon).Association("Genders").Append(g)
}

func (s *PokemonSchema) ScanDescriptons() error {
	return s.db.Model(s.Pokemon).Association("Descripton").Find(&s.Pokemon.Descriptions)
}

func (s *PokemonSchema) AddDescripton(d *Description) error {
	d.PokemonID = s.Pokemon.ID

	return s.db.Create(d).Error
}

func (s *PokemonSchema) ScanCharacteristics() error {
	return s.db.Model(s.Pokemon).Association("Characteristics").Find(&s.Pokemon.Characteristics)
}

func (s *PokemonSchema) AddCharacteristics(c *Characteristic) error {
	return s.db.Model(s.Pokemon).Association("Characteristics").Append(c)
}
