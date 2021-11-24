package pokemon_loader

import (
	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"piteroni/dictionary-go-nuxt-graphql/persistence"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"piteroni/dictionary-go-nuxt-graphql/testing/factories"
	"testing"

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

	cleanup := func() {
		err := itesting.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	c := &findPokemonCommand{db: db}

	t.Run("IDに一致するポケモンの情報をデータベースから取得できる", func(t *testing.T) {
		err := seed(t, db)
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		pokemons, err := c.execute(itesting.Int(3), itesting.Int(1))

		assert.Nil(t, err)
		assert.Len(t, pokemons, 3)

		// ステータス値や関連テーブルなどの情報が存在する場合、適切に取得できる
		assert.Contains(t, pokemons, &model.Pokemon{
			Model: gorm.Model{
				ID: uint(1),
			},
			NationalNo:          1,
			Name:                "pokemon-1",
			ImageURL:            "pokemon-1.jpg",
			Species:             "pokemon-1-species",
			Weight:              "10kg",
			Height:              "1m",
			HeartPoint:          10,
			AttackPoint:         10,
			DefensePoint:        10,
			SpecialAttackPoint:  10,
			SpecialDefensePoint: 10,
			SpeedPoint:          10,
			Genders: []model.Gender{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Name:    "gender-1",
					IconURL: "gender-1.jpg",
				},
				{
					Model: gorm.Model{
						ID: uint(2),
					},
					Name:    "gender-2",
					IconURL: "gender-2.jpg",
				},
			},
			Types: []model.Type{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Name:    "type-1",
					IconURL: "type-1.jpg",
				},
				{
					Model: gorm.Model{
						ID: uint(2),
					},
					Name:    "type-2",
					IconURL: "type-2.jpg",
				},
			},
			Characteristics: []model.Characteristic{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Name:        "characteristics-1",
					Description: "characteristics-1-description",
				},
				{
					Model: gorm.Model{
						ID: uint(2),
					},
					Name:        "characteristics-2",
					Description: "characteristics-2-description",
				},
			},
			Descriptions: []model.Description{
				{
					Model: gorm.Model{
						ID: uint(1),
					},
					Text:   "description-1",
					Series: "series-1",
				},
				{
					Model: gorm.Model{
						ID: uint(2),
					},
					Text:   "description-2",
					Series: "series-2",
				},
			},
		})

		// 中途半端にデータが持っている場合は、不足分はデフォルトの値で埋められる.
		assert.Contains(t, pokemons, &model.Pokemon{
			Model: gorm.Model{
				ID: uint(2),
			},
			NationalNo:          2,
			Name:                "pokemon-2",
			ImageURL:            "",
			Species:             "",
			Weight:              "",
			Height:              "",
			HeartPoint:          0,
			AttackPoint:         0,
			DefensePoint:        0,
			SpecialAttackPoint:  0,
			SpecialDefensePoint: 0,
			SpeedPoint:          0,
			Genders: []model.Gender{
				{
					Model: gorm.Model{
						ID: uint(2),
					},
					Name:    "gender-2",
					IconURL: "gender-2.jpg",
				},
			},
			Types:           nil,
			Descriptions:    nil,
			Characteristics: nil,
		})

		// データを全く持ち合わせていない場合は、全てデフォルト値で埋められる.
		assert.Contains(t, pokemons, &model.Pokemon{
			Model: gorm.Model{
				ID: uint(3),
			},
			NationalNo:          3,
			Name:                "pokemon-3",
			ImageURL:            "",
			Species:             "",
			Weight:              "",
			Height:              "",
			HeartPoint:          0,
			AttackPoint:         0,
			DefensePoint:        0,
			SpecialAttackPoint:  0,
			SpecialDefensePoint: 0,
			SpeedPoint:          0,
			Genders:             nil,
			Types:               nil,
			Descriptions:        nil,
			Characteristics:     nil,
		})
	})
}

