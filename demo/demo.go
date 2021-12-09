package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/tinygo"
)

//go:generate go run github.com/json-iterator/tinygo/gen
type NamedArray = []string

//go:generate go run github.com/json-iterator/tinygo/gen
type RefNamedArray struct {
	Value NamedArray
}

func main() {
	// list all the types you need to unmarshal here
	json := jsoniter.CreateJsonAdapter(RefNamedArray_json{}, NamedArray_json{})

	var val1 RefNamedArray
	var val2 NamedArray
	json.Unmarshal([]byte(`{ "Value": ["hello","world"] }`), &val1)
	json.Unmarshal([]byte(`["hello","world"]`), &val2)
	fmt.Println(val1.Value)
}
