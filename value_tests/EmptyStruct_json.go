package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func EmptyStruct_json_unmarshal(iter *jsoniter.Iterator, out *EmptyStruct) {
}
type EmptyStruct_json struct {
}
func (json EmptyStruct_json) Type() interface{} {
  var val EmptyStruct
  return &val
}
func (json EmptyStruct_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  EmptyStruct_json_unmarshal(iter, val.(*EmptyStruct))
}
