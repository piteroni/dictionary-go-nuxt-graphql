package pageinfo

import (
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/model"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type PageInfoQueryResolver struct {
	DB *gorm.DB
}

func (r *PageInfoQueryResolver) PageInfo(pokemonID int) (graph.PageInfoResult, error) {
	pokemon := &model.Pokemon{}

	tx := r.DB.Model(&model.Pokemon{}).Find(pokemon, pokemonID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected <= 0 {
		return graph.PokemonNotFound{}, nil
	}

	i := graph.PageInfo{
		PrevID: int(pokemon.ID - 1),
		NextID: int(pokemon.ID + 1),
	}

	tx = &gorm.DB{}

	tx = r.DB.Model(&model.Pokemon{}).Where("id = ?", i.PrevID).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, tx.Error
		}
	}

	i.HasPrev = tx.RowsAffected > 0

	tx = r.DB.Model(&model.Pokemon{}).Where("id = ?", i.NextID).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, tx.Error
		}
	}

	i.HasNext = tx.RowsAffected > 0

	return i, nil
}
