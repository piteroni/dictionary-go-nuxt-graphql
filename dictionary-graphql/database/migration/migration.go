package migration

import (
	"piteroni/dictionary-go-nuxt-graphql/model"

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
			return err
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
		"type_possessed",
		"characteristic_possessed",
		"gender_possessed",
	)
}
