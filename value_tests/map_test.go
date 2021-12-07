package value_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_map1(t *testing.T) {
	input := `{}`
	var val1 map[string]int
	var val2 map[string]int
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedMap_json{}), &val1, &val2)
}

func Test_map2(t *testing.T) {
	input := `{"hello":100}`
	var val1 map[string]int
	var val2 map[string]int
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedMap_json{}), &val1, &val2)
}

func Test_map3(t *testing.T) {
	input := `{"hello":"world"}`
	var val1 map[string]int
	var val2 map[string]int
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedMap_json{}), &val1, &val2)
}

func Test_map4(t *testing.T) {
	input := `{"Value":{"hello":"world"}}`
	var val1 AnonymousMap
	var val2 AnonymousMap
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(AnonymousMap_json{}), &val1, &val2)
}
