package pokemons

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	graph_internal "piteroni/dictionary-go-nuxt-graphql/graph/internal"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon_loader"

	"gorm.io/gorm"
)

type PokemonsQueryResolver struct {
	DB     *gorm.DB
	Logger *driver.AppLogger
}

func (r *PokemonsQueryResolver) Pokemons(first *int, after *int) (model.PokemonConnectionResult, error) {
	l := pokemon_loader.NewPokemonLoader(r.DB)

	p, err := l.Pokemons(first, after)
	if err != nil {
		e, ok := err.(*pokemon_loader.PokemonNotFound)
		if ok {
			r.Logger.Print(e.Error())

			return model.PokemonNotFound{
				Message: e.Error(),
			}, nil
		}

		_, ok = err.(*pokemon_loader.IllegalArgument)
		if ok {
			r.Logger.Print(e.Error())

			return model.IllegalArgument{
				Message: e.Error(),
			}, nil
		}

		r.Logger.Error(err)

		return nil, graph_internal.InternalSystemError
	}

	pokemons := *p
	token := pokemons[len(pokemons)-1].ID

	return model.PokemonConnection{
		NextID: token,
		Items:  pokemons,
	}, nil
}
