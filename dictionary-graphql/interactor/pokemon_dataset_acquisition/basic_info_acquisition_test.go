package pokemon_dataset_acquisition

import (
	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/datasource/model"
	"piteroni/dictionary-go-nuxt-graphql/datasource/persistence"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"piteroni/dictionary-go-nuxt-graphql/testing/factories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBasicInfoAcquisition(t *testing.T) {
	db, err := itesting.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	factory := factories.NewPokemonFactory(db)
	dao := persistence.NewPokemonDAO(db)

	cleanup := func() {
		err := itesting.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	i := basicInfoAcquisition{db: db}

	t.Run("指定したIDに一致するポケモンの基礎情報を取得できる", func(t *testing.T) {
		pokemon, err := saveBasicInfo(db, factory, dao)
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		dataset, err := i.getBasicInfo(pokemon)

		assert.NotNil(t, dataset)
		assert.Nil(t, err)

		assert.Equal(t, dataset.NationalNo, 1)
		assert.Equal(t, dataset.Name, "pokemon-1")
		assert.Equal(t, dataset.ImageURL, "pokemon-1.jpg")
		assert.Equal(t, dataset.Height, "2m")
		assert.Equal(t, dataset.Weight, "84kg")
		assert.Equal(t, dataset.Species, "normal")

		assert.Equal(t, dataset.Description, &Description{
			Text:   "description",
			Series: "series-1",
		})

		assert.Equal(t, dataset.Ability, &Ability{
			Heart:          30,
			Attack:         31,
			Defense:        32,
			SpecialAttack:  33,
			SpecialDefense: 34,
			Speed:          35,
		})

		assert.Len(t, dataset.Genders, 2)
		assert.Contains(t, dataset.Genders, &Gender{
			Name:    "gender-1",
			IconURL: "gender-1.jpg",
		})
		assert.Contains(t, dataset.Genders, &Gender{
			Name:    "gender-2",
			IconURL: "gender-2.jpg",
		})

		assert.Len(t, dataset.Types, 2)
		assert.Contains(t, dataset.Types, &Type{
			Name:    "type-1",
			IconURL: "type-1.jpg",
		})
		assert.Contains(t, dataset.Types, &Type{
			Name:    "type-2",
			IconURL: "type-2.jpg",
		})

		assert.Len(t, dataset.Characteristics, 2)
		assert.Contains(t, dataset.Characteristics, &Characteristic{
			Name:        "characteristics-1",
			Description: "characteristics-1-description",
		})
		assert.Contains(t, dataset.Characteristics, &Characteristic{
			Name:        "characteristics-2",
			Description: "characteristics-2-description",
		})

		// Out of jurisdiction.
		assert.Nil(t, dataset.LinkInfo)
		assert.Len(t, dataset.Evolutions, 0)
	})

	t.Run("指定したIDに一致するポケモンが進化するか否かを取得できる", func(t *testing.T) {
		t.Run("進化する場合、canEvolutionがtrueに設定される", func(t *testing.T) {
			pokemon := &model.Pokemon{
				Model:      gorm.Model{ID: 1},
				NationalNo: 1,
				Name:       "pokemon-1",
			}

			err = factory.CreatePokemon(pokemon)
			if err != nil {
				t.Fatal(err)
			}

			evolution := &model.Pokemon{
				Model:      gorm.Model{ID: 2},
				NationalNo: 2,
				Name:       "pokemon-2",
			}

			err = factory.CreatePokemon(evolution)
			if err != nil {
				t.Fatal(err)
			}

			pokemon.EvolutionID = &evolution.ID

			err := db.Save(pokemon).Error
			if err != nil {
				t.Fatal(err)
			}

			defer cleanup()

			dataset, err := i.getBasicInfo(pokemon)

			assert.NotNil(t, dataset)
			assert.Nil(t, err)

			assert.Equal(t, dataset.CanEvolution, true)
		})

		t.Run("進化しない場合、canEvolutionがfalseに設定される", func(t *testing.T) {
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

			dataset, err := i.getBasicInfo(pokemon)

			assert.NotNil(t, dataset)
			assert.Nil(t, err)

			assert.Equal(t, dataset.CanEvolution, false)
		})
	})
}

func saveBasicInfo(db *gorm.DB, factory *factories.PokemonFactory, dao *persistence.PokemonDAO) (*model.Pokemon, error) {
	pokemon := &model.Pokemon{
		Model:               gorm.Model{ID: 1},
		NationalNo:          1,
		Name:                "pokemon-1",
		ImageURL:            "pokemon-1.jpg",
		Height:              "2m",
		Weight:              "84kg",
		Species:             "normal",
		HeartPoint:          30,
		AttackPoint:         31,
		DefensePoint:        32,
		SpecialAttackPoint:  33,
		SpecialDefensePoint: 34,
		SpeedPoint:          35,
	}

	err := factory.CreatePokemon(pokemon)
	if err != nil {
		return nil, err
	}

	genders := []*model.Gender{}
	genders = append(genders, &model.Gender{
		Name:    "gender-1",
		IconURL: "gender-1.jpg",
	})
	genders = append(genders, &model.Gender{
		Name:    "gender-2",
		IconURL: "gender-2.jpg",
	})

	for _, gender := range genders {
		err := factory.CreateGender(gender)
		if err != nil {
			return nil, err
		}

		err = dao.AddGender(pokemon, gender)
		if err != nil {
			return nil, err
		}
	}

	types := []*model.Type{}
	types = append(types, &model.Type{
		Name:    "type-1",
		IconURL: "type-1.jpg",
	})
	types = append(types, &model.Type{
		Name:    "type-2",
		IconURL: "type-2.jpg",
	})

	for _, t := range types {
		err := factory.CreateType(t)
		if err != nil {
			return nil, err
		}

		err = dao.AddType(pokemon, t)
		if err != nil {
			return nil, err
		}
	}

	characteristics := []*model.Characteristic{}
	characteristics = append(characteristics, &model.Characteristic{
		Name:        "characteristics-1",
		Description: "characteristics-1-description",
	})
	characteristics = append(characteristics, &model.Characteristic{
		Name:        "characteristics-2",
		Description: "characteristics-2-description",
	})

	for _, c := range characteristics {
		err := factory.CreateCharacteristic(c)
		if err != nil {
			return nil, err
		}

		err = dao.AddCharacteristics(pokemon, c)
		if err != nil {
			return nil, err
		}
	}

	description := &model.Description{
		Text:   "description",
		Series: "series-1",
	}

	err = dao.AddDescripton(pokemon, description)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
