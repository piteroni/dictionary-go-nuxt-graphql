package document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Description struct {
	ID     primitive.ObjectID `bson:"_id"`
	Text   string             `bson:"text"`
	Series string             `bson:"series"`
}
