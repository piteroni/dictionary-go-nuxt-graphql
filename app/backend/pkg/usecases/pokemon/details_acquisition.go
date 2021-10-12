package pokemon

import (
	"errors"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"

	"gorm.io/gorm"
)

type DetailsAcquisition struct {
	db *gorm.DB
}

func NewDetailsAcquisition(db *gorm.DB) *DetailsAcquisition {
	return &DetailsAcquisition{
		db: db,
	}
}

func (u *DetailsAcquisition) GetDetailsOfPokemon(pokemonId int) (*PokemonDetails, error) {
	pokemon := &models.Pokemon{}

	if err := u.db.Model(&models.Pokemon{}).First(pokemon, pokemonId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &PokemonNotFoundException{
				message: fmt.Sprintf("specified pokemon does not exists, pokemonId = %d", pokemonId),
			}
		}

		return nil, err
	}

	s := pokemon.Schema(u.db)

	if err := s.ScanGenders(); err != nil {
		return nil, err
	}

	var genders []Gender

	for _, g := range s.Pokemon.Genders {
		genders = append(genders, Gender{
			Name:     g.Name,
			IconPath: g.IconName,
		})
	}

	return &PokemonDetails{
		NationalNo: pokemon.NationalNo,
		Name:       pokemon.Name,
		ImagePath:  pokemon.ImageName,
		Genders:    genders,
	}, nil
}

type PokemonDetails struct {
	NationalNo int
	Name       string
	ImagePath  string
	Genders    []Gender
}

type Gender struct {
	Name     string
	IconPath string
}

var _ error = (*PokemonNotFoundException)(nil)

type PokemonNotFoundException struct {
	message string
}

func (e PokemonNotFoundException) Error() string {
	return e.message
}
