package main

import (
	"context"
	"os"
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/mongo/database"

	"github.com/pkg/errors"
)

func main() {
	logger := driver.NewLogger(os.Stdout)

	db, close, err := database.Connect()
	if err != nil {
		logger.Error(errors.WithStack(err))
		return
	}

	defer func() {
		err := close()
		if err != nil {
			logger.Error(errors.WithStack(err))
			return
		}
	}()

	err = database.Seed(context.Background(), db)
	if err != nil {
		logger.Error(errors.WithStack(err))
		return
	}
}
