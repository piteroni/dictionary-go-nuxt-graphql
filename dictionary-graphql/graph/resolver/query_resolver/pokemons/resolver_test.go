package pokemons

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/database"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"piteroni/dictionary-go-nuxt-graphql/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidArguments(t *testing.T) {
	r := &PokemonsQueryResolver{}

	tests := []struct {
		name             string
		first            *int
		after            *string
		expectedFirst    int
		expectedAfterHex string
	}{
		{
			name:             "first of integer and after on hex string of 24 byte",
			first:            testutils.IntPtr(0),
			after:            testutils.StringPtr("562a94d381cb9f1cd6eb0e1c"),
			expectedFirst:    0,
			expectedAfterHex: "562a94d381cb9f1cd6eb0e1c",
		},
		{
			name:             "first of integer that not exceeds threshold",
			first:            testutils.IntPtr(64),
			after:            testutils.StringPtr("562a94d381cb9f1cd6eb0e1a"),
			expectedFirst:    64,
			expectedAfterHex: "562a94d381cb9f1cd6eb0e1a",
		},
		{
			name:  "first of nil",
			first: nil,
			after: testutils.StringPtr("562a94d381cb9f1cd6eb0e1b"),
			// default value is set.
			expectedFirst:    64,
			expectedAfterHex: "562a94d381cb9f1cd6eb0e1b",
		},
		{
			name:             "after on nil",
			first:            testutils.IntPtr(1),
			after:            nil,
			expectedFirst:    1,
			expectedAfterHex: "000000000000000000000000",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			first, after, err := r.prepareParams(test.first, test.after)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedFirst, first)
			assert.Equal(t, test.expectedAfterHex, after.Hex())
		})
	}
}

func TestInValidArguments(t *testing.T) {
	r := &PokemonsQueryResolver{}

	tests := []struct {
		name            string
		first           *int
		after           *string
		expectedFirst   int
		expectedAfter   interface{}
		expectedErrType interface{}
	}{
		{
			name:            "first of negative number",
			first:           testutils.IntPtr(-1),
			after:           testutils.StringPtr("562a94d381cb9f1cd6eb0e1c"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: &pokemon_interactor.IllegalArguments{},
		},
		{
			name:            "first of integer that exceeds threshold",
			first:           testutils.IntPtr(65),
			after:           testutils.StringPtr("562a94d381cb9f1cd6eb0e1c"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: &pokemon_interactor.IllegalArguments{},
		},
		{
			name:            "after of not hex string of 24 byte",
			first:           testutils.IntPtr(1),
			after:           testutils.StringPtr("invalid-text"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: &pokemon_interactor.IllegalArguments{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			first, after, err := r.prepareParams(test.first, test.after)

			assert.Equal(t, test.expectedFirst, first)
			assert.Equal(t, after, test.expectedAfter)
			assert.IsType(t, err, test.expectedErrType)
		})
	}
}

func TestFindPokemons(t *testing.T) {
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

	r := &PokemonsQueryResolver{
		DB:                 db,
		Context:            context.Background(),
		GraphQLModelMapper: &pokemon_interactor.GraphQLModelMapper{},
	}

	t.Run("id of pokemon and number of pokemon to be retrieved", func(t *testing.T) {
		// #region
		data := map[string][]interface{}{
			collection.Pokemons: {
				document.Pokemon{
					ID: testutils.ObjectID(t, "000000000000000000000001"),
				},
				document.Pokemon{
					ID:                  testutils.ObjectID(t, "000000000000000000000002"),
					NationalNo:          2,
					Name:                "pokemon-2",
					ImageURL:            "pokemon-2.jpg",
					Species:             "pokemon-2-species",
					Height:              "1.0m",
					Weight:              "1kg",
					EvolutionID:         testutils.ObjectIDPtr(t, testutils.ObjectID(t, "000000000000000000000010")),
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
					ID:                  testutils.ObjectID(t, "000000000000000000000003"),
					NationalNo:          3,
					Name:                "pokemon-3",
					ImageURL:            "pokemon-3.jpg",
					Species:             "pokemon-3-species",
					Height:              "1.0m",
					Weight:              "1kg",
					EvolutionID:         testutils.ObjectIDPtr(t, testutils.ObjectID(t, "000000000000000000000011")),
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
		expected := model.PokemonConnection{
			HasNext:   true,
			EndCursor: "000000000000000000000003",
			Items: []*model.Pokemon{
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
				},
				{
					ID:           "000000000000000000000003",
					NationalNo:   3,
					Name:         "pokemon-3",
					ImageURL:     "pokemon-3.jpg",
					Species:      "pokemon-3-species",
					Height:       "1.0m",
					Weight:       "1kg",
					CanEvolution: true,
					Description: &model.Description{
						Text:   "description-3",
						Series: "series-3",
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
				},
			},
		}
		// #endregion

		for c, d := range data {
			_, err = db.Collection(c).InsertMany(context.Background(), d)
			if err != nil {
				t.Fatal(err)
			}
		}

		defer cleanup()

		actual, err := r.Pokemons(testutils.IntPtr(3), testutils.StringPtr("000000000000000000000001"))

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("invalid pokemon id", func(t *testing.T) {
		actual, err := r.Pokemons(testutils.IntPtr(1), testutils.StringPtr("invalid-id"))

		assert.Nil(t, err)
		assert.IsType(t, model.IllegalArguments{}, actual)
	})
}

func TestBoundaryOfFindPokemonsResult(t *testing.T) {
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

	r := &PokemonsQueryResolver{
		DB:                 db,
		Context:            context.Background(),
		GraphQLModelMapper: &pokemon_interactor.GraphQLModelMapper{},
	}

	data := []interface{}{
		document.Pokemon{
			ID: testutils.ObjectID(t, "000000000000000000000001"),
		},
		document.Pokemon{
			ID: testutils.ObjectID(t, "000000000000000000000002"),
		},
		document.Pokemon{
			ID: testutils.ObjectID(t, "000000000000000000000003"),
		},
	}

	_, err = db.Collection(collection.Pokemons).InsertMany(r.Context, data)
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	// #region
	tests := []struct {
		name     string
		first    *int
		after    *string
		expected model.PokemonConnection
	}{
		{
			name:  "after paramter is 1",
			first: testutils.IntPtr(1),
			after: testutils.StringPtr("000000000000000000000001"),
			expected: model.PokemonConnection{
				HasNext:   true,
				EndCursor: "000000000000000000000002",
				Items: []*model.Pokemon{
					{
						ID:          "000000000000000000000002",
						Description: &model.Description{},
						Ability:     &model.Ability{},
					},
				},
			},
		},
		{
			name:  "after paramter is 2",
			first: testutils.IntPtr(1),
			after: testutils.StringPtr("000000000000000000000002"),
			expected: model.PokemonConnection{
				HasNext:   true,
				EndCursor: "000000000000000000000003",
				Items: []*model.Pokemon{
					{
						ID:          "000000000000000000000003",
						Description: &model.Description{},
						Ability:     &model.Ability{},
					},
				},
			},
		},
		{
			name:  "after paramter is 3",
			first: testutils.IntPtr(1),
			after: testutils.StringPtr("000000000000000000000003"),
			expected: model.PokemonConnection{
				HasNext:   false,
				EndCursor: "",
				Items:     []*model.Pokemon{},
			},
		},
	}
	// #endregion

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := r.Pokemons(test.first, test.after)

			assert.Nil(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
