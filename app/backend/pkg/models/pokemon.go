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

type PokemonDAO struct {
	db *gorm.DB
}

func NewPokemonDAO(db *gorm.DB) *PokemonDAO {
	return &PokemonDAO{
		db: db,
	}
}

func (dao *PokemonDAO) ScanTypes(p *Pokemon) error {
	return dao.db.Model(p).Association("Types").Find(&p.Types)
}

func (dao *PokemonDAO) AddType(p *Pokemon, t *Type) error {
	return dao.db.Model(p).Association("Types").Append(t)
}

func (dao *PokemonDAO) ScanGenders(p *Pokemon) error {
	return dao.db.Model(p).Association("Genders").Find(&p.Genders)
}

func (dao *PokemonDAO) AddGender(p *Pokemon, g *Gender) error {
	return dao.db.Model(p).Association("Genders").Append(g)
}

func (dao *PokemonDAO) ScanDescriptions(p *Pokemon) error {
	return dao.db.Model(p).Association("Descriptions").Find(&p.Descriptions)
}

func (dao *PokemonDAO) AddDescripton(p *Pokemon, d *Description) error {
	d.PokemonID = p.ID

	return dao.db.Create(d).Error
}

func (dao *PokemonDAO) ScanCharacteristics(p *Pokemon) error {
	return dao.db.Model(p).Association("Characteristics").Find(&p.Characteristics)
}

func (dao *PokemonDAO) AddCharacteristics(p *Pokemon, c *Characteristic) error {
	return dao.db.Model(p).Association("Characteristics").Append(c)
}
