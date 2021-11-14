package graph

import (
	"piteroni/dictionary-go-nuxt-graphql/driver"

	"gorm.io/gorm"
)

type Resolver struct {
	DB     *gorm.DB
	Logger *driver.AppLogger
}
