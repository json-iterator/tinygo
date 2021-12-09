# make `json.Unmarshal` work in tinygo

```go
type NamedArray = []string

type RefNamedArray struct {
	Value NamedArray
}

import "encoding/json"

var val1 RefNamedArray
json.Unmarshal([]byte(`{ "Value": ["hello","world"] }`), &val1)
json.Unmarshal([]byte(`["hello","world"]`), &val2) 
```

The code above does not work in tinygo, due to incomplete runtime reflection support. To fix this, we use code generation instead of runtime reflection to implement `json.Unmarshal`

```go
//go:generate go run github.com/json-iterator/tinygo/gen
type NamedArray = []string

//go:generate go run github.com/json-iterator/tinygo/gen
type RefNamedArray struct {
	Value NamedArray
}

import "github.com/json-iterator/tinygo"
// list all the types you need to unmarshal here
json := jsoniter.CreateJsonAdapter(RefNamedArray_json{}, NamedArray_json{}) 

var val1 RefNamedArray
json.Unmarshal([]byte(`{ "Value": ["hello","world"] }`), &val1)
json.Unmarshal([]byte(`["hello","world"]`), &val2) 
```

run `go generate` command to generate RefNamedArray_json and NamedArray_json.