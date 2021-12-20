package pokemon_interactor

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/database"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"piteroni/dictionary-go-nuxt-graphql/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidArguments(t *testing.T) {
	c := &PokemonSearchCommandImpl{}

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
			first, after, err := c.prepareParams(test.first, test.after)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedFirst, first)
			assert.Equal(t, test.expectedAfterHex, after.Hex())
		})
	}
}

func TestInValidArguments(t *testing.T) {
	c := &PokemonSearchCommandImpl{}

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
			expectedErrType: &IllegalArguments{},
		},
		{
			name:            "first of integer that exceeds threshold",
			first:           testutils.IntPtr(65),
			after:           testutils.StringPtr("562a94d381cb9f1cd6eb0e1c"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: &IllegalArguments{},
		},
		{
			name:            "after of not hex string of 24 byte",
			first:           testutils.IntPtr(1),
			after:           testutils.StringPtr("invalid-text"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: &IllegalArguments{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			first, after, err := c.prepareParams(test.first, test.after)

			assert.Equal(t, test.expectedFirst, first)
			assert.Equal(t, after, test.expectedAfter)
			assert.IsType(t, err, test.expectedErrType)
		})
	}
}

func TestFindPokemon(t *testing.T) {
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

	c := &PokemonSearchCommandImpl{
		DB:      db,
		Context: context.Background(),
	}

	t.Run("id of pokemon and number of pokemon to be retrieved", func(t *testing.T) {
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
							Text:   "description-1",
							Series: "series-1",
						},
						{
							ID:     testutils.ObjectID(t, "000000000000000000000004"),
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

		expected := []*document.Pokemon{
			{
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
				Types: &[]document.Type{
					{
						ID:      testutils.ObjectID(t, "000000000000000000000001"),
						Name:    "type-1",
						IconURL: "type-1.jpg",
					},
					{
						ID:      testutils.ObjectID(t, "000000000000000000000002"),
						Name:    "type-2",
						IconURL: "type-2.jpg",
					},
				},
				Genders: &[]document.Gender{
					{
						ID:      testutils.ObjectID(t, "000000000000000000000001"),
						Name:    "gender-1",
						IconURL: "gender-1.jpg",
					},
					{
						ID:      testutils.ObjectID(t, "000000000000000000000002"),
						Name:    "gender-2",
						IconURL: "gender-2.jpg",
					},
				},
				Characteristics: &[]document.Characteristic{
					{
						ID:          testutils.ObjectID(t, "000000000000000000000001"),
						Name:        "characteristic-1",
						Description: "characteristic-1-description",
					},
					{
						ID:          testutils.ObjectID(t, "000000000000000000000002"),
						Name:        "characteristic-2",
						Description: "characteristic-2-description",
					},
				},
			},
			{
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
						Text:   "description-1",
						Series: "series-1",
					},
					{
						ID:     testutils.ObjectID(t, "000000000000000000000004"),
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
				Types: &[]document.Type{
					{
						ID:      testutils.ObjectID(t, "000000000000000000000001"),
						Name:    "type-1",
						IconURL: "type-1.jpg",
					},
					{
						ID:      testutils.ObjectID(t, "000000000000000000000002"),
						Name:    "type-2",
						IconURL: "type-2.jpg",
					},
				},
				Genders: &[]document.Gender{
					{
						ID:      testutils.ObjectID(t, "000000000000000000000001"),
						Name:    "gender-1",
						IconURL: "gender-1.jpg",
					},
					{
						ID:      testutils.ObjectID(t, "000000000000000000000002"),
						Name:    "gender-2",
						IconURL: "gender-2.jpg",
					},
				},
				Characteristics: &[]document.Characteristic{
					{
						ID:          testutils.ObjectID(t, "000000000000000000000001"),
						Name:        "characteristic-1",
						Description: "characteristic-1-description",
					},
					{
						ID:          testutils.ObjectID(t, "000000000000000000000002"),
						Name:        "characteristic-2",
						Description: "characteristic-2-description",
					},
				},
			},
		}

		for c, d := range data {
			_, err = db.Collection(c).InsertMany(context.Background(), d)
			if err != nil {
				t.Fatal(err)
			}
		}

		defer cleanup()

		pokemons, err := c.Execute(testutils.IntPtr(3), testutils.StringPtr("000000000000000000000001"))

		assert.Nil(t, err)
		// Pokemons for specified acquisition items can be acquired.
		assert.Len(t, pokemons, 2)
		assert.Equal(t, expected, pokemons)
	})
}
