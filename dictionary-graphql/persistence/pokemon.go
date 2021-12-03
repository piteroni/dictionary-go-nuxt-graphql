package persistence

import (
	"piteroni/dictionary-go-nuxt-graphql/model"

	"github.com/pkg/errors"
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

	err := dao.db.Model(p).Association("Types").Find(&p.Types)
	if err != nil {
		return err
	}

	return nil
}

func (dao *PokemonDAO) AddType(p *model.Pokemon, t *model.Type) error {
	err := dao.db.Model(p).Association("Types").Append(t)
	if err != nil {
		return err
	}

	return nil
}

func (dao *PokemonDAO) ScanGenders(p *model.Pokemon) error {
	if len(p.Genders) > 0 {
		return nil
	}

	err := dao.db.Model(p).Association("Genders").Find(&p.Genders)
	if err != nil {
		return err
	}

	return nil
}

func (dao *PokemonDAO) AddGender(p *model.Pokemon, g *model.Gender) error {
	err := dao.db.Model(p).Association("Genders").Append(g)
	if err != nil {
		return err
	}

	return nil
}

func (dao *PokemonDAO) ScanDescriptions(p *model.Pokemon) error {
	if len(p.Descriptions) > 0 {
		return nil
	}

	err := dao.db.Model(p).Association("Descriptions").Find(&p.Descriptions)
	if err != nil {
		return err
	}

	return nil
}

func (dao *PokemonDAO) AddDescripton(p *model.Pokemon, d *model.Description) error {
	d.PokemonID = p.ID

	err := dao.db.Create(d).Error
	if err != nil {
		return err
	}

	return nil
}

func (dao *PokemonDAO) ScanCharacteristics(p *model.Pokemon) error {
	if len(p.Characteristics) > 0 {
		return nil
	}

	err := dao.db.Model(p).Association("Characteristics").Find(&p.Characteristics)
	if err != nil {
		return err
	}

	return nil
}

func (dao *PokemonDAO) AddCharacteristics(p *model.Pokemon, c *model.Characteristic) error {
	err := dao.db.Model(p).Association("Characteristics").Append(c)
	if err != nil {
		return err
	}

	return nil
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
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}
