package pokemon_dataset_acquisition

import (
	"errors"
	"piteroni/dictionary-go-nuxt-graphql/datasource/model"
	"piteroni/dictionary-go-nuxt-graphql/datasource/persistence"

	"gorm.io/gorm"
)

var _ iEvolutionTableAcquisition = (*evolutionTableAcquisition)(nil)

// that provides the pokemon evolution table.
type iEvolutionTableAcquisition interface {
	getEvolutionTable(pokemon *model.Pokemon) ([]*PokemonDataset, error)
}

type evolutionTableAcquisition struct {
	db                   *gorm.DB
	basicInfoAcquisition *basicInfoAcquisition
}

func (i *evolutionTableAcquisition) getEvolutionTable(pokemon *model.Pokemon) ([]*PokemonDataset, error) {
	datasets := []*PokemonDataset{}
	before := &model.Pokemon{}

	r := i.db.Model(&model.Pokemon{}).Where("evolution_id = ?", pokemon.ID).First(before)
	if r.Error != nil {
		// ErrRecordNotFound is an expected error,
		// that occurs when there is no pre-evolution pokemon.
		if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, r.Error
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
					return nil, err
				}
			}

			before = row
		}
	} else {
		before = pokemon
	}

	dao := persistence.NewPokemonDAO(i.db)

	err := dao.ScanEvolution(before)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// when not evolution.
	if before.Evolution == nil {
		return datasets, nil
	}

	// add a starting point.
	dataset, err := i.basicInfoAcquisition.getBasicInfo(before)
	if err != nil {
		return nil, err
	}

	datasets = append(datasets, dataset)

	// tracing evolution.
	for {
		err := dao.ScanEvolution(before)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		if before.Evolution == nil {
			break
		}

		dataset, err := i.basicInfoAcquisition.getBasicInfo(before.Evolution)
		if err != nil {
			return nil, err
		}

		datasets = append(datasets, dataset)

		before = before.Evolution
	}

	return datasets, nil
}
