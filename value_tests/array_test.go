package main

import (
	"encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

//go:generate go run github.com/json-iterator/tinygo/gen
type ArrayOfString = []string

func compareWithStdlib(input string, val1 interface{}, val2 interface{}) {
	iter := jsoniter.ParseBytes([]byte(input))
	jd_array_string(iter, val1.(*[]string))
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

func Test_empty_array(t *testing.T) {
	input := `[]`
	var val1 []string
	var val2 []string
	compareWithStdlib(input, &val1, &val2)
}

func Test_one_element_array(t *testing.T) {
	input := `["hello"]`
	var val1 []string
	var val2 []string
	compareWithStdlib(input, &val1, &val2)
}

func Test_incompatible_array_element(t *testing.T) {
	input := `[1, 2, 3]`
	val1 := []string{"hello"}
	val2 := []string{"hello"}
	compareWithStdlib(input, &val1, &val2)
}
