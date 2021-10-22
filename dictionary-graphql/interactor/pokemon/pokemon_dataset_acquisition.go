package pokemon

import (
	"errors"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"piteroni/dictionary-go-nuxt-graphql/persistence"

	"gorm.io/gorm"
)

type PokemonDatasetAcquisition struct {
	db *gorm.DB
}

func NewPokemonDatasetAcquisition(db *gorm.DB) *PokemonDatasetAcquisition {
	return &PokemonDatasetAcquisition{
		db: db,
	}
}

func (u *PokemonDatasetAcquisition) GetPokemonDataset(pokemonId int) (*PokemonDataset, error) {
	pokemon := &model.Pokemon{}

	if err := u.db.Model(&model.Pokemon{}).First(pokemon, pokemonId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &PokemonNotFound{
				message: fmt.Sprintf("specified pokemon does not exists, pokemonId = %d", pokemonId),
			}
		}

		return nil, err
	}

	dataset, err := u.constructPokemonDataset(pokemon)
	if err != nil {
		return nil, err
	}

	evolutions, err := u.getEvolutionTable(pokemon)
	if err != nil {
		return nil, err
	}

	dataset.Evolutions = evolutions

	return dataset, nil
}

func (u *PokemonDatasetAcquisition) getEvolutionTable(pokemon *model.Pokemon) ([]*PokemonDataset, error) {
	datasets := []*PokemonDataset{}
	before := &model.Pokemon{}

	// backward
	r := u.db.Model(&model.Pokemon{}).Where("evolution_id = ?", pokemon.ID).First(before)
	if r.Error != nil {
		if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, r.Error
		}
	}

	if r.RowsAffected != 0 {
		for {
			beforeId := before.ID
			row := &model.Pokemon{}

			err := u.db.Model(&model.Pokemon{}).Where("evolution_id = ?", beforeId).First(row).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					break
				} else {
					return nil, err
				}
			}

			*before = *row
		}
	} else {
		*before = *pokemon
	}

	dao := persistence.NewPokemonDAO(u.db)

	err := dao.ScanEvolution(before)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// when not evolution
	if before.Evolution == nil {
		return datasets, nil
	}

	// forward
	dataset, err := u.constructPokemonDataset(before)
	if err != nil {
		return nil, err
	}

	datasets = append(datasets, dataset)

	for {
		err := dao.ScanEvolution(before)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		if before.Evolution == nil {
			break
		}

		dataset, err := u.constructPokemonDataset(before.Evolution)
		if err != nil {
			return nil, err
		}

		datasets = append(datasets, dataset)

		before = before.Evolution
	}

	return datasets, nil
}

func (u *PokemonDatasetAcquisition) constructPokemonDataset(pokemon *model.Pokemon) (*PokemonDataset, error) {
	dao := persistence.NewPokemonDAO(u.db)

	if err := dao.ScanEvolution(pokemon); err != nil {
		return nil, err
	}

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

	link := &LinkInfo{
		PrevNationalNo: pokemon.NationalNo - 1,
		NextNationalNo: pokemon.NationalNo + 1,
	}

	var r *gorm.DB

	r = u.db.Model(&model.Pokemon{}).Where("national_no = ?", link.PrevNationalNo).First(&model.Pokemon{})
	if r.Error != nil {
		if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, r.Error
		}
	}

	link.HasPrev = r.RowsAffected > 0

	r = u.db.Model(&model.Pokemon{}).Where("national_no = ?", link.NextNationalNo).First(&model.Pokemon{})
	if r.Error != nil {
		if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, r.Error
		}
	}

	link.HasNext = r.RowsAffected > 0

	canEvolution := pokemon.Evolution != nil

	return &PokemonDataset{
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
		LinkInfo:        link,
		CanEvolution:    canEvolution,
	}, nil
}

type PokemonDataset struct {
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
	LinkInfo        *LinkInfo
	CanEvolution    bool
	Evolutions      []*PokemonDataset
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

type LinkInfo struct {
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
