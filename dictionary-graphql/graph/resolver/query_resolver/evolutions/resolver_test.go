package evolutions

import (
	"testing"

	"piteroni/dictionary-go-nuxt-graphql/database"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
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

	r := &EvolutionsQueryResolver{DB: db}

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

		evolutions, err := r.Evolutions(2)

		assert.NotNil(t, evolutions)
		assert.Nil(t, err)

		assert.IsType(t, graph.Evolutions{}, evolutions)
		assert.Len(t, evolutions.(graph.Evolutions).Pokemons, 3)
		assert.Equal(t, evolutions, graph.Evolutions{
			Pokemons: []*graph.Pokemon{
				{
					ID:           1,
					NationalNo:   1,
					Name:         "pokemon-1",
					CanEvolution: true,
					Ability:      &graph.Ability{},
				},
				{
					ID:           2,
					NationalNo:   2,
					Name:         "pokemon-2",
					CanEvolution: true,
					Ability:      &graph.Ability{},
				},
				{
					ID:           3,
					NationalNo:   3,
					Name:         "pokemon-3",
					CanEvolution: false,
					Ability:      &graph.Ability{},
				},
			},
		})
	})

	t.Run("進化表データには関連テーブル情報が含まれる", func(t *testing.T) {
		data := []*model.Pokemon{
			{
				Model:      gorm.Model{ID: 1},
				NationalNo: 1,
				Name:       "pokemon-1",
				Genders: []model.Gender{
					{
						Model:   gorm.Model{ID: 1},
						Name:    "gender-1",
						IconURL: "gender-1.jpg",
					},
				},
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

		assert.IsType(t, graph.Evolutions{}, evolutions)
		assert.Len(t, evolutions.(graph.Evolutions).Pokemons, 1)
		assert.Equal(t, evolutions, graph.Evolutions{
			Pokemons: []*graph.Pokemon{
				{
					ID:         1,
					NationalNo: 1,
					Name:       "pokemon-1",
					Ability:    &graph.Ability{},
					Genders: []*graph.Gender{
						{
							Name:    "gender-1",
							IconURL: "gender-1.jpg",
						},
					},
				},
			},
		})
	})
}
