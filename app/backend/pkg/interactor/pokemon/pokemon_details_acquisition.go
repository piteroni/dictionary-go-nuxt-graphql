package pokemon

import (
	"errors"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"
	"piteroni/dictionary-go-nuxt-graphql/pkg/persistence"

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
			return nil, &PokemonNotFound{
				message: fmt.Sprintf("specified pokemon does not exists, pokemonId = %d", pokemonId),
			}
		}

		return nil, err
	}

	dao := persistence.NewPokemonDAO(u.db)

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
			Name:    g.Name,
			IconURL: g.IconURL,
		})
	}

	types := []*Type{}
	for _, t := range pokemon.Types {
		types = append(types, &Type{
			Name:    t.Name,
			IconURL: t.IconURL,
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

	ability := &Ability{
		Heart:          pokemon.HeartPoint,
		Attack:         pokemon.AttackPoint,
		Defense:        pokemon.DefensePoint,
		SpecialAttack:  pokemon.SpecialAttachPoint,
		SpecialDefense: pokemon.SpecialDefensePoint,
		Speed:          pokemon.SpeedPoint,
	}

	transition := &TransitionInfo{
		PrevNationalNo: pokemon.NationalNo - 1,
		NextNationalNo: pokemon.NationalNo + 1,
	}

	var r *gorm.DB

	r = u.db.Model(&models.Pokemon{}).Where("national_no = ?", transition.PrevNationalNo).First(&models.Pokemon{})
	if r.Error != nil {
		if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, r.Error
		}
	}

	transition.HasPrev = r.RowsAffected > 0

	r = u.db.Model(&models.Pokemon{}).Where("national_no = ?", transition.NextNationalNo).First(&models.Pokemon{})
	if r.Error != nil {
		if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, r.Error
		}
	}

	transition.HasNext = r.RowsAffected > 0

	return &PokemonDetails{
		NationalNo:      pokemon.NationalNo,
		Name:            pokemon.Name,
		ImageURL:        pokemon.ImageURL,
		Genders:         genders,
		Species:         pokemon.Species,
		Types:           types,
		HeightText:      pokemon.Height,
		WeightText:      pokemon.Weight,
		Characteristics: characteristics,
		Description:     description,
		Ability:         ability,
		TransitionInfo:  transition,
	}, nil
}

type PokemonDetails struct {
	NationalNo      int
	Name            string
	ImageURL        string
	Species         string
	Types           []*Type
	HeightText      string
	WeightText      string
	Genders         []*Gender
	Characteristics []*Characteristic
	Description     *Description
	Ability         *Ability
	TransitionInfo  *TransitionInfo
}

type Type struct {
	Name    string
	IconURL string
}

type Gender struct {
	Name    string
	IconURL string
}

type Characteristic struct {
	Name        string
	Description string
}

type Description struct {
	Text   string
	Series string
}

type Ability struct {
	Heart          int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

type TransitionInfo struct {
	PrevNationalNo int
	NextNationalNo int
	HasPrev        bool
	HasNext        bool
}

var _ error = (*PokemonNotFound)(nil)

type PokemonNotFound struct {
	message string
}

func (e PokemonNotFound) Error() string {
	return e.message
}
