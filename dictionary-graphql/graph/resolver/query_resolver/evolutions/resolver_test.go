package evolutions

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/database"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"piteroni/dictionary-go-nuxt-graphql/testutils"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
)

func TestEvolutionsQueryResolver(t *testing.T) {
	db, close, err := testutils.ConnnectToTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer close()

	cleanup := func() {
		err = database.Drop(context.Background(), db)
		if err != nil {
			t.Fatal(err)
		}
	}

	r := &EvolutionsQueryResolver{
		DB:                 db,
		Context:            context.Background(),
		GraphQLModelMapper: &pokemon_interactor.GraphQLModelMapper{},
	}

	t.Run("id of pokemon", func(t *testing.T) {
		// #region
		data := map[string][]interface{}{
			collection.Pokemons: {
				document.Pokemon{
					ID:                  testutils.ObjectID(t, "000000000000000000000001"),
					NationalNo:          1,
					Name:                "pokemon-1",
					ImageURL:            "pokemon-1.jpg",
					Species:             "pokemon-1-species",
					Height:              "1.0m",
					Weight:              "1kg",
					EvolutionID:         testutils.ObjectIDPtr(t, "000000000000000000000002"),
					HeartPoint:          10,
					AttackPoint:         10,
					DefensePoint:        10,
					SpecialAttackPoint:  10,
					SpecialDefensePoint: 10,
					SpeedPoint:          10,
					Descriptions: []document.Description{
						{
							ID:     testutils.ObjectID(t, "000000000000000000000001"),
							Text:   "description-1",
							Series: "series-1",
						},
						{
							ID:     testutils.ObjectID(t, "000000000000000000000002"),
							Text:   "description-2",
							Series: "series-2",
						},
					},
					References: document.PokemonReferences{
						Types: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
						Genders: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
						Characteristics: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
					},
				},
				document.Pokemon{
					ID:                  testutils.ObjectID(t, "000000000000000000000002"),
					NationalNo:          2,
					Name:                "pokemon-2",
					ImageURL:            "pokemon-2.jpg",
					Species:             "pokemon-2-species",
					Height:              "1.0m",
					Weight:              "1kg",
					EvolutionID:         testutils.ObjectIDPtr(t, "000000000000000000000003"),
					HeartPoint:          10,
					AttackPoint:         10,
					DefensePoint:        10,
					SpecialAttackPoint:  10,
					SpecialDefensePoint: 10,
					SpeedPoint:          10,
					Descriptions: []document.Description{
						{
							ID:     testutils.ObjectID(t, "000000000000000000000003"),
							Text:   "description-3",
							Series: "series-3",
						},
						{
							ID:     testutils.ObjectID(t, "000000000000000000000004"),
							Text:   "description-4",
							Series: "series-4",
						},
					},
					References: document.PokemonReferences{
						Types: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
						Genders: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
						Characteristics: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
					},
				},
				document.Pokemon{
					ID:                  testutils.ObjectID(t, "000000000000000000000003"),
					NationalNo:          3,
					Name:                "pokemon-3",
					ImageURL:            "pokemon-3.jpg",
					Species:             "pokemon-3-species",
					Height:              "1.0m",
					Weight:              "1kg",
					EvolutionID:         nil,
					HeartPoint:          10,
					AttackPoint:         10,
					DefensePoint:        10,
					SpecialAttackPoint:  10,
					SpecialDefensePoint: 10,
					SpeedPoint:          10,
					Descriptions: []document.Description{
						{
							ID:     testutils.ObjectID(t, "000000000000000000000005"),
							Text:   "description-5",
							Series: "series-5",
						},
						{
							ID:     testutils.ObjectID(t, "000000000000000000000006"),
							Text:   "description-6",
							Series: "series-6",
						},
					},
					References: document.PokemonReferences{
						Types: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
						Genders: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
						Characteristics: []primitive.ObjectID{
							testutils.ObjectID(t, "000000000000000000000001"),
							testutils.ObjectID(t, "000000000000000000000002"),
						},
					},
				},
			},
			collection.Types: {
				document.Type{
					ID:      testutils.ObjectID(t, "000000000000000000000001"),
					Name:    "type-1",
					IconURL: "type-1.jpg",
				},
				document.Type{
					ID:      testutils.ObjectID(t, "000000000000000000000002"),
					Name:    "type-2",
					IconURL: "type-2.jpg",
				},
			},
			collection.Genders: {
				document.Gender{
					ID:      testutils.ObjectID(t, "000000000000000000000001"),
					Name:    "gender-1",
					IconURL: "gender-1.jpg",
				},
				document.Gender{
					ID:      testutils.ObjectID(t, "000000000000000000000002"),
					Name:    "gender-2",
					IconURL: "gender-2.jpg",
				},
			},
			collection.Characteristics: {
				document.Characteristic{
					ID:          testutils.ObjectID(t, "000000000000000000000001"),
					Name:        "characteristic-1",
					Description: "characteristic-1-description",
				},
				document.Characteristic{
					ID:          testutils.ObjectID(t, "000000000000000000000002"),
					Name:        "characteristic-2",
					Description: "characteristic-2-description",
				},
			},
		}
		// #endregion

		// #region
		expected := model.Evolutions{
			Pokemons: []*model.Pokemon{
				{
					ID:           "000000000000000000000001",
					NationalNo:   1,
					Name:         "pokemon-1",
					ImageURL:     "pokemon-1.jpg",
					Species:      "pokemon-1-species",
					Height:       "1.0m",
					Weight:       "1kg",
					CanEvolution: true,
					Description: &model.Description{
						Text:   "description-1",
						Series: "series-1",
					},
					Ability: &model.Ability{
						Heart:          10,
						Attack:         10,
						Defense:        10,
						SpecialAttack:  10,
						SpecialDefense: 10,
						Speed:          10,
					},
					Types: []*model.Type{
						{
							Name:    "type-1",
							IconURL: "type-1.jpg",
						},
						{
							Name:    "type-2",
							IconURL: "type-2.jpg",
						},
					},
					Genders: []*model.Gender{
						{
							Name:    "gender-1",
							IconURL: "gender-1.jpg",
						},
						{
							Name:    "gender-2",
							IconURL: "gender-2.jpg",
						},
					},
					Characteristics: []*model.Characteristic{
						{
							Name:        "characteristic-1",
							Description: "characteristic-1-description",
						},
						{
							Name:        "characteristic-2",
							Description: "characteristic-2-description",
						},
					},
				},
				{
					ID:           "000000000000000000000002",
					NationalNo:   2,
					Name:         "pokemon-2",
					ImageURL:     "pokemon-2.jpg",
					Species:      "pokemon-2-species",
					Height:       "1.0m",
					Weight:       "1kg",
					CanEvolution: true,
					Description: &model.Description{
						Text:   "description-3",
						Series: "series-3",
					},
					Ability: &model.Ability{
						Heart:          10,
						Attack:         10,
						Defense:        10,
						SpecialAttack:  10,
						SpecialDefense: 10,
						Speed:          10,
					},
					Types: []*model.Type{
						{
							Name:    "type-1",
							IconURL: "type-1.jpg",
						},
						{
							Name:    "type-2",
							IconURL: "type-2.jpg",
						},
					},
					Genders: []*model.Gender{
						{
							Name:    "gender-1",
							IconURL: "gender-1.jpg",
						},
						{
							Name:    "gender-2",
							IconURL: "gender-2.jpg",
						},
					},
					Characteristics: []*model.Characteristic{
						{
							Name:        "characteristic-1",
							Description: "characteristic-1-description",
						},
						{
							Name:        "characteristic-2",
							Description: "characteristic-2-description",
						},
					},
				},
				{
					ID:           "000000000000000000000003",
					NationalNo:   3,
					Name:         "pokemon-3",
					ImageURL:     "pokemon-3.jpg",
					Species:      "pokemon-3-species",
					Height:       "1.0m",
					Weight:       "1kg",
					CanEvolution: false,
					Description: &model.Description{
						Text:   "description-5",
						Series: "series-5",
					},
					Ability: &model.Ability{
						Heart:          10,
						Attack:         10,
						Defense:        10,
						SpecialAttack:  10,
						SpecialDefense: 10,
						Speed:          10,
					},
					Types: []*model.Type{
						{
							Name:    "type-1",
							IconURL: "type-1.jpg",
						},
						{
							Name:    "type-2",
							IconURL: "type-2.jpg",
						},
					},
					Genders: []*model.Gender{
						{
							Name:    "gender-1",
							IconURL: "gender-1.jpg",
						},
						{
							Name:    "gender-2",
							IconURL: "gender-2.jpg",
						},
					},
					Characteristics: []*model.Characteristic{
						{
							Name:        "characteristic-1",
							Description: "characteristic-1-description",
						},
						{
							Name:        "characteristic-2",
							Description: "characteristic-2-description",
						},
					},
				},
			},
		}
		// #endregion

		tests := []struct {
			name   string
			params string
		}{
			{
				name:   "first id",
				params: "000000000000000000000001",
			},
			{
				name:   "middle id",
				params: "000000000000000000000002",
			},
			{
				name:   "last id",
				params: "000000000000000000000003",
			},
		}

		for c, d := range data {
			_, err = db.Collection(c).InsertMany(r.Context, d)
			if err != nil {
				t.Fatal(err)
			}
		}

		defer cleanup()

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual, err := r.Evolutions(test.params)

				assert.Nil(t, err)
				assert.Equal(t, expected, actual)
			})
		}
	})
}
