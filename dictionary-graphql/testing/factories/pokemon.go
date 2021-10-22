package factories

import (
	"piteroni/dictionary-go-nuxt-graphql/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/imdario/mergo"
	"gorm.io/gorm"
)

type PokemonFactory struct {
	db *gorm.DB
}

func NewPokemonFactory(db *gorm.DB) *PokemonFactory {
	return &PokemonFactory{
		db: db,
	}
}

func (f *PokemonFactory) CreateGender(gender *model.Gender) (*model.Gender, error) {
	defaults := &model.Gender{
		Name:    gofakeit.Name(),
		IconURL: gofakeit.UUID(),
	}

	err := mergo.Merge(gender, *defaults)
	if err != nil {
		return nil, err
	}

	err = f.db.Create(gender).Error
	if err != nil {
		return nil, err
	}

	return gender, nil
}

func (f *PokemonFactory) CreateType(t *model.Type) (*model.Type, error) {
	defaults := &model.Type{
		Name:    gofakeit.Name(),
		IconURL: gofakeit.UUID(),
	}

	err := mergo.Merge(t, *defaults)
	if err != nil {
		return nil, err
	}

	err = f.db.Create(t).Error
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *PokemonFactory) CreateCharacteristic(c *model.Characteristic) (*model.Characteristic, error) {
	defaults := &model.Characteristic{
		Name:        gofakeit.Name(),
		Description: gofakeit.Name(),
	}

	err := mergo.Merge(c, *defaults)
	if err != nil {
		return nil, err
	}

	err = f.db.Create(c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (f *PokemonFactory) CreatePokemon(pokemon *model.Pokemon) (*model.Pokemon, error) {
	defaults := &model.Pokemon{
		NationalNo:          gofakeit.Number(1, 2048),
		Name:                gofakeit.Name(),
		Species:             gofakeit.Name(),
		ImageURL:            gofakeit.UUID(),
		Height:              gofakeit.Noun(),
		Weight:              gofakeit.Noun(),
		HeartPoint:          gofakeit.Number(1, 100),
		AttackPoint:         gofakeit.Number(1, 100),
		DefensePoint:        gofakeit.Number(1, 100),
		SpecialAttachPoint:  gofakeit.Number(1, 100),
		SpecialDefensePoint: gofakeit.Number(1, 100),
		SpeedPoint:          gofakeit.Number(1, 100),
	}

	err := mergo.Merge(pokemon, *defaults)
	if err != nil {
		return nil, err
	}

	err = f.db.Create(pokemon).Error
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
