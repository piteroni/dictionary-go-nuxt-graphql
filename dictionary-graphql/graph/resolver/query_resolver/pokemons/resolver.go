package pokemons

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"

	"gorm.io/gorm"
)

type PokemonsQueryResolver struct {
	*gorm.DB
	*driver.AppLogger
	*pokemon_interactor.GraphQLModelMapper
	pokemon_interactor.FindPokemonCommand
}

func (r *PokemonsQueryResolver) Pokemons(first *int, after *int) (model.PokemonConnectionResult, error) {
	p, err := r.FindPokemonCommand.Execute(first, after)
	if err != nil {
		var e error

		e, ok := err.(*pokemon_interactor.PokemonNotFound)
		if ok {
			r.AppLogger.Print(e.Error())

			return model.PokemonNotFound{
				Message: e.Error(),
			}, nil
		}

		e, ok = err.(*pokemon_interactor.IllegalArgument)
		if ok {
			r.AppLogger.Print(e.Error())

			return model.IllegalArgument{
				Message: e.Error(),
			}, nil
		}
	}

	pokemons := []*graph.Pokemon{}

	for _, pokemon := range p {
		pokemons = append(pokemons, r.GraphQLModelMapper.Mapping(pokemon))
	}

	token := pokemons[len(pokemons)-1].ID

	return model.PokemonConnection{
		NextID: token,
		Items:  pokemons,
	}, nil
}
