package pokemon_dataset_acquisition

import (
	"errors"
	"piteroni/dictionary-go-nuxt-graphql/datasource/model"

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
		PrevNationalNo: pokemon.NationalNo - 1,
		NextNationalNo: pokemon.NationalNo + 1,
	}

	var tx *gorm.DB

	tx = i.db.Model(&model.Pokemon{}).Where("national_no = ?", link.PrevNationalNo).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, tx.Error
		}
	}

	link.HasPrev = tx.RowsAffected > 0

	tx = i.db.Model(&model.Pokemon{}).Where("national_no = ?", link.NextNationalNo).First(&model.Pokemon{})
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, tx.Error
		}
	}

	link.HasNext = tx.RowsAffected > 0

	return link, nil
}
