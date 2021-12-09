package value_tests

import jsoniter "github.com/json-iterator/tinygo"
import json "encoding/json"

func NamedStruct_json_unmarshal(iter *jsoniter.Iterator, out *NamedStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !NamedStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func NamedStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *NamedStruct) bool {
  switch {
  case field == `Name`:
    iter.ReadString(&(*out).Name)
    return true
  case field == `Price`:
    NamedStruct_ptr1_json_unmarshal(iter, &(*out).Price)
    return true
  }
  return false
}
func NamedStruct_ptr1_json_unmarshal (iter *jsoniter.Iterator, out **json.Number) {
    var val json.Number
    iter.ReadNumber((*jsoniter.Number)(&val))
    if iter.Error == nil {
      *out = &val
    }
}
type NamedStruct_json struct {
}
func (json NamedStruct_json) Type() interface{} {
  var val NamedStruct
  return &val
}
func (json NamedStruct_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  NamedStruct_json_unmarshal(iter, val.(*NamedStruct))
}
