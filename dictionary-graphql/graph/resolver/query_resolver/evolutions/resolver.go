package evolutions

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/dao"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EvolutionsQueryResolver struct {
	DB      *mongo.Database
	Context context.Context
	*pokemon_interactor.GraphQLModelMapper
}

func (r *EvolutionsQueryResolver) Evolutions(pokemonID string) (model.EvolutionsResult, error) {
	objectID, err := primitive.ObjectIDFromHex(pokemonID)
	if err != nil {
		return model.IllegalArguments{Message: err.Error()}, nil
	}

	pokemon := document.Pokemon{}

	condition := bson.D{{Key: "_id", Value: objectID}}
	opt := options.FindOneOptions{Projection: bson.D{{Key: "_id", Value: 1}}}

	err = r.DB.Collection(collection.Pokemons).FindOne(r.Context, condition, &opt).Decode(&pokemon)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.PokemonNotFound{}, nil
		}

		return nil, errors.WithStack(err)
	}

	err = r.tracingPreEvolution(&pokemon)
	if err != nil {
		return nil, err
	}

	pokemons, err := r.getEvolutions(pokemon.ID)
	if err != nil {
		return nil, err
	}

	p := []*model.Pokemon{}

	for _, pokemon := range pokemons {
		g := r.GraphQLModelMapper.Mapping(pokemon)
		p = append(p, g)
	}

	return model.Evolutions{Pokemons: p}, nil
}

func (r *EvolutionsQueryResolver) tracingPreEvolution(pokemon *document.Pokemon) error {
	opt := options.FindOneOptions{Projection: bson.D{{Key: "_id", Value: 1}}}

	for {
		row := &document.Pokemon{}
		condition := bson.D{{Key: "evolution_id", Value: pokemon.ID}}

		err := r.DB.Collection(collection.Pokemons).FindOne(r.Context, condition, &opt).Decode(&row)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				break
			} else {
				return errors.WithStack(err)
			}
		}

		*pokemon = *row
	}

	return nil
}

func (r *EvolutionsQueryResolver) getEvolutions(pokemonID primitive.ObjectID) ([]*document.Pokemon, error) {
	pokemonDAO := dao.PokemonDAO{
		DB:      r.DB,
		Context: r.Context,
	}

	objectID := &pokemonID
	pokemons := []*document.Pokemon{}

	for {
		if objectID == nil {
			break
		}

		pokemon := document.Pokemon{}

		err := pokemonDAO.FindOneWithLookup(&pokemon, *objectID)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		pokemons = append(pokemons, &pokemon)

		objectID = pokemon.EvolutionID
	}

	// return empty list when evolution.
	if len(pokemons) == 1 && pokemons[0].ID == pokemonID {
		return []*document.Pokemon{}, nil
	}

	return pokemons, nil
}
