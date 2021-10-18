package pokemon

import (
	"errors"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"

	"gorm.io/gorm"
)

type PokemonDetailsAcquisition struct {
	db *gorm.DB
}

func NewPokemonDetailsAcquisition(db *gorm.DB) *PokemonDetailsAcquisition {
	return &PokemonDetailsAcquisition{
		db: db,
	}
}

func (u *PokemonDetailsAcquisition) GetPokemonDetails(pokemonId int) (*PokemonDetails, error) {
	pokemon := &models.Pokemon{}

	if err := u.db.Model(&models.Pokemon{}).First(pokemon, pokemonId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &PokemonNotFoundException{
				message: fmt.Sprintf("specified pokemon does not exists, pokemonId = %d", pokemonId),
			}
		}

		return nil, err
	}

	dao := models.NewPokemonDAO(u.db)

	if err := dao.ScanGenders(pokemon); err != nil {
		return nil, err
	}

	if err := dao.ScanTypes(pokemon); err != nil {
		return nil, err
	}

	if err := dao.ScanCharacteristics(pokemon); err != nil {
		return nil, err
	}

	if err := dao.ScanDescriptions(pokemon); err != nil {
		return nil, err
	}

	genders := []*Gender{}
	for _, g := range pokemon.Genders {
		genders = append(genders, &Gender{
			Name:     g.Name,
			IconName: g.IconName,
		})
	}

	types := []*Type{}
	for _, t := range pokemon.Types {
		types = append(types, &Type{
			Name:     t.Name,
			IconName: t.IconName,
		})
	}

	characteristics := []*Characteristic{}
	for _, c := range pokemon.Characteristics {
		characteristics = append(characteristics, &Characteristic{
			Name:        c.Name,
			Description: c.Description,
		})
	}

	description := &Description{}
	if len(pokemon.Descriptions) > 0 {
		description = &Description{
			Text:   pokemon.Descriptions[0].Text,
			Series: pokemon.Descriptions[0].Series,
		}
	}

	return &PokemonDetails{
		NationalNo:      pokemon.NationalNo,
		Name:            pokemon.Name,
		ImageName:       pokemon.ImageName,
		Genders:         genders,
		Species:         pokemon.Species,
		Types:           types,
		HeightText:      pokemon.Height,
		WeightText:      pokemon.Weight,
		Characteristics: characteristics,
		Description:     description,
	}, nil
}

type PokemonDetails struct {
	NationalNo      int
	Name            string
	ImageName       string
	Species         string
	Types           []*Type
	HeightText      string
	WeightText      string
	Genders         []*Gender
	Characteristics []*Characteristic
	Description     *Description
}

type Type struct {
	Name     string
	IconName string
}

type Gender struct {
	Name     string
	IconName string
}

type Characteristic struct {
	Name        string
	Description string
}

type Description struct {
	Text   string
	Series string
}

var _ error = (*PokemonNotFoundException)(nil)

type PokemonNotFoundException struct {
	message string
}

func (e PokemonNotFoundException) Error() string {
	return e.message
}
