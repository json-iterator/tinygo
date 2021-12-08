package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func WithEmbedStructBase1_json_unmarshal(iter *jsoniter.Iterator, out *WithEmbedStructBase1) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !WithEmbedStructBase1_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func WithEmbedStructBase1_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithEmbedStructBase1) bool {
  switch {
  case field == `Field1`:
    iter.ReadString(&(*out).Field1)
    return true
  }
  return false
}
type WithEmbedStructBase1_json struct {
}
func (json WithEmbedStructBase1_json) Type() interface{} {
  var val WithEmbedStructBase1
  return &val
}
func (json WithEmbedStructBase1_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  WithEmbedStructBase1_json_unmarshal(iter, val.(*WithEmbedStructBase1))
}
