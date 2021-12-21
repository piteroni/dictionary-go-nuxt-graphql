package pokemon

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
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

	pipe := mongo.Pipeline{
		{{
			Key: "$match", Value: bson.M{"_id": bson.M{"$eq": objectID}},
		}},
	}

	pipe = append(pipe, document.PokemonAggregate{}.StagesOfLookUp()...)

	cursor, err := r.DB.Collection(collection.Pokemons).Aggregate(r.Context, pipe)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pokemons := []document.Pokemon{}

	err = cursor.All(r.Context, &pokemons)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = cursor.Close(r.Context)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(pokemons) == 0 {
		return model.PokemonNotFound{}, nil
	}

	p := r.GraphQLModelMapper.Mapping(&pokemons[0])

	return p, nil
}
