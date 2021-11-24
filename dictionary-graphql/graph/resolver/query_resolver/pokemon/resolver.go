package pokemon

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"

	"gorm.io/gorm"
)

type PokemonQueryResolver struct {
	DB     *gorm.DB
	Logger *driver.AppLogger
}

func (r *PokemonQueryResolver) Pokemon(pokemonID int) (model.PokemonResult, error) {
	command := pokemon_interactor.FindPokemonCommand{DB: r.DB}

	first := 0

	pokemons, err := command.Execute(&first, &pokemonID)
	if err != nil {
		_, ok := err.(*pokemon_interactor.PokemonNotFound)
		if ok {
			return &model.PokemonNotFound{}, nil
		}
	}

	p := pokemon_interactor.GraphQLModel((pokemons)[0])

	return p, nil
}
