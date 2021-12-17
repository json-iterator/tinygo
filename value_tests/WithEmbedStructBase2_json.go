package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type WithEmbedStructBase2_json struct {
}
func (json WithEmbedStructBase2_json) Type() interface{} {
  var val WithEmbedStructBase2
  return val
}
func (json WithEmbedStructBase2_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  WithEmbedStructBase2_json_unmarshal(iter, out.(*WithEmbedStructBase2))
}
func (json WithEmbedStructBase2_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  WithEmbedStructBase2_json_marshal(stream, val.(WithEmbedStructBase2))
}
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
func WithEmbedStructBase2_json_marshal(stream *jsoniter.Stream, val WithEmbedStructBase2) {
    stream.WriteObjectHead()
    stream.WriteObjectField(`Field2`)
    stream.WriteString(val.Field2)
    stream.WriteObjectTail()
}
