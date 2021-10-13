package factories

import (
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/imdario/mergo"
	"gorm.io/gorm"
)

type pokemonFactory struct {
	db *gorm.DB
}

func NewPokemonFactory(db *gorm.DB) *pokemonFactory {
	return &pokemonFactory{
		db: db,
	}
}

func (f *pokemonFactory) CreateGender(gender *models.Gender) (*models.Gender, error) {
	defaults := &models.Gender{
		Name:     gofakeit.Name(),
		IconName: gofakeit.UUID(),
	}

	if err := mergo.Merge(gender, *defaults); err != nil {
		return nil, err
	}

	if err := f.db.Create(gender).Error; err != nil {
		return nil, err
	}

	return gender, nil
}

func (f *pokemonFactory) CreatePokemon(pokemon *models.Pokemon) (*models.Pokemon, error) {
	defaults := &models.Pokemon{
		NationalNo:          gofakeit.Number(1, 2048),
		Name:                gofakeit.Name(),
		Species:             gofakeit.Name(),
		ImageName:           gofakeit.UUID(),
		Height:              gofakeit.Noun(),
		Weight:              gofakeit.Noun(),
		HeartPoint:          gofakeit.Number(1, 100),
		AttackPoint:         gofakeit.Number(1, 100),
		DefensePoint:        gofakeit.Number(1, 100),
		SpecialAttachPoint:  gofakeit.Number(1, 100),
		SpecialDefensePoint: gofakeit.Number(1, 100),
		SpeedPoint:          gofakeit.Number(1, 100),
	}

	if err := mergo.Merge(pokemon, *defaults); err != nil {
		return nil, err
	}

	if err := f.db.Create(pokemon).Error; err != nil {
		return nil, err
	}

	return pokemon, nil
}
