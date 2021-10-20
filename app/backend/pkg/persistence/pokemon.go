package persistence

import (
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"

	"gorm.io/gorm"
)

type PokemonDAO struct {
	db *gorm.DB
}

func NewPokemonDAO(db *gorm.DB) *PokemonDAO {
	return &PokemonDAO{
		db: db,
	}
}

func (dao *PokemonDAO) ScanTypes(p *models.Pokemon) error {
	return dao.db.Model(p).Association("Types").Find(&p.Types)
}

func (dao *PokemonDAO) AddType(p *models.Pokemon, t *models.Type) error {
	return dao.db.Model(p).Association("Types").Append(t)
}

func (dao *PokemonDAO) ScanGenders(p *models.Pokemon) error {
	return dao.db.Model(p).Association("Genders").Find(&p.Genders)
}

func (dao *PokemonDAO) AddGender(p *models.Pokemon, g *models.Gender) error {
	return dao.db.Model(p).Association("Genders").Append(g)
}

func (dao *PokemonDAO) ScanDescriptions(p *models.Pokemon) error {
	return dao.db.Model(p).Association("Descriptions").Find(&p.Descriptions)
}

func (dao *PokemonDAO) AddDescripton(p *models.Pokemon, d *models.Description) error {
	d.PokemonID = p.ID

	return dao.db.Create(d).Error
}

func (dao *PokemonDAO) ScanCharacteristics(p *models.Pokemon) error {
	return dao.db.Model(p).Association("Characteristics").Find(&p.Characteristics)
}

func (dao *PokemonDAO) AddCharacteristics(p *models.Pokemon, c *models.Characteristic) error {
	return dao.db.Model(p).Association("Characteristics").Append(c)
}
