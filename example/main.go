package main

import (
	"github.com/evenyosua18/typedecoder"
	uuid_decoder "github.com/evenyosua18/uuid-decoder"
	"github.com/google/uuid"
	"log"
	"reflect"
)

// implement uuid decoder into type decoder package
func main() {
	// add uuid decoder
	var uuidDecoder uuid.UUID
	typedecoder.AddDecodeFunction(reflect.TypeOf(uuidDecoder).String(), uuid_decoder.UuidDecoder)

	// init data
	idString := uuid.New().String()
	var result uuid.UUID

	// try to decode
	if err := typedecoder.Decode(idString, &result); err != nil {
		panic(err)
	}

	log.Println(result, reflect.TypeOf(result))
}
