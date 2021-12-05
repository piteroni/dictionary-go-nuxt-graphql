package pokemon

import (
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"testing"

	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

func TestPokemonQueryResolver(t *testing.T) {
	t.Run("指定したIDに一致するポケモンのデータセットを取得できる", func(t *testing.T) {
		r := &PokemonQueryResolver{
			GraphQLModelMapper:   &pokemon_interactor.GraphQLModelMapper{},
			PokemonSearchCommand: &pokemonSearchCommandMock{t: t},
		}

		expected := &graph.Pokemon{
			ID:          100,
			NationalNo:  100,
			Name:        "pokemon-100",
			Ability:     &graph.Ability{},
			Description: &graph.Description{},
		}

		actual, err := r.Pokemon(100)

		assert.NotNil(t, actual)
		assert.Nil(t, err)

		assert.Equal(t, expected, actual)
	})

	t.Run("指定したIDに一致するポケモンが存在しない場合、例外が返る", func(t *testing.T) {
		r := &PokemonQueryResolver{
			GraphQLModelMapper:   &pokemon_interactor.GraphQLModelMapper{},
			PokemonSearchCommand: &pokemonSearchCommandMockWhenNotFound{t: t},
		}

		expected := &graph.PokemonNotFound{}

		actual, err := r.Pokemon(101)

		assert.NotNil(t, actual)
		assert.Nil(t, err)

		assert.Equal(t, expected, actual)
	})
}

var _ pokemon_interactor.PokemonSearchCommand = (*pokemonSearchCommandMock)(nil)

type pokemonSearchCommandMock struct{ t *testing.T }

func (m *pokemonSearchCommandMock) Execute(first *int, after *int) ([]*model.Pokemon, error) {
	assert.Equal(m.t, *first, 0)
	assert.Equal(m.t, *after, 100)

	return []*model.Pokemon{
		{
			Model:      gorm.Model{ID: 100},
			Name:       "pokemon-100",
			NationalNo: 100,
		},
	}, nil
}

type pokemonSearchCommandMockWhenNotFound struct{ t *testing.T }

var _ pokemon_interactor.PokemonSearchCommand = (*pokemonSearchCommandMock)(nil)

func (m *pokemonSearchCommandMockWhenNotFound) Execute(first *int, after *int) ([]*model.Pokemon, error) {
	assert.Equal(m.t, *first, 0)
	assert.Equal(m.t, *after, 101)

	return nil, &pokemon_interactor.PokemonNotFound{}
}
