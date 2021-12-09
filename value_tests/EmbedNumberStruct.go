package value_tests

import (
	"encoding/json"
)

//go:generate go run github.com/json-iterator/tinygo/gen
type EmbedNumberStruct struct {
	json.Number
}
