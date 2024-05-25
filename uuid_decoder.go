package uuid_decoder

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
)

func UuidDecoder(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	// set value if input and output have same type
	if inVal.Kind() == out.Kind() {
		out.Set(inVal)
		return
	}

	switch inVal.Kind() {
	case reflect.String:
		if res, err := uuid.Parse(inVal.String()); err == nil {
			out.Set(reflect.ValueOf(res))
		}
	default:
		err = fmt.Errorf("got %s type, expected %s", inVal.Kind(), out.Type())
	}

	return
}
