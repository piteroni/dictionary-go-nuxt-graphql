package database

import (
	"piteroni/dictionary-go-nuxt-graphql/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := []interface{}{
		&model.Type{},
		&model.Gender{},
		&model.Characteristic{},
		&model.Description{},
		&model.Pokemon{},
	}

	for _, model := range m {
		err := db.AutoMigrate(model)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&model.Type{},
		&model.Gender{},
		&model.Characteristic{},
		&model.Description{},
		&model.Pokemon{},
		"pokemon_types",
		"pokemon_characteristics",
		"pokemon_genders",
	)
}
