package main

import (
	"context"
	"errors"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/interactor/pokemon_dataset_acquisition"

	"github.com/aws/aws-lambda-go/lambda"
)

type Arguments struct {
	PokemonId int `json:"pokemonId"`
}

func Handler(ctx context.Context, arguments Arguments) (*model.Pokemon, error) {
	if arguments.PokemonId == 0 {
		message := fmt.Sprintf("Illegal arguments arguments.PokemonId = %d", arguments.PokemonId)
		return nil, errors.New(message)
	}

	db, err := database.ConnectToDatabase()
	if err != nil {
		return nil, err
	}

	u := pokemon_dataset_acquisition.New(db)

	p, err := u.GetPokemonDataset(arguments.PokemonId)
	if err != nil {
		if _, ok := err.(*pokemon_dataset_acquisition.PokemonNotFound); ok {
			fmt.Println(err.Error())

			return nil, err
		}

		return nil, err
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

	description := (*model.Description)(p.Description)
	ability := (*model.Ability)(p.Ability)
	link := (*model.LinkInfo)(p.LinkInfo)

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
			Height:          evolution.Height,
			Weight:          evolution.Weight,
			Genders:         genders,
			Types:           types,
			Characteristics: characteristics,
			Description:     description,
			Ability:         ability,
			LinkInfo:        link,
			CanEvolution:    evolution.CanEvolution,
		})
	}

	return &model.Pokemon{
		NationalNo:      p.NationalNo,
		Name:            p.Name,
		ImageURL:        p.ImageURL,
		Species:         p.Species,
		Height:          p.Height,
		Weight:          p.Weight,
		Genders:         genders,
		Types:           types,
		Characteristics: characteristics,
		Description:     description,
		Ability:         ability,
		LinkInfo:        link,
		Evolutions:      evolutions,
		CanEvolution:    p.CanEvolution,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
