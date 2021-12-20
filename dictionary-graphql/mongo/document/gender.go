package document

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gender struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	IconURL string             `bson:"icon_url"`
}
