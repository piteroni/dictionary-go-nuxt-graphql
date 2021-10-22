package main

import (
	"os"
	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/database/migration"
	"piteroni/dictionary-go-nuxt-graphql/driver"

	"github.com/joho/godotenv"
)

const (
	statusError = 1
	statusFatal = 2
)

func main() {
	logger := driver.NewLogger(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		logger.Errorf("unexpected error occurred during loading .env: %v", err)
		os.Exit(statusFatal)
	}

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Errorf("unexpected error occurred during connect database: %v", err)
		os.Exit(statusFatal)
	}

	err = migration.Migrate(db)
	if err != nil {
		logger.Error(err)
		os.Exit(statusError)
	}

	err = migration.Seed(db)
	if err != nil {
		logger.Error(err)
		os.Exit(statusError)
	}
}
