package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Drop(ctx context.Context, db *mongo.Database) error {
	c, err := db.ListCollectionNames(ctx, bson.D{{}})
	if err != nil {
		return err
	}

	for _, v := range c {
		err := db.Collection(v).Drop(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
