package value_tests

import (
	jsoniter "github.com/json-iterator/tinygo"
)

//go:generate go run github.com/json-iterator/tinygo/gen
type EmbedNumberStruct struct {
	*jsoniter.Number
}
