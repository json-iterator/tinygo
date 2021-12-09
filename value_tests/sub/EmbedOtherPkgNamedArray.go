package sub

import "github.com/json-iterator/tinygo/value_tests"

//go:generate go run github.com/json-iterator/tinygo/gen
type EmbedOtherPkgNamedArray struct {
	value_tests.NamedArray
}
