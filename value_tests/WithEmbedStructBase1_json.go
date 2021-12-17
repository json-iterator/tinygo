package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type WithEmbedStructBase1_json struct {
}
func (json WithEmbedStructBase1_json) Type() interface{} {
  var val WithEmbedStructBase1
  return val
}
func (json WithEmbedStructBase1_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  WithEmbedStructBase1_json_unmarshal(iter, out.(*WithEmbedStructBase1))
}
func (json WithEmbedStructBase1_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  WithEmbedStructBase1_json_marshal(stream, val.(WithEmbedStructBase1))
}
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
func WithEmbedStructBase1_json_marshal(stream *jsoniter.Stream, val WithEmbedStructBase1) {
    stream.WriteObjectHead()
    WithEmbedStructBase1_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func WithEmbedStructBase1_json_marshal_field(stream *jsoniter.Stream, val WithEmbedStructBase1) {
    stream.WriteObjectField(`Field1`)
    stream.WriteString(val.Field1)
    stream.WriteMore()
}
