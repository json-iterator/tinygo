package tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_empty_array(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`[]`))
	if iter.ReadArray() {
		t.Fail()
	}
}

func Test_one_element_array(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`[1]`))
	if !iter.ReadArray() {
		t.Fail()
	}
	if iter.ReadInt() != 1 {
		t.Fail()
	}
	if iter.ReadArray() {
		t.Fail()
	}
}
