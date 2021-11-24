package pageinfo

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
	graph_internal "piteroni/dictionary-go-nuxt-graphql/graph/internal"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/model"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type PageInfoQueryResolver struct {
	DB     *gorm.DB
	Logger *driver.AppLogger
}

func (r *PageInfoQueryResolver) PageInfo(pokemonID int) (graph.PageInfoResult, error) {
	pokemon := &model.Pokemon{}

	err := r.DB.Model(&model.Pokemon{}).Find(pokemon, pokemonID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return graph.PokemonNotFound{}, nil
		}

		r.Logger.Error(err)

		return nil, graph_internal.InternalSystemError
	}

	i := graph.PageInfo{
		PrevID: int(pokemon.ID - 1),
		NextID: int(pokemon.ID + 1),
	}

	var tx *gorm.DB

	tx = r.DB.Model(&model.Pokemon{}).Where("id = ?", i.PrevID).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(tx.Error)
		}
	}

	i.HasPrev = tx.RowsAffected > 0

	tx = r.DB.Model(&model.Pokemon{}).Where("id = ?", i.NextID).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(tx.Error)
		}
	}

	i.HasNext = tx.RowsAffected > 0

	return i, nil
}
