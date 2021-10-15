package pokemon

import (
	"piteroni/dictionary-go-nuxt-graphql/pkg/database"
	"piteroni/dictionary-go-nuxt-graphql/pkg/database/factories"
	"piteroni/dictionary-go-nuxt-graphql/pkg/drivers"
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPokemonDetailsAcquisition(t *testing.T) {
	db, err := drivers.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	if err := database.Migrate(db); err != nil {
		t.Fatal(err)
	}

	if err := seed(db); err != nil {
		t.Fatal(err)
	}

	detailsAcquisition := NewPokemonDetailsAcquisition(db)

	t.Cleanup(func() {
		if err := drivers.RefreshInMemoryDatabase(db); err != nil {
			t.Fatal(err)
		}

		if err := seed(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("指定したIDに一致するポケモンの詳細を取得できる", func(t *testing.T) {
		details, err := detailsAcquisition.GetPokemonDetails(1)

		assert.NotNil(t, details)
		assert.Nil(t, err)

		assert.Equal(t, details.NationalNo, 30)
		assert.Equal(t, details.Name, "pokemon-30")
		assert.Equal(t, details.ImageName, "pokemon-30.jpg")
		assert.Equal(t, details.HeightText, "2m")
		assert.Equal(t, details.WeightText, "84kg")
		assert.Equal(t, details.Species, "normal")

		assert.Len(t, details.Genders, 2)
		assert.Contains(t, details.Genders, &Gender{
			Name:     "gender-1",
			IconName: "gender-1.jpg",
		})
		assert.Contains(t, details.Genders, &Gender{
			Name:     "gender-2",
			IconName: "gender-2.jpg",
		})

		assert.Len(t, details.Types, 2)
		assert.Contains(t, details.Types, &Type{
			Name:     "type-1",
			IconName: "type-1.jpg",
		})
		assert.Contains(t, details.Types, &Type{
			Name:     "type-2",
			IconName: "type-2.jpg",
		})

		assert.Len(t, details.Characteristics, 2)
		assert.Contains(t, details.Characteristics, &Characteristic{
			Name:        "characteristics-1",
			Description: "characteristics-1-description",
		})
		assert.Contains(t, details.Characteristics, &Characteristic{
			Name:        "characteristics-2",
			Description: "characteristics-2-description",
		})
	})

	t.Run("指定したIDに一致するポケモンが存在しない場合、エラーが送出される", func(t *testing.T) {
		details, err := detailsAcquisition.GetPokemonDetails(2)

		assert.Nil(t, details)
		assert.NotNil(t, err)
		assert.IsType(t, err, &PokemonNotFoundException{})
	})
}

func seed(db *gorm.DB) error {
	factory := factories.NewPokemonFactory(db)

	pokemon, err := factory.CreatePokemon(&models.Pokemon{
		Model:      gorm.Model{ID: 1},
		NationalNo: 30,
		Name:       "pokemon-30",
		ImageName:  "pokemon-30.jpg",
		Height:     "2m",
		Weight:     "84kg",
		Species:    "normal",
	})
	if err != nil {
		return err
	}

	dao := models.NewPokemonDAO(db)

	genders := []*models.Gender{}
	genders = append(genders, &models.Gender{
		Name:     "gender-1",
		IconName: "gender-1.jpg",
	})
	genders = append(genders, &models.Gender{
		Name:     "gender-2",
		IconName: "gender-2.jpg",
	})

	for _, gender := range genders {
		g, err := factory.CreateGender(gender)
		if err != nil {
			return err
		}

		if err := dao.AddGender(pokemon, g); err != nil {
			return err
		}
	}

	types := []*models.Type{}
	types = append(types, &models.Type{
		Name:     "type-1",
		IconName: "type-1.jpg",
	})
	types = append(types, &models.Type{
		Name:     "type-2",
		IconName: "type-2.jpg",
	})

	for _, t := range types {
		t, err := factory.CreateType(t)
		if err != nil {
			return err
		}

		if err := dao.AddType(pokemon, t); err != nil {
			return err
		}
	}

	characteristics := []*models.Characteristic{}
	characteristics = append(characteristics, &models.Characteristic{
		Name:        "characteristics-1",
		Description: "characteristics-1-description",
	})
	characteristics = append(characteristics, &models.Characteristic{
		Name:        "characteristics-2",
		Description: "characteristics-2-description",
	})

	for _, c := range characteristics {
		c, err := factory.CreateCharacteristic(c)
		if err != nil {
			return err
		}

		if err := dao.AddCharacteristics(pokemon, c); err != nil {
			return err
		}
	}

	return nil
}
