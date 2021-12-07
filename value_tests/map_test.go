package value_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_map1(t *testing.T) {
	input := `{}`
	var val1 map[string]string
	var val2 map[string]string
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedMap_json{}), &val1, &val2)
}