func seed(t *testing.T, db *gorm.DB) error {
	dao := persistence.NewPokemonDAO(db)
	factory := factories.NewPokemonFactory(db)

	genders := []*model.Gender{}
	gender1 := &model.Gender{
		Model: gorm.Model{
			ID: uint(1),
		},
		Name:    "gender-1",
		IconURL: "gender-1.jpg",
	}
	genders = append(genders, gender1)
	gender2 := &model.Gender{
		Model: gorm.Model{
			ID: uint(2),
		},
		Name:    "gender-2",
		IconURL: "gender-2.jpg",
	}
	genders = append(genders, gender2)

	for _, gender := range genders {
		err := factory.CreateGender(gender)
		if err != nil {
			return err
		}
	}

	types := []*model.Type{}
	type1 := &model.Type{
		Model: gorm.Model{
			ID: uint(1),
		},
		Name:    "type-1",
		IconURL: "type-1.jpg",
	}
	types = append(types, type1)
	type2 := &model.Type{
		Model: gorm.Model{
			ID: uint(2),
		},
		Name:    "type-2",
		IconURL: "type-2.jpg",
	}
	types = append(types, type2)

	for _, t := range types {
		err := factory.CreateType(t)
		if err != nil {
			return err
		}
	}

	characteristics := []*model.Characteristic{}
	characteristic1 := &model.Characteristic{
		Model: gorm.Model{
			ID: uint(1),
		},
		Name:        "characteristics-1",
		Description: "characteristics-1-description",
	}
	characteristics = append(characteristics, characteristic1)
	characteristic2 := &model.Characteristic{
		Model: gorm.Model{
			ID: uint(2),
		},
		Name:        "characteristics-2",
		Description: "characteristics-2-description",
	}
	characteristics = append(characteristics, characteristic2)

	for _, c := range characteristics {
		err := db.Create(c).Error
		if err != nil {
			return err
		}
	}

	pokemon := &model.Pokemon{
		Model:               gorm.Model{ID: 1},
		NationalNo:          1,
		Name:                "pokemon-1",
		ImageURL:            "pokemon-1.jpg",
		Species:             "pokemon-1-species",
		Weight:              "10kg",
		Height:              "1m",
		HeartPoint:          10,
		AttackPoint:         10,
		DefensePoint:        10,
		SpecialAttackPoint:  10,
		SpecialDefensePoint: 10,
		SpeedPoint:          10,
	}

	err := db.Create(pokemon).Error
	if err != nil {
		return err
	}

	description1 := &model.Description{
		Model: gorm.Model{
			ID: uint(1),
		},
		Text:   "description-1",
		Series: "series-1",
	}
	description2 := &model.Description{
		Model: gorm.Model{
			ID: uint(2),
		},
		Text:   "description-2",
		Series: "series-2",
	}

	err = dao.AddDescripton(pokemon, description1)
	if err != nil {
		return err
	}

	err = dao.AddDescripton(pokemon, description2)
	if err != nil {
		return err
	}

	err = dao.AddGender(pokemon, gender1)
	if err != nil {
		return err
	}

	err = dao.AddGender(pokemon, gender2)
	if err != nil {
		return err
	}

	err = dao.AddType(pokemon, type1)
	if err != nil {
		return err
	}

	err = dao.AddType(pokemon, type2)
	if err != nil {
		return err
	}

	err = dao.AddCharacteristics(pokemon, characteristic1)
	if err != nil {
		return err
	}

	err = dao.AddCharacteristics(pokemon, characteristic2)
	if err != nil {
		return err
	}

	pokemon = &model.Pokemon{
		Model:      gorm.Model{ID: 2},
		NationalNo: 2,
		Name:       "pokemon-2",
	}

	err = db.Create(pokemon).Error
	if err != nil {
		return err
	}

	err = dao.AddGender(pokemon, gender2)
	if err != nil {
		return err
	}

	pokemon = &model.Pokemon{
		Model:      gorm.Model{ID: 3},
		NationalNo: 3,
		Name:       "pokemon-3",
	}

	err = db.Create(pokemon).Error
	if err != nil {
		return err
	}

	return nil
}
