package database

import (
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/driver"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() (*gorm.DB, error) {
	username, err := driver.Env("DB_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := driver.Env("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	host, err := driver.Env("DB_HOST")
	if err != nil {
		return nil, err
	}

	port, err := driver.Env("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbname, err := driver.Env("DB_NAME")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
