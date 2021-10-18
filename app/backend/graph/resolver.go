package graph

import (
	"piteroni/dictionary-go-nuxt-graphql/pkg/drivers"

	"gorm.io/gorm"
)

type Resolver struct {
	DB     *gorm.DB
	Logger drivers.AppLogger
}
