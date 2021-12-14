package document

import (
	"time"
)

type Description struct {
	Text   string `bson:"text"`
	Series string `bson:"series"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
