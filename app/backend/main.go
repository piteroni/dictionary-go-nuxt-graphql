package main

import (
	"os"
	"piteroni/dictionary-go-nuxt-graphql/pkg/database"
	"piteroni/dictionary-go-nuxt-graphql/pkg/drivers"
	"piteroni/dictionary-go-nuxt-graphql/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := drivers.NewLogger(os.Stdout)

	if err := godotenv.Load(); err != nil {
		logger.Errorf("unexpected error occurred during loading .env: %v", err)
		os.Exit(1)
	}

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Errorf("unexpected error occurred during connect database: %v", err)
		os.Exit(1)
	}

	e := gin.Default()

	routes.InitAPIRouting(e, db)

	e.Run(":8080")
}
