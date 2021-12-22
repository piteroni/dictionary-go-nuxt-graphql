package dao

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonDAO struct {
	DB      *mongo.Database
	Context context.Context
}

func (dao *PokemonDAO) FindOneWithLookup(pokemon *document.Pokemon, pokemonID primitive.ObjectID, options ...bson.D) error {
	pipe := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": bson.M{"$eq": pokemonID}}}},
	}

	if options != nil {
		pipe = append(pipe, options...)
	}

	pokemons := []*document.Pokemon{}

	err := dao.FindWithLookup(&pokemons, pipe...)
	if err != nil {
		return err
	}

	if len(pokemons) == 0 {
		return nil
	}

	*pokemon = *pokemons[0]

	return nil
}

func (dao *PokemonDAO) FindWithLookup(pokemons *[]*document.Pokemon, stages ...bson.D) error {
	pipe := mongo.Pipeline{}

	pipe = append(pipe, []bson.D{
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
	}...)

	if stages != nil {
		pipe = append(pipe, stages...)
	}

	cursor, err := dao.DB.Collection(collection.Pokemons).Aggregate(dao.Context, pipe)
	if err != nil {
		return errors.WithStack(err)
	}

	err = cursor.All(dao.Context, pokemons)
	if err != nil {
		return errors.WithStack(err)
	}

	err = cursor.Close(dao.Context)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
