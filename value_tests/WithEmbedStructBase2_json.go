package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func WithEmbedStructBase2_json_unmarshal(iter *jsoniter.Iterator, out *WithEmbedStructBase2) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !WithEmbedStructBase2_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func WithEmbedStructBase2_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithEmbedStructBase2) bool {
  switch {
  case field == `Field2`:
    iter.ReadString(&(*out).Field2)
    return true
  }
  return false
}
type WithEmbedStructBase2_json struct {
}
func (json WithEmbedStructBase2_json) Type() interface{} {
  var val WithEmbedStructBase2
  return &val
}
func (json WithEmbedStructBase2_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  WithEmbedStructBase2_json_unmarshal(iter, val.(*WithEmbedStructBase2))
}
