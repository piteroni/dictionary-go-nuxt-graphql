package main

import (
	"os"
	"piteroni/dictionary-go-nuxt-graphql/cmd/internal"
	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/database/migration"
	"piteroni/dictionary-go-nuxt-graphql/driver"
)

func main() {
	logger := driver.NewLogger(os.Stdout)

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Errorf("unexpected error occurred during connect database: %v", err)
		os.Exit(internal.StatusFatal)
	}

	err = migration.DropTables(db)
	if err != nil {
		logger.Error(err)
		os.Exit(internal.StatusError)
	}
}
