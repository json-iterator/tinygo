package main

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

//go:generate go run github.com/json-iterator/tinygo/gen
type ArrayOfString = []string

func Test_empty_array(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`[]`))
	val := jd_array_string(iter)
	if len(val) != 0 {
		t.Fail()
	}
}

func Test_one_element_array(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`["hello"]`))
	val := jd_array_string(iter)
	if len(val) != 1 {
		t.Fail()
	}
	if val[0] != "hello" {
		t.Fail()
	}
}
