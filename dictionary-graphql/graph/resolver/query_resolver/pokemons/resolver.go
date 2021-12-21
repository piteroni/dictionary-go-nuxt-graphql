package pokemons

import (
	"context"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonsQueryResolver struct {
	DB      *mongo.Database
	Context context.Context
	*pokemon_interactor.GraphQLModelMapper
}

func (r *PokemonsQueryResolver) Pokemons(first *int, after *string) (model.PokemonConnectionResult, error) {
	f, a, err := r.prepareParams(first, after)
	if err != nil {
		e, ok := err.(*pokemon_interactor.IllegalArguments)
		if ok {
			return model.IllegalArguments{Message: e.Error()}, nil
		}

		return nil, err
	}

	p, err := r.findPokemons(f, a)
	if err != nil {
		return nil, err
	}

	pokemons := []*model.Pokemon{}

	for _, pokemon := range p {
		pokemons = append(pokemons, r.GraphQLModelMapper.Mapping(pokemon))
	}

	endCursor := ""

	if len(pokemons) != 0 {
		endCursor = pokemons[len(pokemons)-1].ID
	}

	hasNext := len(pokemons) != 0
	pokemonConnection := model.PokemonConnection{
		HasNext:   hasNext,
		EndCursor: endCursor,
		Items:     pokemons,
	}

	return pokemonConnection, nil
}

func (r *PokemonsQueryResolver) prepareParams(first *int, after *string) (int, primitive.ObjectID, error) {
	f := 0

	const (
		min = 0
		max = 64
	)

	if first != nil {
		value := *first

		if value < min {
			err := errors.Cause(&pokemon_interactor.IllegalArguments{
				Message: fmt.Sprintf("first less then %d: first = %d", min, value),
			})

			return 0, primitive.NilObjectID, err
		}

		if value > max {
			err := errors.Cause(&pokemon_interactor.IllegalArguments{
				Message: fmt.Sprintf("first graeter then %d: first = %d", max, value),
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
		e := errors.Cause(&pokemon_interactor.IllegalArguments{Message: err.Error()})
		return 0, primitive.NilObjectID, e
	}

	return f, a, nil
}

func (r *PokemonsQueryResolver) findPokemons(first int, after primitive.ObjectID) ([]*document.Pokemon, error) {
	pipe := document.PokemonAggregate{}.StagesOfLookUp()

	if !after.IsZero() {
		pipe = append(pipe, bson.D{{
			Key: "$match", Value: bson.M{"_id": bson.M{"$gt": after}},
		}})
	}

	pipe = append(pipe, bson.D{{Key: "$limit", Value: first}})

	cursor, err := r.DB.Collection(collection.Pokemons).Aggregate(r.Context, pipe)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pokemons := []*document.Pokemon{}

	err = cursor.All(r.Context, &pokemons)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = cursor.Close(r.Context)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pokemons, nil
}
