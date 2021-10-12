package main

import (
	"os"
	"piteroni/dictionary-go-nuxt-graphql/pkg/driver"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := driver.NewLogger(os.Stdout)

	if err := godotenv.Load(); err != nil {
		logger.Errorf("unexpected error occurred during loading .env: %v", err)
		os.Exit(1)
	}

	_, err := driver.ConnectToDatabase()
	if err != nil {
		logger.Errorf("unexpected error occurred during connect database: %v", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
