package pokemon_dataset_acquisition

import (
	"piteroni/dictionary-go-nuxt-graphql/datasource/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonDatasetAcquisition(t *testing.T) {
	t.Run("指定したIDに一致するポケモンのデータセットを取得できる", func(t *testing.T) {
		datasetAcquisition := &pokemonDatasetAcquisition{
			commandToFindPokemon:      &commandToFindPokemonMock{t: t},
			basicInfoAcquisition:      &basicInfoAcquisitionMock{t: t},
			linkInfoAcquisition:       &linkInfoAcquisitionMock{t: t},
			evolutionTableAcquisition: &evolutionTableAcquisitionMock{t: t},
		}

		dataset, err := datasetAcquisition.GetPokemonDataset(201)

		assert.NotNil(t, dataset)
		assert.Nil(t, err)

		// basic information of pokemon is set.
		assert.Equal(t, dataset.NationalNo, 201)
		assert.Equal(t, dataset.Name, "pokemon-201")

		// link information is set.
		assert.NotNil(t, dataset.LinkInfo)
		assert.Equal(t, dataset.LinkInfo, &LinkInfo{
			PrevNationalNo: 200,
			NextNationalNo: 202,
			HasPrev:        true,
			HasNext:        true,
		})

		// evolution table of pokemon is set.
		assert.Len(t, dataset.Evolutions, 3)
		assert.Equal(t, dataset.Evolutions[0].NationalNo, 201)
		assert.Equal(t, dataset.Evolutions[0].Name, "pokemon-201")
		assert.Equal(t, dataset.Evolutions[1].NationalNo, 202)
		assert.Equal(t, dataset.Evolutions[1].Name, "pokemon-202")
		assert.Equal(t, dataset.Evolutions[2].NationalNo, 203)
		assert.Equal(t, dataset.Evolutions[2].Name, "pokemon-203")
	})
}

var _ iCommandToFindPokemon = (*commandToFindPokemonMock)(nil)

type commandToFindPokemonMock struct {
	t *testing.T
}

func (m *commandToFindPokemonMock) execute(pokemonID int) (*model.Pokemon, error) {
	assert.Equal(m.t, pokemonID, 201)

	return &model.Pokemon{
		NationalNo: 201,
		Name:       "pokemon-201",
	}, nil
}

var _ iBasicInfoAcquisition = (*basicInfoAcquisitionMock)(nil)

type basicInfoAcquisitionMock struct {
	t *testing.T
}

var _ iBasicInfoAcquisition = (*basicInfoAcquisitionMock)(nil)

func (m *basicInfoAcquisitionMock) getBasicInfo(pokemon *model.Pokemon) (*PokemonDataset, error) {
	assert.Equal(m.t, pokemon.NationalNo, 201)
	assert.Equal(m.t, pokemon.Name, "pokemon-201")

	return &PokemonDataset{
		NationalNo: 201,
		Name:       "pokemon-201",
	}, nil
}

var _ iLinkInfoAcquisition = (*linkInfoAcquisitionMock)(nil)

type linkInfoAcquisitionMock struct {
	t *testing.T
}

func (m *linkInfoAcquisitionMock) getLinkInfo(pokemon *model.Pokemon) (*LinkInfo, error) {
	assert.Equal(m.t, pokemon.NationalNo, 201)
	assert.Equal(m.t, pokemon.Name, "pokemon-201")

	return &LinkInfo{
		PrevNationalNo: 200,
		NextNationalNo: 202,
		HasPrev:        true,
		HasNext:        true,
	}, nil
}

var _ iEvolutionTableAcquisition = (*evolutionTableAcquisitionMock)(nil)

type evolutionTableAcquisitionMock struct {
	t *testing.T
}

func (m *evolutionTableAcquisitionMock) getEvolutionTable(pokemon *model.Pokemon) ([]*PokemonDataset, error) {
	assert.Equal(m.t, pokemon.NationalNo, 201)
	assert.Equal(m.t, pokemon.Name, "pokemon-201")

	return []*PokemonDataset{
		{
			NationalNo: pokemon.NationalNo,
			Name:       pokemon.Name,
		},
		{
			NationalNo: 202,
			Name:       "pokemon-202",
		},
		{
			NationalNo: 203,
			Name:       "pokemon-203",
		},
	}, nil
}
