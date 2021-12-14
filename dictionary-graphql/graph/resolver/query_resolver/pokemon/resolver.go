package pokemon

// import (
// 	"piteroni/dictionary-go-nuxt-graphql/graph/model"
// 	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"

// 	"gorm.io/gorm"
// )

// type PokemonQueryResolver struct {
// 	*gorm.DB
// 	*pokemon_interactor.GraphQLModelMapper
// 	pokemon_interactor.PokemonSearchCommand
// }

// func (r *PokemonQueryResolver) Pokemon(pokemonID int) (model.PokemonResult, error) {
// 	first := 0

// 	pokemons, err := r.PokemonSearchCommand.Execute(&first, &pokemonID)
// 	if err != nil {
// 		_, ok := err.(*pokemon_interactor.PokemonNotFound)
// 		if ok {
// 			return &model.PokemonNotFound{}, nil
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	p := r.GraphQLModelMapper.Mapping((pokemons)[0])

// 	return p, nil
// }
