package value_tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_empty_struct(t *testing.T) {
	bytes, _ := json.Marshal(&StructOfStringInt{
		Name:  "hello",
		Price: 100,
	})
	fmt.Println(string(bytes))
}
