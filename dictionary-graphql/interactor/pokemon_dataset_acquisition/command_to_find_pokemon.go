package pokemon_dataset_acquisition

import (
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/model"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

var _ iCommandToFindPokemon = (*commandToFindPokemon)(nil)

type iCommandToFindPokemon interface {
	execute(pokemonID int) (*model.Pokemon, error)
}

type commandToFindPokemon struct {
	db *gorm.DB
}

func (c *commandToFindPokemon) execute(pokemonID int) (*model.Pokemon, error) {
	pokemon := &model.Pokemon{}

	err := c.db.Model(&model.Pokemon{}).First(pokemon, pokemonID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(&PokemonNotFound{
				message: fmt.Sprintf("specified pokemon does not exists, pokemonID = %d", pokemonID),
			})
		}

		return nil, errors.WithStack(err)
	}

	return pokemon, nil
}
