package testutils

import (
	"context"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() (string, string, error) {
	username, err := driver.Env("DB_USERNAME")
	if err != nil {
		return "", "", err
	}

	password, err := driver.Env("DB_PASSWORD")
	if err != nil {
		return "", "", err
	}

	host, err := driver.Env("DB_HOST")
	if err != nil {
		return "", "", err
	}

	port, err := driver.Env("DB_PORT")
	if err != nil {
		return "", "", err
	}

	dbname, err := driver.Env("TEST_DB_NAME")
	if err != nil {
		return "", "", err
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	return uri, dbname, nil
}

func ConnnectToTestDatabase() (*mongo.Database, func() error, error) {
	uri, dbname, err := DatabaseConnection()
	if err != nil {
		return nil, nil, err
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
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

func ObjectID(t *testing.T, value string) primitive.ObjectID {
	objectId, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		t.Fatal(err)
	}

	return objectId
}
