package pokemon_dataset_acquisition

import (
	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/model"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"piteroni/dictionary-go-nuxt-graphql/testing/factories"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCommandToFindPokemon(t *testing.T) {
	db, err := itesting.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	factory := factories.NewPokemonFactory(db)

	cleanup := func() {
		err := itesting.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	c := commandToFindPokemon{db: db}

	t.Run("指定したIDに一致するポケモンを取得できる", func(t *testing.T) {
		pokemon := &model.Pokemon{
			Model:      gorm.Model{ID: 1},
			NationalNo: 1,
			Name:       "pokemon-1",
		}

		err := factory.CreatePokemon(pokemon)
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		pokemon, err = c.execute(1)

		assert.Nil(t, err)
		assert.NotNil(t, pokemon)
		assert.Equal(t, pokemon.ID, uint(1))
		assert.Equal(t, pokemon.NationalNo, 1)
		assert.Equal(t, pokemon.Name, "pokemon-1")
	})

	t.Run("指定したIDに一致するポケモンが存在しない場合、エラーが送出される", func(t *testing.T) {
		pokemon, err := c.execute(1)

		assert.Nil(t, pokemon)
		assert.NotNil(t, err)
		assert.IsType(t, errors.Cause(err), &PokemonNotFound{})
	})
}
