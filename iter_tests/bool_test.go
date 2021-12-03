package iter_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_read_true(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte("true"))
	iter.AssertIsBool()
	if iter.ReadBool() == false {
		t.Fatal()
	}
}

func Test_read_false(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte("false"))
	iter.AssertIsBool()
	if iter.ReadBool() == true {
		t.Fatal()
	}
}
