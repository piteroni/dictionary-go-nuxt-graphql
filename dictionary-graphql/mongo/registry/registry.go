package registry

import (
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

func NewRegistry() *bsoncodec.Registry {
	rb := bsoncodec.NewRegistryBuilder()

	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bson.PrimitiveCodecs{}.RegisterPrimitiveCodecs(rb)

	rb.RegisterTypeEncoder(reflect.TypeOf(document.Record{}), &RecordEncoder{})

	return rb.Build()
}
