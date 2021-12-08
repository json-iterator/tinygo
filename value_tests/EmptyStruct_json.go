package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func EmptyStruct_json_unmarshal(iter *jsoniter.Iterator, out *EmptyStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !EmptyStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func EmptyStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *EmptyStruct) bool {
  return false
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
