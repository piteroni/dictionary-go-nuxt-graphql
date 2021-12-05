package evolutions

import (
	"testing"

	"piteroni/dictionary-go-nuxt-graphql/database"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/model"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestEvolutionsQueryResolver(t *testing.T) {
	db, err := itesting.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		err := itesting.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	r := &EvolutionsQueryResolver{
		DB:                 db,
		GraphQLModelMapper: &pokemon_interactor.GraphQLModelMapper{},
	}

	t.Run("指定したポケモンの進化表を取得できる", func(t *testing.T) {
		data := []*model.Pokemon{
			{
				Model:       gorm.Model{ID: 1},
				NationalNo:  1,
				Name:        "pokemon-1",
				EvolutionID: itesting.UInt(2),
			},
			{
				Model:       gorm.Model{ID: 2},
				NationalNo:  2,
				Name:        "pokemon-2",
				EvolutionID: itesting.UInt(3),
			},
			{
				Model:      gorm.Model{ID: 3},
				NationalNo: 3,
				Name:       "pokemon-3",
			},
		}

		err := db.Create(data).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		// same result can be obtained by specifying ID 1 to 3.
		for _, id := range []int{1, 2, 3} {
			evolutions, err := r.Evolutions(id)

			assert.NotNil(t, evolutions)
			assert.Nil(t, err)

			assert.Len(t, evolutions.(graph.Evolutions).Pokemons, 3)
			assert.Equal(t, evolutions, graph.Evolutions{
				Pokemons: []*graph.Pokemon{
					{
						ID:           1,
						NationalNo:   1,
						Name:         "pokemon-1",
						CanEvolution: true,
						Ability:      &graph.Ability{},
						Description:  &graph.Description{},
					},
					{
						ID:           2,
						NationalNo:   2,
						Name:         "pokemon-2",
						CanEvolution: true,
						Ability:      &graph.Ability{},
						Description:  &graph.Description{},
					},
					{
						ID:           3,
						NationalNo:   3,
						Name:         "pokemon-3",
						CanEvolution: false,
						Ability:      &graph.Ability{},
						Description:  &graph.Description{},
					},
				},
			})
		}
	})

	t.Run("進化表データには関連テーブル情報が含まれる", func(t *testing.T) {
		data := []*model.Pokemon{
			{
				Model:       gorm.Model{ID: 1},
				NationalNo:  1,
				Name:        "pokemon-1",
				EvolutionID: itesting.UInt(2),
				Genders: []model.Gender{
					{
						Model:   gorm.Model{ID: 1},
						Name:    "gender-1",
						IconURL: "gender-1.jpg",
					},
				},
				Types: []model.Type{
					{
						Model:   gorm.Model{ID: 1},
						Name:    "type-1",
						IconURL: "type-1.jpg",
					},
				},
				Characteristics: []model.Characteristic{
					{
						Model:       gorm.Model{ID: 1},
						Name:        "characteristic-1",
						Description: "characteristic-1-description",
					},
				},
				Descriptions: []model.Description{
					{
						Model:  gorm.Model{ID: 1},
						Text:   "description-1-text",
						Series: "description-1-series",
					},
				},
			},
			{
				Model:      gorm.Model{ID: 2},
				NationalNo: 2,
				Name:       "pokemon-2",
			},
		}

		err := db.Create(data).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		evolutions, err := r.Evolutions(1)

		assert.NotNil(t, evolutions)
		assert.Nil(t, err)

		assert.Len(t, evolutions.(graph.Evolutions).Pokemons, 2)
		assert.Equal(t, evolutions, graph.Evolutions{
			Pokemons: []*graph.Pokemon{
				{
					ID:           1,
					NationalNo:   1,
					Name:         "pokemon-1",
					CanEvolution: true,
					Ability:      &graph.Ability{},
					Genders: []*graph.Gender{
						{
							Name:    "gender-1",
							IconURL: "gender-1.jpg",
						},
					},
					Types: []*graph.Type{
						{
							Name:    "type-1",
							IconURL: "type-1.jpg",
						},
					},
					Characteristics: []*graph.Characteristic{
						{
							Name:        "characteristic-1",
							Description: "characteristic-1-description",
						},
					},
					Description: &graph.Description{
						Text:   "description-1-text",
						Series: "description-1-series",
					},
				},
				{
					ID:           2,
					NationalNo:   2,
					Name:         "pokemon-2",
					CanEvolution: false,
					Ability:      &graph.Ability{},
					Description:  &graph.Description{},
				},
			},
		})
	})

	t.Run("指定したポケモンが進化しない場合、空の進化表が返ってくる", func(t *testing.T) {
		data := []*model.Pokemon{
			{
				Model:       gorm.Model{ID: 1},
				NationalNo:  1,
				Name:        "pokemon-1",
				EvolutionID: nil,
			},
		}

		err := db.Create(data).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		evolutions, err := r.Evolutions(1)

		assert.NotNil(t, evolutions)
		assert.Nil(t, err)

		assert.Len(t, evolutions.(graph.Evolutions).Pokemons, 0)
		assert.Equal(t, graph.Evolutions{Pokemons: []*graph.Pokemon{}}, evolutions)
	})

	t.Run("指定したポケモンが存在しない場合、例外が送出される", func(t *testing.T) {
		actual, err := r.Evolutions(1)
		expected := graph.PokemonNotFound{}

		assert.NotNil(t, actual)
		assert.Nil(t, err)

		assert.IsType(t, expected, actual)
	})
}
