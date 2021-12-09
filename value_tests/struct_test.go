package value_tests

import (
	"encoding/json"
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

func Test_struct4(t *testing.T) {
	var val1 WithEmbedStruct
	var val2 WithEmbedStruct
	compareWithStdlib(`{"Field1":"hello","Field2":"world","Field3":"abc","Embed3":"123"}`,
		jsoniter.CreateJsonAdapter(WithEmbedStruct_json{}), &val1, &val2)
}

func Test_struct5(t *testing.T) {
	var val1 WithEmbedStruct
	var val2 WithEmbedStruct
	compareWithStdlib(`{"Field2":"world","Field3":"abc","Embed3":"123"}`,
		jsoniter.CreateJsonAdapter(WithEmbedStruct_json{}), &val1, &val2)
}

func Test_struct6(t *testing.T) {
	var val1 WithNamelessField
	var val2 WithNamelessField
	compareWithStdlib(`{"WithNamelessField_f1":12.34,"WithNamelessField_f2":true,"string":"hello","int":1}`,
		jsoniter.CreateJsonAdapter(WithNamelessField_json{}), &val1, &val2)
}

func Test_struct7(t *testing.T) {
	var val1 EmbedNumberStruct
	iter := jsoniter.ParseBytes([]byte(`{"Number":100}`))
	EmbedNumberStruct_json_unmarshal(iter, &val1)
	if iter.Error != nil {
		t.Fatal(iter.Error)
	}
	if val1.Number != json.Number("100") {
		t.Fatal(val1.Number)
	}
}
