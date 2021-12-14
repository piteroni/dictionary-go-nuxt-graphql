package registry

import (
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"
	"reflect"
	"time"

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

	if val.FieldByName("ID").IsNil() {
		ew, err := w.WriteDocumentElement("_id")
		if err != nil {
			return errors.WithStack(err)
		}

		ew.WriteObjectID(primitive.NewObjectID())
	}

	if val.FieldByName("CreatedAt").FieldByName("wall").Uint() == 0 {
		ew, err := w.WriteDocumentElement("created_at")
		if err != nil {
			return errors.WithStack(err)
		}

		ew.WriteDateTime(time.Now().Unix())
	}

	if val.FieldByName("UpdatedAt").FieldByName("wall").Uint() == 0 {
		ew, err := w.WriteDocumentElement("updated_at")
		if err != nil {
			return errors.WithStack(err)
		}

		ew.WriteDateTime(time.Now().Unix())
	}

	w.WriteDocumentEnd()

	return nil
}
