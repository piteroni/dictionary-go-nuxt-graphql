package pokemon

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	graph_internal "piteroni/dictionary-go-nuxt-graphql/graph/internal"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon_loader"

	"gorm.io/gorm"
)

type PokemonQueryResolver struct {
	DB     *gorm.DB
	Logger *driver.AppLogger
}

func (r *PokemonQueryResolver) Pokemon(pokemonID int) (model.PokemonResult, error) {
	l := pokemon_loader.NewPokemonLoader(r.DB)
	first := 0

	pokemons, err := l.Pokemons(&first, &pokemonID)
	if err != nil {
		_, ok := err.(*pokemon_loader.PokemonNotFound)
		if ok {
			return &model.PokemonNotFound{}, nil
		}

		r.Logger.Error(err)

		return nil, graph_internal.InternalSystemError
	}

	p := (*pokemons)[0]

	evolutions, err := l.Evolutions(uint(p.ID))
	if err != nil {
		return nil, err
	}

	p.Evolutions = *evolutions

	return p, nil
}
