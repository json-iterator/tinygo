# make `json.Unmarshal` work in tinygo

make `json.Unmarshal` work in tinygo with minimal code size.

* for the value that parsed out, it will always validate the input
* for the value that mismatching type or not needed, it will NOT validate the input

# usage

```go
//go:generate go run github.com/json-iterator/tinygo/gen
type NamedArray = []string

//go:generate go run github.com/json-iterator/tinygo/gen
type RefNamedArray struct {
	Value NamedArray
}
```

`go generate` will produce the following function in `NamedArray_json.go` and `RefNamedArray_json.go`

```go
package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func NamedArray_json_unmarshal(iter *jsoniter.Iterator, out *NamedArray) {
...
}
type NamedArray_json struct {
}
func (json NamedArray_json) Type() interface{} {
  var val NamedArray
  return &val
}
func (json NamedArray_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  NamedArray_json_unmarshal(iter, val.(*NamedArray))
}
```

```go
package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func RefNamedArray_json_unmarshal(iter *jsoniter.Iterator, out *RefNamedArray) {
...
}
type RefNamedArray_json struct {
}
func (json RefNamedArray_json) Type() interface{} {
  var val RefNamedArray
  return &val
}
func (json RefNamedArray_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  RefNamedArray_json_unmarshal(iter, val.(*RefNamedArray))
}
```

We can use `NamedArray_json_unmarshal` or `RefNamedArray_json_unmarshal` directly. Or we can use `json.Unmarshal`

```go
var val1 RefNamedArray
json := jsoniter.CreateJsonAdapter(RefNamedArray_json{}, NamedArray_json{}) // list all the types you need to unmarshal here
err := json.Unmarshal([]byte(`{ "Value": ["hello","world"] }`), &val1)
if err != nil {
    ...
} else {
    fmt.Println(val1)
}
var val2 NamedArray
json.Unmarshal([]byte(`["hello","world"]`), &val2) // json.Unmarshal will choose the decoder based on type
```
