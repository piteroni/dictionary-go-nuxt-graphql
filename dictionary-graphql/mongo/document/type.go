package document

import (
	"time"
)

type Type struct {
	Name    string `bson:"name"`
	IconURL string `bson:"icon_url"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
