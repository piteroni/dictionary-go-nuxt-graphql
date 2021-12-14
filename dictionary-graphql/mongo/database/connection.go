package database

import (
	"context"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/mongo/registry"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, func() error, error) {
	username, err := driver.Env("DB_USERNAME")
	if err != nil {
		return nil, nil, err
	}

	password, err := driver.Env("DB_PASSWORD")
	if err != nil {
		return nil, nil, err
	}

	host, err := driver.Env("DB_HOST")
	if err != nil {
		return nil, nil, err
	}

	port, err := driver.Env("DB_PORT")
	if err != nil {
		return nil, nil, err
	}

	dbname, err := driver.Env("DB_NAME")
	if err != nil {
		return nil, nil, err
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	registry := registry.NewRegistry()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetRegistry(registry))
	if err != nil {
		return nil, nil, err
	}

	close := func() error {
		if err := client.Disconnect(context.TODO()); err != nil {
			return err
		}

		return nil
	}

	return client.Database(dbname), close, nil
}
