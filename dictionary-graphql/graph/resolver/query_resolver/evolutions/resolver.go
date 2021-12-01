package evolutions

import (
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"piteroni/dictionary-go-nuxt-graphql/persistence"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type EvolutionsQueryResolver struct {
	*gorm.DB
	*pokemon_interactor.GraphQLModelMapper
}

func (r *EvolutionsQueryResolver) Evolutions(pokemonID int) (graph.EvolutionsResult, error) {
	pokemon := &model.Pokemon{}

	tx := r.DB.Model(&model.Pokemon{}).Find(pokemon, pokemonID)
	if tx.Error != nil {
		return nil, errors.WithStack(tx.Error)
	}

	if tx.RowsAffected == 0 {
		return graph.PokemonNotFound{}, nil
	}

	err := r.tracingPreEvolution(pokemon)
	if err != nil {
		return nil, err
	}

	pokemons, err := r.getEvolutions(pokemon)
	if err != nil {
		return nil, err
	}

	p := []*graph.Pokemon{}

	for _, pokemon := range pokemons {
		g := r.GraphQLModelMapper.Mapping(pokemon)
		p = append(p, g)
	}

	return graph.Evolutions{Pokemons: p}, nil
}

func (r *EvolutionsQueryResolver) tracingPreEvolution(pokemon *model.Pokemon) error {
	for {
		row := &model.Pokemon{}

		err := r.DB.Model(&model.Pokemon{}).Where("evolution_id = ?", pokemon.ID).First(row).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			} else {
				return errors.WithStack(err)
			}
		}

		*pokemon = *row
	}

	return nil
}

func (r *EvolutionsQueryResolver) getEvolutions(pokemon *model.Pokemon) ([]*model.Pokemon, error) {
	pokemons := []*model.Pokemon{}

	// add a starting point.
	err := r.resolveRelations(pokemon)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pokemons = append(pokemons, pokemon)

	for {
		err := r.resolveRelations(pokemon)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if pokemon.Evolution == nil {
			break
		}

		pokemons = append(pokemons, pokemon.Evolution)

		pokemon = pokemon.Evolution
	}

	// return empty list when evolution.
	if len(pokemons) == 1 && pokemons[0].ID == pokemon.ID {
		return []*model.Pokemon{}, nil
	}

	return pokemons, nil
}

func (r *EvolutionsQueryResolver) resolveRelations(pokemon *model.Pokemon) error {
	dao := persistence.NewPokemonDAO(r.DB)

	err := dao.ScanEvolution(pokemon)
	if err != nil {
		return errors.WithStack(err)
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
