package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/graph/generated"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon"
)

func (r *queryResolver) Pokemon(ctx context.Context, pokemonID int) (model.PokemonResult, error) {
	qr := pokemon.PokemonQueryResolver{DB: r.DB, Logger: r.Logger}

	return qr.Pokemon(pokemonID)
}

func (r *queryResolver) PageInfo(ctx context.Context, pokemonID int) (model.PageInfoResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Evolutions(ctx context.Context, pokemonID int) (model.EvolutionsResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pokemons(ctx context.Context, first *int, after *int) (*model.PokemonConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
