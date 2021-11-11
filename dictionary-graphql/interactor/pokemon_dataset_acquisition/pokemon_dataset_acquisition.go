package pokemon_dataset_acquisition

import (
	"gorm.io/gorm"
)

// Facade object that provides the pokemon dataset.
type pokemonDatasetAcquisition struct {
	commandToFindPokemon      iCommandToFindPokemon
	basicInfoAcquisition      iBasicInfoAcquisition
	linkInfoAcquisition       iLinkInfoAcquisition
	evolutionTableAcquisition iEvolutionTableAcquisition
}

func New(db *gorm.DB) *pokemonDatasetAcquisition {
	commandToFindPokemon := &commandToFindPokemon{
		db: db,
	}
	basicInfoAcquisition := &basicInfoAcquisition{
		db: db,
	}
	linkInfoAcquisition := &linkInfoAcquisition{
		db: db,
	}
	evolutionTableAcquisition := &evolutionTableAcquisition{
		db:                   db,
		basicInfoAcquisition: basicInfoAcquisition,
	}

	return &pokemonDatasetAcquisition{
		commandToFindPokemon:      commandToFindPokemon,
		basicInfoAcquisition:      basicInfoAcquisition,
		linkInfoAcquisition:       linkInfoAcquisition,
		evolutionTableAcquisition: evolutionTableAcquisition,
	}
}

func (u *pokemonDatasetAcquisition) GetPokemonDataset(pokemonID int) (*PokemonDataset, error) {
	// memo: pokemon Modelは以降でも使い回すから、なるだけstructは下位のInteractorクラスには値渡しで渡したいけど結構むずそう.
	// go言語でdeepcopyはかなりめんどくさそう、gormの仕様をきっちり調べて何が入力としてあればいいのか把握して、下位のInteractorオブジェクトごとにinputオブジェクトを用意する必要がありそう
	pokemon, err := u.commandToFindPokemon.execute(pokemonID)
	if err != nil {
		return nil, err
	}

	dataset, err := u.basicInfoAcquisition.getBasicInfo(pokemon)
	if err != nil {
		return nil, err
	}

	dataset.LinkInfo, err = u.linkInfoAcquisition.getLinkInfo(pokemon)
	if err != nil {
		return nil, err
	}

	dataset.Evolutions, err = u.evolutionTableAcquisition.getEvolutionTable(pokemon)

	return dataset, nil
}

type PokemonDataset struct {
	NationalNo      int
	Name            string
	ImageURL        string
	Description     *Description
	Species         string
	Types           []*Type
	EvolutionID     *uint
	Height          string
	Weight          string
	Genders         []*Gender
	Ability         *Ability
	Characteristics []*Characteristic
	LinkInfo        *LinkInfo
	Evolutions      []*PokemonDataset
	CanEvolution    bool
}

type Description struct {
	Text   string
	Series string
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

type Ability struct {
	Heart          int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

type LinkInfo struct {
	PrevID  int
	NextID  int
	HasPrev bool
	HasNext bool
}

var _ error = (*PokemonNotFound)(nil)

type PokemonNotFound struct {
	message string
}

func (e PokemonNotFound) Error() string {
	return e.message
}
