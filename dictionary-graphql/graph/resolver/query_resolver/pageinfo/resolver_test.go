package pageinfo

import (
	"context"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/database"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"piteroni/dictionary-go-nuxt-graphql/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageInfoQueryResolver(t *testing.T) {
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

	r := &PageInfoQueryResolver{DB: db, Context: context.Background()}

	t.Run("id of pokemon", func(t *testing.T) {
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

		tests := []struct {
			name     string
			params   string
			expected model.PageInfo
		}{
			{
				name:   "first pokemon id",
				params: "000000000000000000000001",
				expected: model.PageInfo{
					PrevID:  "",
					NextID:  "000000000000000000000002",
					HasPrev: false,
					HasNext: true,
				},
			},
			{
				name:   "middle pokemon id",
				params: "000000000000000000000002",
				expected: model.PageInfo{
					PrevID:  "000000000000000000000001",
					NextID:  "000000000000000000000003",
					HasPrev: true,
					HasNext: true,
				},
			},
			{
				name:   "last pokemon id",
				params: "000000000000000000000003",
				expected: model.PageInfo{
					PrevID:  "000000000000000000000002",
					NextID:  "",
					HasPrev: true,
					HasNext: false,
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual, err := r.PageInfo(test.params)

				assert.Nil(t, err, fmt.Errorf("%+v", err))
				assert.Equal(t, test.expected, actual)
			})
		}
	})

	t.Run("id of unexists pokemon", func(t *testing.T) {
		actual, err := r.PageInfo("000000000000000000000001")

		assert.Nil(t, err)
		assert.IsType(t, model.PokemonNotFound{}, actual)
	})

	t.Run("invalid pokemon id", func(t *testing.T) {
		actual, err := r.PageInfo("invalid-pokemon-id")

		assert.Nil(t, err)
		assert.IsType(t, model.IllegalArguments{}, actual)
	})
}
