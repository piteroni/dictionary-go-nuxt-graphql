package pokemon_loader

import (
	"piteroni/dictionary-go-nuxt-graphql/model"
	"piteroni/dictionary-go-nuxt-graphql/persistence"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type evolutionTableAcquisition struct {
	db *gorm.DB
}

func (i *evolutionTableAcquisition) getEvolutionTable(pokemonID uint) (*[]*model.Pokemon, error) {
	var evolutionID *uint

	i.db.Model(&model.Pokemon{}).Select("evolutionID").Where("id = ?", pokemonID).Scan(evolutionID)

	// when not evolution.
	if evolutionID == nil {
		return &[]*model.Pokemon{}, nil
	}

	before := &model.Pokemon{}

	r := i.db.Model(&model.Pokemon{}).Where("evolution_id = ?", evolutionID).First(before)
	if r.Error != nil {
		// ErrRecordNotFound is an expected error,
		// that occurs when there is no pre-evolution pokemon.
		if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(r.Error)
		}
	}

	// tracing pre-evolution.
	if r.RowsAffected != 0 {
		for {
			row := &model.Pokemon{}

			err := i.db.Model(&model.Pokemon{}).Where("evolution_id = ?", before.ID).First(row).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					break
				} else {
					return nil, errors.WithStack(err)
				}
			}

			before = row
		}
	}

	dao := persistence.NewPokemonDAO(i.db)

	err := dao.ScanEvolution(before)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(err)
		}
	}

	// add a starting point.
	pokemons := []*model.Pokemon{before}

	// tracing evolution.
	for {
		err := dao.ScanEvolution(before)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.WithStack(err)
			}
		}

		if before.Evolution == nil {
			break
		}

		pokemons = append(pokemons, before.Evolution)

		before = before.Evolution
	}

	return &pokemons, nil
}
