package pokemon_interactor

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/database"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"testing"

	"github.com/pkg/errors"
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
			first:            itesting.Int(0),
			after:            itesting.String("562a94d381cb9f1cd6eb0e1c"),
			expectedFirst:    0,
			expectedAfterHex: "562a94d381cb9f1cd6eb0e1c",
		},
		{
			name:             "first of integer that not exceeds threshold",
			first:            itesting.Int(64),
			after:            itesting.String("562a94d381cb9f1cd6eb0e1a"),
			expectedFirst:    64,
			expectedAfterHex: "562a94d381cb9f1cd6eb0e1a",
		},
		{
			name:  "first of nil",
			first: nil,
			after: itesting.String("562a94d381cb9f1cd6eb0e1b"),
			// default value is set.
			expectedFirst:    64,
			expectedAfterHex: "562a94d381cb9f1cd6eb0e1b",
		},
		{
			name:             "after on nil",
			first:            itesting.Int(1),
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
			first:           itesting.Int(-1),
			after:           itesting.String("562a94d381cb9f1cd6eb0e1c"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: &IllegalArguments{},
		},
		{
			name:            "first of integer that exceeds threshold",
			first:           itesting.Int(65),
			after:           itesting.String("562a94d381cb9f1cd6eb0e1c"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: &IllegalArguments{},
		},
		{
			name:            "after of not hex string of 24 byte",
			first:           itesting.Int(1),
			after:           itesting.String("invalid-text"),
			expectedFirst:   0,
			expectedAfter:   primitive.NilObjectID,
			expectedErrType: errors.WithStack(errors.New("")),
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
	db, close, err := itesting.ConnnectToTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer close()

	cleanup := func() {
		err = database.Drop(context.Background(), db)
		if err != nil {
			panic(err)
		}
	}

	c := &PokemonSearchCommandImpl{
		DB:      db,
		Context: context.Background(),
	}

	t.Run("id of pokemon and number of pokemon to be retrieved", func(t *testing.T) {
		_, err := db.Collection(collection.Pokemons).InsertOne(context.Background(), document.Pokemon{})
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		pokemons, err := c.Execute(itesting.Int(1), nil)

		assert.Nil(t, err)
		assert.NotNil(t, pokemons)
	})
}
