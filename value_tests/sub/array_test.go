package sub

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

func Test_RefOtherPkgNamedArray(t *testing.T) {
	input := `{ "Value": ["hello","world"] }`
	var val1 RefOtherPkgNamedArray
	var val2 RefOtherPkgNamedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(RefOtherPkgNamedArray_json{}), &val1, &val2)
}

func Test_EmbedOtherPkgNamedArray(t *testing.T) {
	input := `{ "NamedArray": ["hello","world"] }`
	var val1 EmbedOtherPkgNamedArray
	var val2 EmbedOtherPkgNamedArray
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(EmbedOtherPkgNamedArray_json{}), &val1, &val2)
}

func Test_EmbedViaPtr(t *testing.T) {
	input := `{ "NamedArray": ["hello","world"] }`
	var val1 EmbedViaPtr
	var val2 EmbedViaPtr
	compareWithStdlib(input, jsoniter.CreateJsonAdapter(EmbedViaPtr_json{}), &val1, &val2)
}
