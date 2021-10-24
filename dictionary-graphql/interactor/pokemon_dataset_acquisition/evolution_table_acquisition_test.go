package pokemon_dataset_acquisition

import (
	"piteroni/dictionary-go-nuxt-graphql/database/migration"
	"piteroni/dictionary-go-nuxt-graphql/model"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"piteroni/dictionary-go-nuxt-graphql/testing/factories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestEvolutionTableAcquisition(t *testing.T) {
	db, err := itesting.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = migration.Migrate(db)
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

	i := evolutionTableAcquisition{
		db: db,
		basicInfoAcquisition: &basicInfoAcquisition{
			db: db,
		},
	}

	t.Run("指定したポケモンの進化表を取得できる", func(t *testing.T) {
		pokemon := &model.Pokemon{
			Model:      gorm.Model{ID: 1},
			NationalNo: 1,
			Name:       "pokemon-1",
		}

		err = factory.CreatePokemon(pokemon)
		if err != nil {
			t.Fatal(err)
		}

		evolution1 := &model.Pokemon{
			Model:      gorm.Model{ID: 2},
			NationalNo: 2,
			Name:       "pokemon-2",
		}

		err = factory.CreatePokemon(evolution1)
		if err != nil {
			t.Fatal(err)
		}

		pokemon.EvolutionID = &evolution1.ID

		err := db.Save(pokemon).Error
		if err != nil {
			t.Fatal(err)
		}

		evolution2 := &model.Pokemon{
			Model:      gorm.Model{ID: 3},
			NationalNo: 3,
			Name:       "pokemon-3",
		}

		err = factory.CreatePokemon(evolution2)
		if err != nil {
			t.Fatal(err)
		}

		evolution1.EvolutionID = &evolution2.ID

		err = db.Save(evolution1).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		evolutions, err := i.getEvolutionTable(evolution2)

		assert.NotNil(t, evolutions)
		assert.Nil(t, err)
		assert.Len(t, evolutions, 3)

		assert.Equal(t, evolutions[0].NationalNo, 1)
		assert.Equal(t, evolutions[0].Name, "pokemon-1")

		assert.Equal(t, evolutions[1].NationalNo, 2)
		assert.Equal(t, evolutions[1].Name, "pokemon-2")

		assert.Equal(t, evolutions[2].NationalNo, 3)
		assert.Equal(t, evolutions[2].Name, "pokemon-3")
	})

	t.Run("進化しないポケモンが指定された場合、空の進化表が戻る", func(t *testing.T) {
		pokemon := &model.Pokemon{
			Model:      gorm.Model{ID: 1},
			NationalNo: 1,
			Name:       "pokemon-1",
		}

		err = factory.CreatePokemon(pokemon)
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		evolutions, err := i.getEvolutionTable(pokemon)

		assert.NotNil(t, evolutions)
		assert.Nil(t, err)
		assert.Len(t, evolutions, 0)
	})
}
