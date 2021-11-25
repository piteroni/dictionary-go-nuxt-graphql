package pokemon_interactor

import (
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/model"
)

func GraphQLModel(pokemon *model.Pokemon) *graph.Pokemon {
	m := &graph.Pokemon{
		ID:           int(pokemon.ID),
		NationalNo:   pokemon.NationalNo,
		Name:         pokemon.Name,
		ImageURL:     pokemon.ImageURL,
		Species:      pokemon.Species,
		Height:       pokemon.Height,
		Weight:       pokemon.Weight,
		CanEvolution: pokemon.EvolutionID != nil,
		Ability: &graph.Ability{
			Heart:          pokemon.HeartPoint,
			Attack:         pokemon.AttackPoint,
			Defense:        pokemon.DefensePoint,
			SpecialAttack:  pokemon.SpecialAttackPoint,
			SpecialDefense: pokemon.SpecialDefensePoint,
			Speed:          pokemon.SpeedPoint,
		},
	}

	for _, g := range pokemon.Genders {
		m.Genders = append(m.Genders, &graph.Gender{
			Name:    g.Name,
			IconURL: g.IconURL,
		})
	}

	for _, t := range pokemon.Types {
		m.Types = append(m.Types, &graph.Type{
			Name:    t.Name,
			IconURL: t.IconURL,
		})
	}

	for _, c := range pokemon.Characteristics {
		m.Characteristics = append(m.Characteristics, &graph.Characteristic{
			Name:        c.Name,
			Description: c.Description,
		})
	}

	if len(pokemon.Descriptions) > 0 {
		m.Description = &graph.Description{
			Text:   pokemon.Descriptions[0].Text,
			Series: pokemon.Descriptions[0].Series,
		}
	}

	return m
}
