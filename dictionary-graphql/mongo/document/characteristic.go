package document

import (
	"time"
)

type Characteristic struct {
	Record
	Name        string `bson:"name"`
	Description string `bson:"description"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
