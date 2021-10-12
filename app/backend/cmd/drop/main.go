package main

import (
	"os"
	"piteroni/dictionary-go-nuxt-graphql/pkg/driver"

	"github.com/joho/godotenv"
)

const (
	statusError = 1
	statusFatal = 2
)

func main() {
	logger := driver.NewLogger(os.Stdout)

	if err := godotenv.Load(); err != nil {
		logger.Errorf("unexpected error occurred during loading .env: %v", err)
		os.Exit(statusFatal)
	}

	db, err := driver.ConnectToDatabase()
	if err != nil {
		logger.Errorf("unexpected error occurred during connect database: %v", err)
		os.Exit(statusFatal)
	}

	if err := driver.Drop(db); err != nil {
		logger.Error(err)
		os.Exit(statusError)
	}
}
