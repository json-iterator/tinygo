package value_tests

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_empty_struct(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`{"Name":"hello","Price":100}`))
	var val StructOfStringInt
	jd_StructOfStringInt(iter, &val)
	if iter.Error != nil {
		t.Fatal(iter.Error)
	}
	fmt.Println(val)
}
