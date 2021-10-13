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

		assert.Len(t, details.Genders, 2)
		assert.Contains(t, details.Genders, Gender{
			Name:     "gender-1",
			IconPath: "gender-1.jpg",
		})
		assert.Contains(t, details.Genders, Gender{
			Name:     "gender-2",
			IconPath: "gender-2.jpg",
		})
	})

	t.Run("指定したIDに一致するポケモンが存在しない場合、エラーが送出される", func(t *testing.T) {
		details, err := detailsAcquisition.GetPokemonDetails(3)

		assert.Nil(t, details)
		assert.NotNil(t, err)
		assert.IsType(t, err, &PokemonNotFoundException{})
	})
}

func seed(db *gorm.DB) error {
	factory := factories.NewPokemonFactory(db)

	pokemon, err := factory.CreatePokemon(&models.Pokemon{
		NationalNo: 30,
		Name:       "pokemon-30",
		ImageName:  "pokemon-30.jpg",
	})
	if err != nil {
		return err
	}

	s := pokemon.Schema(db)

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

		if err := s.AddGender(g); err != nil {
			return err
		}
	}

	return nil
}
