package pokemon

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

func TestPokemonQueryResolver(t *testing.T) {
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

	r := &PokemonQueryResolver{
		DB:                 db,
		Context:            context.Background(),
		GraphQLModelMapper: &pokemon_interactor.GraphQLModelMapper{},
	}

	t.Run("id of pokemon", func(t *testing.T) {
		// #region
		data := map[string][]interface{}{
			collection.Pokemons: {
				document.Pokemon{
					ID:                  testutils.ObjectID(t, "000000000000000000000100"),
					NationalNo:          100,
					Name:                "pokemon-100",
					ImageURL:            "pokemon-100.jpg",
					Species:             "pokemon-100-species",
					Height:              "1.0m",
					Weight:              "1kg",
					EvolutionID:         testutils.ObjectIDPtr(t, testutils.ObjectID(t, "000000000000000000001000")),
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

		for c, d := range data {
			_, err = db.Collection(c).InsertMany(r.Context, d)
			if err != nil {
				t.Fatal(err)
			}
		}

		defer cleanup()

		// #region
		expected := &model.Pokemon{
			ID:           "000000000000000000000100",
			NationalNo:   100,
			Name:         "pokemon-100",
			ImageURL:     "pokemon-100.jpg",
			Species:      "pokemon-100-species",
			Height:       "1.0m",
			Weight:       "1kg",
			CanEvolution: true,
			Description: &model.Description{
				Text:   "description-1",
				Series: "series-1",
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
			Ability: &model.Ability{
				Heart:          10,
				Attack:         10,
				Defense:        10,
				SpecialAttack:  10,
				SpecialDefense: 10,
				Speed:          10,
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
		}
		// #endregion

		actual, err := r.Pokemon("000000000000000000000100")

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("invalid pokemon id", func(t *testing.T) {
		actual, err := r.Pokemon("invalid-id")

		assert.Nil(t, err)
		assert.IsType(t, model.IllegalArguments{}, actual)
	})

	t.Run("id of unexists pokemon", func(t *testing.T) {
		actual, err := r.Pokemon("000000000000000000000001")

		assert.Nil(t, err)
		assert.IsType(t, model.PokemonNotFound{}, actual)
	})
}
