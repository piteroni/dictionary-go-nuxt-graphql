package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() (*gorm.DB, error) {
	username, err := env("DB_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := env("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	host, err := env("DB_HOST")
	if err != nil {
		return nil, err
	}

	port, err := env("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbname, err := env("DB_NAME")
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

func env(key string) (string, error) {
	message := "environment variables for access aws are not set: %s"

	value, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf(message, key))
	}

	return value, nil
}
