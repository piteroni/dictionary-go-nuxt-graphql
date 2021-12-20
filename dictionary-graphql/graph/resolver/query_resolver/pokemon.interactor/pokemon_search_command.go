package pokemon_interactor

import (
	"context"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonSearchCommand interface {
	Execute(first *int, after *string) ([]*document.Pokemon, error)
}

var _ PokemonSearchCommand = (*PokemonSearchCommandImpl)(nil)

type PokemonSearchCommandImpl struct {
	DB      *mongo.Database
	Context context.Context
}

func (c *PokemonSearchCommandImpl) Execute(first *int, after *string) ([]*document.Pokemon, error) {
	f, a, err := c.prepareParams(first, after)
	if err != nil {
		return nil, err
	}

	pokemons, err := c.findPokemons(f, a)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (c *PokemonSearchCommandImpl) prepareParams(first *int, after *string) (int, primitive.ObjectID, error) {
	f := 0

	const (
		min = 0
		max = 64
	)

	if first != nil {
		value := *first

		if value < min {
			err := errors.Cause(&IllegalArguments{
				message: fmt.Sprintf("first less then %d: first = %d", min, value),
			})

			return 0, primitive.NilObjectID, err
		}

		if value > max {
			err := errors.Cause(&IllegalArguments{
				message: fmt.Sprintf("first graeter then %d: first = %d", max, value),
			})

			return 0, primitive.NilObjectID, err
		}

		f = value
	} else {
		f = max
	}

	if after == nil {
		return f, primitive.ObjectID{}, nil
	}

	a, err := primitive.ObjectIDFromHex(*after)
	if err != nil {
		e := errors.Cause(&IllegalArguments{message: err.Error()})
		return 0, primitive.NilObjectID, e
	}

	return f, a, nil
}

func (c *PokemonSearchCommandImpl) findPokemons(first int, after primitive.ObjectID) ([]*document.Pokemon, error) {
	pipeline := mongo.Pipeline{
		{{
			Key: "$limit", Value: first,
		}},
		{{
			Key: "$lookup", Value: bson.M{
				"from":         collection.Types,
				"localField":   "references.types",
				"foreignField": "_id",
				"as":           "types",
			},
		}},
		{{
			Key: "$lookup", Value: bson.M{
				"from":         collection.Genders,
				"localField":   "references.genders",
				"foreignField": "_id",
				"as":           "genders",
			},
		}},
		{{
			Key: "$lookup", Value: bson.M{
				"from":         collection.Characteristics,
				"localField":   "references.characteristics",
				"foreignField": "_id",
				"as":           "characteristics",
			},
		}},
	}

	if !after.IsZero() {
		pipeline = append(pipeline, bson.D{
			{
				Key: "$match", Value: bson.M{
					"_id": bson.M{
						"$gt": after,
					},
				},
			},
		})
	}

	cursor, err := c.DB.Collection(collection.Pokemons).Aggregate(c.Context, pipeline)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pokemons := []*document.Pokemon{}

	err = cursor.All(c.Context, &pokemons)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pokemons, nil
}
