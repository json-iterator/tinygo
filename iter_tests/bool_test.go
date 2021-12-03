package iter_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_read_true(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte("true"))
	var val bool
	iter.ReadBool(&val)
	if val == false {
		t.Fatal()
	}
}

func Test_read_false(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte("false"))
	var val bool
	iter.ReadBool(&val)
	if val == true {
		t.Fatal()
	}
}
