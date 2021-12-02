package pokemons

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
)

type PokemonsQueryResolver struct {
	*driver.AppLogger
	*pokemon_interactor.GraphQLModelMapper
	pokemon_interactor.PokemonSearchCommand
}

func (r *PokemonsQueryResolver) Pokemons(first *int, after *int) (model.PokemonConnectionResult, error) {
	p, err := r.PokemonSearchCommand.Execute(first, after)
	if err != nil {
		var e error

		e, ok := err.(*pokemon_interactor.PokemonNotFound)
		if ok {
			r.AppLogger.Print(e.Error())

			return model.PokemonNotFound{
				Message: e.Error(),
			}, nil
		}

		e, ok = err.(*pokemon_interactor.IllegalArguments)
		if ok {
			r.AppLogger.Print(e.Error())

			return model.IllegalArguments{
				Message: e.Error(),
			}, nil
		}

		return nil, err
	}

	pokemons := []*graph.Pokemon{}

	for _, pokemon := range p {
		pokemons = append(pokemons, r.GraphQLModelMapper.Mapping(pokemon))
	}

	token := pokemons[len(pokemons)-1].ID + 1

	return model.PokemonConnection{
		NextID: token,
		Items:  pokemons,
	}, nil
}
