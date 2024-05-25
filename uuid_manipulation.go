package uuid_decoder

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
)

func UuidToStringManipulation(in interface{}) (interface{}, error) {
	// get value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	var id uuid.UUID

	if inVal.Kind() == reflect.ValueOf(id).Kind() {
		id = in.(uuid.UUID)
		return id.String(), nil
	}

	return nil, errors.New("invalid type of uuid when trying to manipulate input")
}
