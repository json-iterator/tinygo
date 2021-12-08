package value_tests

import "bytes"

//go:generate go run github.com/json-iterator/tinygo/gen
type NamedStruct struct {
	Name         string
	Price        interface{}
	privateField *bytes.Buffer
}
