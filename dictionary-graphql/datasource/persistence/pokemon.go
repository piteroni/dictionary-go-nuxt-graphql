package persistence

import (
	"piteroni/dictionary-go-nuxt-graphql/datasource/model"

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

func (dao *PokemonDAO) ScanTypes(p *model.Pokemon) error {
	if len(p.Types) > 0 {
		return nil
	}

	return dao.db.Model(p).Association("Types").Find(&p.Types)
}

func (dao *PokemonDAO) AddType(p *model.Pokemon, t *model.Type) error {
	return dao.db.Model(p).Association("Types").Append(t)
}

func (dao *PokemonDAO) ScanGenders(p *model.Pokemon) error {
	if len(p.Genders) > 0 {
		return nil
	}

	return dao.db.Model(p).Association("Genders").Find(&p.Genders)
}

func (dao *PokemonDAO) AddGender(p *model.Pokemon, g *model.Gender) error {
	return dao.db.Model(p).Association("Genders").Append(g)
}

func (dao *PokemonDAO) ScanDescriptions(p *model.Pokemon) error {
	if len(p.Descriptions) > 0 {
		return nil
	}

	return dao.db.Model(p).Association("Descriptions").Find(&p.Descriptions)
}

func (dao *PokemonDAO) AddDescripton(p *model.Pokemon, d *model.Description) error {
	d.PokemonID = p.ID

	return dao.db.Create(d).Error
}

func (dao *PokemonDAO) ScanCharacteristics(p *model.Pokemon) error {
	if len(p.Characteristics) > 0 {
		return nil
	}

	return dao.db.Model(p).Association("Characteristics").Find(&p.Characteristics)
}

func (dao *PokemonDAO) AddCharacteristics(p *model.Pokemon, c *model.Characteristic) error {
	return dao.db.Model(p).Association("Characteristics").Append(c)
}

func (dao *PokemonDAO) ScanEvolution(p *model.Pokemon) error {
	if p.Evolution != nil {
		return nil
	}

	if p.EvolutionID == nil {
		return nil
	}

	err := dao.db.Model(&model.Pokemon{}).First(&p.Evolution, *p.EvolutionID).Error
	if err != nil {
		return err
	}

	return nil
}
