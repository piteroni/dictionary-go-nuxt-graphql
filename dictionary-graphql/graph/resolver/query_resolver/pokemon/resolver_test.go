package pokemon

import (
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	pokemon_interactor "piteroni/dictionary-go-nuxt-graphql/graph/resolver/query_resolver/pokemon.interactor"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"piteroni/dictionary-go-nuxt-graphql/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonQueryResolver(t *testing.T) {
	t.Run("指定したIDに一致するポケモンのデータセットを取得できる", func(t *testing.T) {
		r := &PokemonQueryResolver{
			GraphQLModelMapper:   &pokemon_interactor.GraphQLModelMapper{},
			PokemonSearchCommand: &pokemonSearchCommandMock{t: t},
		}

		expected := &graph.Pokemon{
			ID:          "000000000000000000000100",
			Ability:     &graph.Ability{},
			Description: &graph.Description{},
		}

		actual, err := r.Pokemon("000000000000000000000100")

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

		actual, err := r.Pokemon("000000000000000000000101")

		assert.NotNil(t, actual)
		assert.Nil(t, err)

		assert.Equal(t, expected, actual)
	})
}

type pokemonSearchCommandMock struct{ t *testing.T }

var _ pokemon_interactor.PokemonSearchCommand = (*pokemonSearchCommandMock)(nil)

func (m *pokemonSearchCommandMock) Execute(first *int, after *string) ([]*document.Pokemon, error) {
	assert.Equal(m.t, *first, 0)
	assert.Equal(m.t, *after, "000000000000000000000100")

	pokemons := []*document.Pokemon{
		{
			ID: testutils.ObjectID(m.t, "000000000000000000000100"),
		},
	}

	return pokemons, nil
}

type pokemonSearchCommandMockWhenNotFound struct{ t *testing.T }

var _ pokemon_interactor.PokemonSearchCommand = (*pokemonSearchCommandMockWhenNotFound)(nil)

func (m *pokemonSearchCommandMockWhenNotFound) Execute(first *int, after *string) ([]*document.Pokemon, error) {
	assert.Equal(m.t, *first, 0)
	assert.Equal(m.t, *after, "000000000000000000000101")

	return []*document.Pokemon{}, nil
}
