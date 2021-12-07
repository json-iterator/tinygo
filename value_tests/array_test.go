package value_tests

import (
	"encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func compareWithStdlib(input string, json1 jsoniter.JsonAdapter, val1 interface{}, val2 interface{}) {
	err1 := json1.Unmarshal([]byte(input), val1)
	bytes1, err := json.Marshal(val1)
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal([]byte(input), val2)
	bytes2, err := json.Marshal(val2)
	if err != nil {
		panic(err)
	}
	if err1 != nil && err2 == nil {
		panic(fmt.Errorf("expect no error, but found error: %w", err1))
	}
	if err1 == nil && err2 != nil {
		panic("expect error, but found no error")
	}
	if string(bytes1) != string(bytes2) {
		panic(fmt.Sprintf("expect ...%s..., actual ...%s...", string(bytes2), string(bytes1)))
	}
}

func Test_array1(t *testing.T) {
	input := `[]`
	var val1 NamedArray
	var val2 NamedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedArray_json{}), &val1, &val2)
}

func Test_array2(t *testing.T) {
	input := `["hello"]`
	var val1 NamedArray
	var val2 NamedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedArray_json{}), &val1, &val2)
}

func Test_array3(t *testing.T) {
	input := `[10, 20, 30]`
	val1 := NamedArray{"hello"}
	val2 := NamedArray{"hello"}
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedArray_json{}), &val1, &val2)
}

func Test_array4(t *testing.T) {
	input := `[100, "world"]`
	val1 := NamedArray{"hello"}
	val2 := NamedArray{"hello"}
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedArray_json{}), &val1, &val2)
}

func Test_array5(t *testing.T) {
	input := `[null, "world"]`
	val1 := NamedArray{"hello"}
	val2 := NamedArray{"hello"}
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NamedArray_json{}), &val1, &val2)
}

func Test_array6(t *testing.T) {
	input := `{ "Value": ["hello","world"] }`
	var val1 AnonymousArray
	var val2 AnonymousArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(AnonymousArray_json{}), &val1, &val2)
}

func Test_array7(t *testing.T) {
	input := `{ "Value": ["hello","world"] }`
	var val1 RefNamedArray
	var val2 RefNamedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(RefNamedArray_json{}), &val1, &val2)
}

func Test_array8(t *testing.T) {
	input := `[["hello"]]`
	var val1 NestedArray
	var val2 NestedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NestedArray_json{}), &val1, &val2)
}

func Test_array9(t *testing.T) {
	input := `[["a","b","c"]]`
	var val1 NestedArray
	var val2 NestedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NestedArray_json{}), &val1, &val2)
}

func Test_array10(t *testing.T) {
	input := `[[]]`
	var val1 NestedArray
	var val2 NestedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(NestedArray_json{}), &val1, &val2)
}
