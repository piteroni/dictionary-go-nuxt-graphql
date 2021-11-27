package evolutions

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"piteroni/dictionary-go-nuxt-graphql/persistence"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type EvolutionsQueryResolver struct {
	DB     *gorm.DB
	Logger *driver.AppLogger
}

func (r *EvolutionsQueryResolver) Evolutions(pokemonID int) (graph.EvolutionsResult, error) {
	var evolutionID uint

	err := r.DB.Model(&model.Pokemon{}).Select("evolution_id").Where("id = ?", pokemonID).Row().Scan(&evolutionID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return graph.PokemonNotFound{}, nil
		}

		return nil, errors.WithStack(err)
	}

	// when not evolution.
	if evolutionID == 0 {
		return graph.Evolutions{}, nil
	}

	before := &model.Pokemon{}

	tx := r.DB.Model(&model.Pokemon{}).Where("evolution_id = ?", evolutionID).First(before)
	if tx.Error != nil {
		// ErrRecordNotFound is an expected error,
		// that occurs when there is no pre-evolution pokemon.
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(tx.Error)
		}
	}

	// tracing pre-evolution.
	if tx.RowsAffected != 0 {
		for {
			row := &model.Pokemon{}

			err := r.DB.Model(&model.Pokemon{}).Where("evolution_id = ?", before.ID).First(row).Error
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

	err = r.resolveRelations(before)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// add a starting point.
	pokemons := []*model.Pokemon{before}

	// tracing evolution.
	for {
		err = r.resolveRelations(before)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if before.Evolution == nil {
			break
		}

		pokemons = append(pokemons, before.Evolution)

		before = before.Evolution
	}

	p := []*graph.Pokemon{}

	for _, pokemon := range pokemons {
		g := pokemon_interactor.MappingGraphQLModel(pokemon)
		p = append(p, g)
	}

	return graph.Evolutions{Pokemons: p}, nil
}

func (r *EvolutionsQueryResolver) resolveRelations(pokemon *model.Pokemon) error {
	dao := persistence.NewPokemonDAO(r.DB)

	err := dao.ScanEvolution(pokemon)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithStack(err)
		}
	}

	err = dao.ScanGenders(pokemon)
	if err != nil {
		return errors.WithStack(err)
	}

	err = dao.ScanTypes(pokemon)
	if err != nil {
		return errors.WithStack(err)
	}

	err = dao.ScanCharacteristics(pokemon)
	if err != nil {
		return errors.WithStack(err)
	}

	err = dao.ScanDescriptions(pokemon)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
