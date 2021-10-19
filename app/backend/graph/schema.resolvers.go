package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/generated"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/pkg/usecases/pokemon"
)

func (r *queryResolver) Pokemon(ctx context.Context, pokemonID int) (*model.Pokemon, error) {
	u := pokemon.NewPokemonDetailsAcquisition(r.DB)

	p, err := u.GetPokemonDetails(pokemonID)
	if err != nil {
		if _, ok := err.(*pokemon.PokemonNotFoundException); ok {
			r.Logger.Warn(err.Error())

			return nil, err
		}

		return nil, internalSystemError
	}

	genders := []*model.Gender{}
	for _, gender := range p.Genders {
		genders = append(genders, &model.Gender{
			Name:    gender.Name,
			IconURL: gender.IconURL,
		})
	}

	types := []*model.Type{}
	for _, t := range p.Types {
		types = append(types, &model.Type{
			Name:    t.Name,
			IconURL: t.IconURL,
		})
	}

	characteristics := []*model.Characteristic{}
	for _, characteristic := range p.Characteristics {
		characteristics = append(characteristics, &model.Characteristic{
			Name:        characteristic.Name,
			Description: characteristic.Description,
		})
	}

	description := (*model.Description)(p.Description)

	ability := (*model.Ability)(p.Ability)

	return &model.Pokemon{
		NationalNo:      p.NationalNo,
		Name:            p.Name,
		ImageURL:        p.ImageURL,
		Species:         p.Species,
		Height:          p.HeightText,
		Weight:          p.WeightText,
		Genders:         genders,
		Types:           types,
		Characteristics: characteristics,
		Description:     description,
		Ability:         ability,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
