package registry

import (
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"reflect"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ bsoncodec.ValueEncoder = (*RecordEncoder)(nil)

type RecordEncoder struct{}

func (rc *RecordEncoder) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	t := reflect.TypeOf(document.Record{})

	if val.Type() != t {
		err := bsoncodec.ValueEncoderError{
			Name:     "RecordEncodeValue",
			Kinds:    []reflect.Kind{t.Kind()},
			Received: val,
		}

		return errors.Cause(err)
	}

	w, err := vw.WriteDocument()
	if err != nil {
		return errors.WithStack(err)
	}

	if !reflect.DeepEqual(val.FieldByName("ID"), primitive.NilObjectID) {
		ew, err := w.WriteDocumentElement("_id")
		if err != nil {
			return errors.WithStack(err)
		}

		ew.WriteObjectID(primitive.NewObjectID())
	}

	w.WriteDocumentEnd()

	return nil
}
