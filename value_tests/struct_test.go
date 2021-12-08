package value_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_struct1(t *testing.T) {
	var val1 NamedStruct
	var val2 NamedStruct
	compareWithStdlib(`{"Name":"hello","Price":100}`, jsoniter.CreateJsonAdapter(NamedStruct_json{}), &val1, &val2)
}

func Test_struct2(t *testing.T) {
	var val1 AnonymousStruct
	var val2 AnonymousStruct
	compareWithStdlib(`{ "Value": {"Name":"hello","Price":100} }`, jsoniter.CreateJsonAdapter(AnonymousStruct_json{}), &val1, &val2)
}

func Test_struct3(t *testing.T) {
	var val1 WithStructTag
	var val2 WithStructTag
	compareWithStdlib(`{ "field1": "hello", "Field2": "world" }`, jsoniter.CreateJsonAdapter(WithStructTag_json{}), &val1, &val2)
}
