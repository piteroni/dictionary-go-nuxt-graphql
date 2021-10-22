package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/generated"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/pkg/interactor/pokemon"
)

func (r *queryResolver) Pokemon(ctx context.Context, pokemonID int) (*model.Pokemon, error) {
	u := pokemon.NewPokemonDatasetAcquisition(r.DB)

	p, err := u.GetPokemonDataset(pokemonID)
	if err != nil {
		if _, ok := err.(*pokemon.PokemonNotFound); ok {
			r.Logger.Warn(err.Error())

			return nil, err
		}

		return nil, internalSystemError
	}

	genders := []*model.Gender{}
	for _, gender := range p.Genders {
		genders = append(genders, (*model.Gender)(gender))
	}

	types := []*model.Type{}
	for _, t := range p.Types {
		types = append(types, (*model.Type)(t))
	}

	characteristics := []*model.Characteristic{}
	for _, characteristic := range p.Characteristics {
		characteristics = append(characteristics, (*model.Characteristic)(characteristic))
	}

	evolutions := []*model.Pokemon{}
	for _, evolution := range p.Evolutions {
		genders := []*model.Gender{}
		for _, gender := range evolution.Genders {
			genders = append(genders, (*model.Gender)(gender))
		}

		types := []*model.Type{}
		for _, t := range evolution.Types {
			types = append(types, (*model.Type)(t))
		}

		characteristics := []*model.Characteristic{}
		for _, characteristic := range evolution.Characteristics {
			characteristics = append(characteristics, (*model.Characteristic)(characteristic))
		}

		description := (*model.Description)(evolution.Description)
		ability := (*model.Ability)(evolution.Ability)
		link := (*model.LinkInfo)(evolution.LinkInfo)

		evolutions = append(evolutions, &model.Pokemon{
			NationalNo:      evolution.NationalNo,
			Name:            evolution.Name,
			ImageURL:        evolution.ImageURL,
			Species:         evolution.Species,
			Height:          evolution.HeightText,
			Weight:          evolution.WeightText,
			Genders:         genders,
			Types:           types,
			Characteristics: characteristics,
			Description:     description,
			Ability:         ability,
			LinkInfo:        link,
			// 多分いらない
			// Evolutions: evolution.Evolutions,
		})
	}

	description := (*model.Description)(p.Description)

	ability := (*model.Ability)(p.Ability)

	link := (*model.LinkInfo)(p.LinkInfo)

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
		LinkInfo:        link,
		Evolutions:      evolutions,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }