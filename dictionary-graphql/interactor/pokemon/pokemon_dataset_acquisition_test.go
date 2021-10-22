package pokemon

import (
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/database/migration"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"piteroni/dictionary-go-nuxt-graphql/persistence"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"piteroni/dictionary-go-nuxt-graphql/testing/factories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPokemonDatasetAcquisition(t *testing.T) {
	db, err := itesting.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = migration.Migrate(db)
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

	datasetAcquisition := NewPokemonDatasetAcquisition(db)

	t.Run("指定したIDに一致するポケモンの詳細を取得できる", func(t *testing.T) {
		err := seed(db, factory, dao)
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		dataset, err := datasetAcquisition.GetPokemonDataset(2)

		assert.NotNil(t, dataset)
		assert.Nil(t, err)

		assert.Equal(t, dataset.NationalNo, 2)
		assert.Equal(t, dataset.Name, "pokemon-2")
		assert.Equal(t, dataset.ImageURL, "pokemon-2.jpg")
		assert.Equal(t, dataset.HeightText, "2m")
		assert.Equal(t, dataset.WeightText, "84kg")
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
	})

	t.Run("データベース上で前後にポケモンが登録されているか取得できる", func(t *testing.T) {
		pokemon := &model.Pokemon{
			Model:      gorm.Model{ID: 1},
			NationalNo: 1,
			Name:       "pokemon-1",
		}

		err = factory.CreatePokemon(pokemon)
		if err != nil {
			t.Fatal(err)
		}

		next := &model.Pokemon{
			Model:      gorm.Model{ID: 2},
			NationalNo: 2,
			Name:       "pokemon-2",
		}

		err = factory.CreatePokemon(next)
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		t.Run("IDが1の場合", func(t *testing.T) {
			dataset, err := datasetAcquisition.GetPokemonDataset(1)

			assert.NotNil(t, dataset)
			assert.Nil(t, err)

			assert.Equal(t, dataset.LinkInfo, &LinkInfo{
				PrevNationalNo: 0,
				NextNationalNo: 2,
				HasPrev:        false,
				HasNext:        true,
			})
		})

		t.Run("IDが2の場合", func(t *testing.T) {
			dataset, err := datasetAcquisition.GetPokemonDataset(2)

			assert.NotNil(t, dataset)
			assert.Nil(t, err)

			assert.Equal(t, dataset.LinkInfo, &LinkInfo{
				PrevNationalNo: 1,
				NextNationalNo: 3,
				HasPrev:        true,
				HasNext:        false,
			})
		})
	})

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

		err := db.Save(pokemon).Error
		if err != nil {
			t.Fatal(err)
		}

		err = db.Save(evolution1).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		evolutions, err := datasetAcquisition.getEvolutionTable(evolution1)

		fmt.Printf("evolutions[0].CanEvolution: %v\n", evolutions[0].CanEvolution)
		fmt.Printf("evolutions[1].CanEvolution: %v\n", evolutions[1].CanEvolution)
		fmt.Printf("evolutions[2].CanEvolution: %v\n", evolutions[2].CanEvolution)

		assert.NotNil(t, evolutions)
		assert.Nil(t, err)
		assert.Len(t, evolutions, 3)
		// @TODO: 構造体チェック、一部属性のみチェックってどうやるかわからないけど
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

		evolutions, err := datasetAcquisition.getEvolutionTable(pokemon)

		assert.NotNil(t, evolutions)
		assert.Nil(t, err)
		assert.Len(t, evolutions, 0)
	})

	t.Run("指定したIDに一致するポケモンが存在しない場合、エラーが送出される", func(t *testing.T) {
		dataset, err := datasetAcquisition.GetPokemonDataset(2)

		assert.Nil(t, dataset)
		assert.NotNil(t, err)
		assert.IsType(t, err, &PokemonNotFound{})
	})
}

func seed(db *gorm.DB, factory *factories.PokemonFactory, dao *persistence.PokemonDAO) error {
	pokemon := &model.Pokemon{
		Model:               gorm.Model{ID: 2},
		NationalNo:          2,
		Name:                "pokemon-2",
		ImageURL:            "pokemon-2.jpg",
		Height:              "2m",
		Weight:              "84kg",
		Species:             "normal",
		HeartPoint:          30,
		AttackPoint:         31,
		DefensePoint:        32,
		SpecialAttachPoint:  33,
		SpecialDefensePoint: 34,
		SpeedPoint:          35,
	}

	// first pokemon.
	err := factory.CreatePokemon(pokemon)
	if err != nil {
		return err
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
			return err
		}

		err = dao.AddGender(pokemon, gender)
		if err != nil {
			return err
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
			return err
		}

		err = dao.AddType(pokemon, t)
		if err != nil {
			return err
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
			return err
		}

		err = dao.AddCharacteristics(pokemon, c)
		if err != nil {
			return err
		}
	}

	description := &model.Description{
		Text:   "description",
		Series: "series-1",
	}

	err = dao.AddDescripton(pokemon, description)
	if err != nil {
		return err
	}

	return nil
}
