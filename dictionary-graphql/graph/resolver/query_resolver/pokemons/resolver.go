package pokemons

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"

	"gorm.io/gorm"
)

type PokemonsQueryResolver struct {
	DB     *gorm.DB
	Logger *driver.AppLogger
}

func (r *PokemonsQueryResolver) Pokemons(first *int, after *int) (model.PokemonConnectionResult, error) {
	command := pokemon_interactor.FindPokemonCommand{DB: r.DB}

	p, err := command.Execute(first, after)
	if err != nil {
		var e error

		e, ok := err.(*pokemon_interactor.PokemonNotFound)
		if ok {
			r.Logger.Print(e.Error())

			return model.PokemonNotFound{
				Message: e.Error(),
			}, nil
		}

		e, ok = err.(*pokemon_interactor.IllegalArgument)
		if ok {
			r.Logger.Print(e.Error())

			return model.IllegalArgument{
				Message: e.Error(),
			}, nil
		}
	}

	pokemons := []*graph.Pokemon{}

	for _, pokemon := range p {
		pokemons = append(pokemons, pokemon_interactor.MappingGraphQLModel(pokemon))
	}

	token := pokemons[len(pokemons)-1].ID

	return model.PokemonConnection{
		NextID: token,
		Items:  pokemons,
	}, nil
}
