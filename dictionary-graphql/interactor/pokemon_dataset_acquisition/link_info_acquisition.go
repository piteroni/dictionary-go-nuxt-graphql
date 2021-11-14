package pokemon_dataset_acquisition

import (
	"piteroni/dictionary-go-nuxt-graphql/model"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

var _ iLinkInfoAcquisition = (*linkInfoAcquisition)(nil)

// that provides the link info.
type iLinkInfoAcquisition interface {
	getLinkInfo(pokemon *model.Pokemon) (*LinkInfo, error)
}

type linkInfoAcquisition struct {
	db *gorm.DB
}

func (i *linkInfoAcquisition) getLinkInfo(pokemon *model.Pokemon) (*LinkInfo, error) {
	link := &LinkInfo{
		PrevID: int(pokemon.ID - 1),
		NextID: int(pokemon.ID + 1),
	}

	var tx *gorm.DB

	tx = i.db.Model(&model.Pokemon{}).Where("id = ?", link.PrevID).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(tx.Error)
		}
	}

	link.HasPrev = tx.RowsAffected > 0

	tx = i.db.Model(&model.Pokemon{}).Where("id = ?", link.NextID).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(tx.Error)
		}
	}

	link.HasNext = tx.RowsAffected > 0

	return link, nil
}
