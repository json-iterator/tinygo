package iter_tests

import (
	"encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func readEmptyInterface(input string, expectedValue interface{}) {
	iter := jsoniter.ParseBytes([]byte(input))
	var val interface{}
	iter.ReadInterface(&val)
	if val != expectedValue {
		panic(fmt.Errorf("expect %s, acutal %s", expectedValue, val))
	}
}

func Test_read_empty_interface(t *testing.T) {
	readEmptyInterface("true", true)
	readEmptyInterface("false", false)
	readEmptyInterface("100", float64(100))
	readEmptyInterface(`"hello"`, "hello")
}

func Test_read_array_as_empty_interface(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`[1,true,"hello"]`))
	var val interface{}
	iter.ReadInterface(&val)
	bytes, err := json.Marshal(val)
	if err != nil {
		t.Fatal()
	}
	if string(bytes) != `[1,true,"hello"]` {
		t.Fatal()
	}
}

func Test_read_object_as_empty_interface(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`{"hello":"world"}`))
	var val interface{}
	iter.ReadInterface(&val)
	bytes, err := json.Marshal(val)
	if err != nil {
		t.Fatal()
	}
	if string(bytes) != `{"hello":"world"}` {
		t.Fatal()
	}
}
