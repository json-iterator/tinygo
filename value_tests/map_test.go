package value_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_map1(t *testing.T) {
	input := `{}`
	var val1 map[string]int
	var val2 map[string]int
	stream := jsoniter.NewStream()
	stream.Prefix = ""
	stream.Indent = "    "
	NamedMap_json_marshal(stream, val1)
	if string(stream.Buffer()) != "{}" {
		t.Fatal()
	}
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

func Test_map5(t *testing.T) {
	input := `{"100":"world","200":"hello"}`
	var val1 map[int]string
	var val2 map[int]string
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(IntKeyMap_json{}), &val1, &val2)
}

func Test_map6(t *testing.T) {
	input := `{"100":"world","abc":"hello"}`
	var val1 map[int]string
	var val2 map[int]string
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(IntKeyMap_json{}), &val1, &val2)
}
