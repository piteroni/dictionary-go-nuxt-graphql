package testutils

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IntPtr(v int) *int {
	return &v
}

func StringPtr(v string) *string {
	return &v
}

func UIntptr(v uint) *uint {
	return &v
}

func BytesPtr(v []byte) *[]byte {
	return &v
}

func ObjectIDPtr(t *testing.T, value primitive.ObjectID) *primitive.ObjectID {
	return &value
}
