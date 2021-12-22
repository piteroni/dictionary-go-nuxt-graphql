package pokemon

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/mongo/dao"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonQueryResolver struct {
	DB      *mongo.Database
	Context context.Context
	*pokemon_interactor.GraphQLModelMapper
}

func (r *PokemonQueryResolver) Pokemon(pokemonID string) (model.PokemonResult, error) {
	objectID, err := primitive.ObjectIDFromHex(pokemonID)
	if err != nil {
		return model.IllegalArguments{Message: err.Error()}, nil
	}

	pokemonDAO := dao.PokemonDAO{
		DB:      r.DB,
		Context: r.Context,
	}

	pokemon := document.Pokemon{}

	err = pokemonDAO.FindOneWithLookup(&pokemon, objectID)
	if err != nil {
		return nil, err
	}

	if pokemon.ID.IsZero() {
		return model.PokemonNotFound{}, nil
	}

	p := r.GraphQLModelMapper.Mapping(&pokemon)

	return p, nil
}
