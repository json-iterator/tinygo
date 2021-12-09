package value_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

type SomeStruct struct {
	Field1 **string
}

func Test_ptr1(t *testing.T) {
	var val1 NamedPtr
	var val2 NamedPtr
	compareWithStdlib(`"hello"`, jsoniter.CreateJsonAdapter(NamedPtr_json{}), &val1, &val2)
}
