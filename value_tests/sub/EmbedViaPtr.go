package sub

import "github.com/json-iterator/tinygo/value_tests"

//go:generate go run github.com/json-iterator/tinygo/gen
type EmbedViaPtr struct {
	*value_tests.NamedArray
}
