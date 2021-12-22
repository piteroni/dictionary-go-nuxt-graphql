package graph

import (
	"context"
	"piteroni/dictionary-go-nuxt-graphql/driver"

	"go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
	DB      *mongo.Database
	Logger  *driver.AppLogger
	Context context.Context
}
