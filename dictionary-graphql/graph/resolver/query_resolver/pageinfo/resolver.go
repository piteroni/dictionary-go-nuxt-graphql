package pageinfo

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"
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
	i, err := r.getPageInfo(pokemonID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return graph.PokemonNotFound{}, nil
		}

		return nil, err
	}

	return i, nil
}

func (r *PageInfoQueryResolver) getPageInfo(pokemonID int) (graph.PageInfoResult, error) {
	pokemon := &model.Pokemon{}

	err := r.DB.Model(&model.Pokemon{}).Find(pokemon, pokemonID).Error
	if err != nil {
		// return nil, err
		return nil, errors.WithStack(err)
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
