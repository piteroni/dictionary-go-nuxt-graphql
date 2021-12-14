package pokemons

// import (
// 	"io/ioutil"
// 	"testing"

// 	"piteroni/dictionary-go-nuxt-graphql/driver"
// 	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
// 	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
// 	"piteroni/dictionary-go-nuxt-graphql/model"
// 	itesting "piteroni/dictionary-go-nuxt-graphql/testing"

// 	"gorm.io/gorm"

// 	"github.com/stretchr/testify/assert"
// )

// func TestPokemonsQueryResolver(t *testing.T) {
// 	t.Run("ポケモンのIDとそれに続く範囲を指定することで、複数のポケモンのデータを取得できる", func(t *testing.T) {
// 		r := &PokemonsQueryResolver{
// 			GraphQLModelMapper:   &pokemon_interactor.GraphQLModelMapper{},
// 			PokemonSearchCommand: &pokemonSearchCommandMock{t: t},
// 		}

// 		expected := graph.PokemonConnection{
// 			Items: []*graph.Pokemon{
// 				{
// 					ID:          200,
// 					NationalNo:  200,
// 					Name:        "pokemon-200",
// 					Ability:     &graph.Ability{},
// 					Description: &graph.Description{},
// 				},
// 				{
// 					ID:          201,
// 					NationalNo:  201,
// 					Name:        "pokemon-201",
// 					Ability:     &graph.Ability{},
// 					Description: &graph.Description{},
// 				},
// 				{
// 					ID:          202,
// 					NationalNo:  202,
// 					Name:        "pokemon-202",
// 					Ability:     &graph.Ability{},
// 					Description: &graph.Description{},
// 				},
// 			},
// 			NextID: 203,
// 		}

// 		actual, err := r.Pokemons(itesting.Int(3), itesting.Int(200))

// 		assert.NotNil(t, actual)
// 		assert.Nil(t, err)

// 		assert.Equal(t, expected, actual)
// 	})

// 	t.Run("指f定したIDに一致するポケモンが存在しない場合、例外が返る", func(t *testing.T) {
// 		logger := driver.NewLogger(ioutil.Discard)

// 		r := &PokemonsQueryResolver{
// 			GraphQLModelMapper:   &pokemon_interactor.GraphQLModelMapper{},
// 			PokemonSearchCommand: &pokemonSearchCommandMockWhenNotFound{t: t},
// 			AppLogger:            logger,
// 		}

// 		expected := graph.PokemonNotFound{}

// 		actual, err := r.Pokemons(itesting.Int(2), itesting.Int(203))

// 		assert.NotNil(t, actual)
// 		assert.Nil(t, err)

// 		assert.Equal(t, expected, actual)
// 	})

// 	t.Run("パラメーターの値が不正な場合、例外が返る", func(t *testing.T) {
// 		logger := driver.NewLogger(ioutil.Discard)

// 		r := &PokemonsQueryResolver{
// 			GraphQLModelMapper:   &pokemon_interactor.GraphQLModelMapper{},
// 			PokemonSearchCommand: &PokemonSearchCommandMockWhenIlligalArguments{t: t},
// 			AppLogger:            logger,
// 		}

// 		expected := graph.IllegalArguments{}

// 		actual, err := r.Pokemons(itesting.Int(2), itesting.Int(203))

// 		assert.NotNil(t, actual)
// 		assert.Nil(t, err)

// 		assert.Equal(t, expected, actual)
// 	})
// }

// type pokemonSearchCommandMock struct{ t *testing.T }

// var _ pokemon_interactor.PokemonSearchCommand = (*pokemonSearchCommandMock)(nil)

// func (m *pokemonSearchCommandMock) Execute(first *int, after *int) ([]*model.Pokemon, error) {
// 	assert.Equal(m.t, *first, 3)
// 	assert.Equal(m.t, *after, 200)

// 	return []*model.Pokemon{
// 		{
// 			Model:      gorm.Model{ID: 200},
// 			Name:       "pokemon-200",
// 			NationalNo: 200,
// 		},
// 		{
// 			Model:      gorm.Model{ID: 201},
// 			Name:       "pokemon-201",
// 			NationalNo: 201,
// 		},
// 		{
// 			Model:      gorm.Model{ID: 202},
// 			Name:       "pokemon-202",
// 			NationalNo: 202,
// 		},
// 	}, nil
// }

// type pokemonSearchCommandMockWhenNotFound struct{ t *testing.T }

// var _ pokemon_interactor.PokemonSearchCommand = (*pokemonSearchCommandMockWhenNotFound)(nil)

// func (m *pokemonSearchCommandMockWhenNotFound) Execute(first *int, after *int) ([]*model.Pokemon, error) {
// 	assert.Equal(m.t, *first, 2)
// 	assert.Equal(m.t, *after, 203)

// 	return nil, &pokemon_interactor.PokemonNotFound{}
// }

// type PokemonSearchCommandMockWhenIlligalArguments struct{ t *testing.T }

// var _ pokemon_interactor.PokemonSearchCommand = (*PokemonSearchCommandMockWhenIlligalArguments)(nil)

// func (m *PokemonSearchCommandMockWhenIlligalArguments) Execute(first *int, after *int) ([]*model.Pokemon, error) {
// 	assert.Equal(m.t, *first, 2)
// 	assert.Equal(m.t, *after, 203)

// 	return nil, &pokemon_interactor.IllegalArguments{}
// }
