package value_tests

import (
	"bytes"
	"encoding/json"
)

//go:generate go run github.com/json-iterator/tinygo/gen
type NamedStruct struct {
	Name         string
	Price        *json.Number
	Raw          json.RawMessage
	privateField *bytes.Buffer
}
