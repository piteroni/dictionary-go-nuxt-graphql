package main

import (
	"log"
	"os"
	"piteroni/dictionary-go-nuxt-graphql/database"

	"github.com/joho/godotenv"
)

const (
	statusError = 1
	statusFatal = 2
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Ldate|log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		logger.Fatalf("unexpected error occurred during loading .env: %v", err)
	}

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Fatalf("unexpected error occurred during connect database: %v", err)
		os.Exit(statusFatal)
	}

	err = database.Migrate(db)
	if err != nil {
		logger.Fatal(err)
	}

	err = database.Seed(db)
	if err != nil {
		logger.Fatal(err)
	}
}
