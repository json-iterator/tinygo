package value_tests

import (
	"encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func compareWithStdlib(input string, val1 interface{}, val2 interface{}) {
	iter := jsoniter.ParseBytes([]byte(input))
	NamedArray_json_unmarshal(iter, val1.(*[]string))
	bytes1, err := json.Marshal(val1)
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal([]byte(input), val2)
	bytes2, err := json.Marshal(val2)
	if err != nil {
		panic(err)
	}
	if iter.Error != nil && err2 == nil {
		panic(fmt.Errorf("expect no error, but found error: %w", iter.Error))
	}
	if iter.Error == nil && err2 != nil {
		panic("expect error, but found no error")
	}
	if string(bytes1) != string(bytes2) {
		panic(fmt.Sprintf("expect ...%s..., actual ...%s...", string(bytes1), string(bytes2)))
	}
}

func Test_array1(t *testing.T) {
	input := `[]`
	var val1 NamedArray
	var val2 NamedArray
	compareWithStdlib(input, &val1, &val2)
}

func Test_array2(t *testing.T) {
	input := `["hello"]`
	var val1 NamedArray
	var val2 NamedArray
	compareWithStdlib(input, &val1, &val2)
}

func Test_array3(t *testing.T) {
	input := `[10, 20, 30]`
	val1 := NamedArray{"hello"}
	val2 := NamedArray{"hello"}
	compareWithStdlib(input, &val1, &val2)
}

func Test_array4(t *testing.T) {
	input := `[100, "world"]`
	val1 := NamedArray{"hello"}
	val2 := NamedArray{"hello"}
	compareWithStdlib(input, &val1, &val2)
}

func Test_array5(t *testing.T) {
	input := `[null, "world"]`
	val1 := NamedArray{"hello"}
	val2 := NamedArray{"hello"}
	compareWithStdlib(input, &val1, &val2)
}
