package pageinfo

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PageInfoQueryResolver struct {
	DB      *mongo.Database
	Context context.Context
}

func (r *PageInfoQueryResolver) PageInfo(pokemonID string) (model.PageInfoResult, error) {
	objectID, err := primitive.ObjectIDFromHex(pokemonID)
	if err != nil {
		return model.IllegalArguments{Message: err.Error()}, nil
	}

	condition := bson.D{{Key: "_id", Value: objectID}}
	pokemon := document.Pokemon{}
	opt := options.FindOneOptions{Projection: bson.D{{Key: "_id", Value: 1}}}

	err = r.DB.Collection(collection.Pokemons).FindOne(r.Context, condition, &opt).Decode(&pokemon)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.PokemonNotFound{}, nil
		}

		return nil, errors.WithStack(err)
	}

	condition = bson.D{{Key: "_id", Value: bson.M{"$lt": pokemon.ID}}}
	prev := document.Pokemon{}
	opt = options.FindOneOptions{
		Projection: bson.D{{Key: "_id", Value: 1}},
		Sort:       bson.D{{Key: "_id", Value: -1}},
	}

	err = r.DB.Collection(collection.Pokemons).FindOne(r.Context, condition, &opt).Decode(&prev)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.WithStack(err)
		}
	}

	condition = bson.D{{Key: "_id", Value: bson.M{"$gt": pokemon.ID}}}
	next := document.Pokemon{}
	opt = options.FindOneOptions{
		Projection: bson.D{{Key: "_id", Value: 1}},
		Sort:       bson.D{{Key: "_id", Value: 1}},
	}

	err = r.DB.Collection(collection.Pokemons).FindOne(r.Context, condition, &opt).Decode(&next)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.WithStack(err)
		}
	}

	pageInfo := model.PageInfo{}

	if !prev.ID.IsZero() {
		pageInfo.HasPrev = true
		pageInfo.PrevID = prev.ID.Hex()
	}

	if !next.ID.IsZero() {
		pageInfo.HasNext = true
		pageInfo.NextID = next.ID.Hex()
	}

	return pageInfo, nil
}
