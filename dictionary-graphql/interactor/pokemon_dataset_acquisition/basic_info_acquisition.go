package pokemon_dataset_acquisition

import (
	"piteroni/dictionary-go-nuxt-graphql/model"
	"piteroni/dictionary-go-nuxt-graphql/persistence"

	"gorm.io/gorm"
)

var _ iBasicInfoAcquisition = (*basicInfoAcquisition)(nil)

// that provides the pokemon basic info.
type iBasicInfoAcquisition interface {
	getBasicInfo(pokemon *model.Pokemon) (*PokemonDataset, error)
}

type basicInfoAcquisition struct {
	db *gorm.DB
}

func (i *basicInfoAcquisition) getBasicInfo(pokemon *model.Pokemon) (*PokemonDataset, error) {
	dao := persistence.NewPokemonDAO(i.db)

	err := dao.ScanEvolution(pokemon)
	if err != nil {
		return nil, err
	}

	err = dao.ScanGenders(pokemon)
	if err != nil {
		return nil, err
	}

	err = dao.ScanTypes(pokemon)
	if err != nil {
		return nil, err
	}

	err = dao.ScanCharacteristics(pokemon)
	if err != nil {
		return nil, err
	}

	err = dao.ScanDescriptions(pokemon)
	if err != nil {
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
		SpecialAttack:  pokemon.SpecialAttackPoint,
		SpecialDefense: pokemon.SpecialDefensePoint,
		Speed:          pokemon.SpeedPoint,
	}

	canEvolution := pokemon.Evolution != nil

	return &PokemonDataset{
		ID:              int(pokemon.ID),
		NationalNo:      pokemon.NationalNo,
		Name:            pokemon.Name,
		ImageURL:        pokemon.ImageURL,
		Genders:         genders,
		Species:         pokemon.Species,
		Types:           types,
		Height:          pokemon.Height,
		Weight:          pokemon.Weight,
		Characteristics: characteristics,
		Description:     description,
		Ability:         ability,
		CanEvolution:    canEvolution,
	}, nil
}
