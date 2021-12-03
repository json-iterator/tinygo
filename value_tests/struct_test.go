package value_tests

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_empty_struct(t *testing.T) {
	input := `{"Name":"hello","Price":100}`
	iter := jsoniter.ParseBytes([]byte(input))
	var val1 StructOfStringInt
	jd_StructOfStringInt(iter, &val1)
	if iter.Error != nil {
		t.Fatal(iter.Error)
	}
	bytes1, err := json.Marshal(val1)
	if err != nil {
		t.Fatal(err)
	}
	var val2 StructOfStringInt
	err = json.Unmarshal([]byte(input), &val2)
	if err != nil {
		t.Fatal(err)
	}
	bytes2, err := json.Marshal(val2)
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes1) != string(bytes2) {
		t.Fatal()
	}
}
