package pokemon

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/graph/internal"
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
	limit := 1

	pokemons, err := l.Load(&limit, &pokemonID)
	if err != nil {
		_, ok := err.(*pokemon_loader.PokemonNotFound)
		if ok {
			return model.PokemonNotFound{}, nil
		}

		r.Logger.Error(err)

		return nil, internal.InternalSystemError
	}

	return (*pokemons)[0], nil
}
