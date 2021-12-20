package pokemon

import (
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"

	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonQueryResolver struct {
	*mongo.Database
	*pokemon_interactor.GraphQLModelMapper
	pokemon_interactor.PokemonSearchCommand
}

func (r *PokemonQueryResolver) Pokemon(pokemonID string) (model.PokemonResult, error) {
	first := 0

	pokemons, err := r.PokemonSearchCommand.Execute(&first, &pokemonID)
	if err != nil {
		return nil, err
	}

	if len(pokemons) == 0 {
		return &model.PokemonNotFound{}, nil
	}

	p := r.GraphQLModelMapper.Mapping((pokemons)[0])

	return p, nil
}
