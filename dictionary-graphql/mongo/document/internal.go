package document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID primitive.ObjectID `bson:"_id"`
}
