package migration

import (
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := []interface{}{
		&models.Type{},
		&models.Gender{},
		&models.Characteristic{},
		&models.Description{},
		&models.Pokemon{},
	}

	for _, model := range m {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}

func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.Type{},
		&models.Gender{},
		&models.Characteristic{},
		&models.Description{},
		&models.Pokemon{},
		"type_possessed",
		"characteristic_possessed",
		"gender_possessed",
	)
}
