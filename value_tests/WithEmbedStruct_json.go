package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type WithEmbedStruct_json struct {
}
func (json WithEmbedStruct_json) Type() interface{} {
  var val WithEmbedStruct
  return val
}
func (json WithEmbedStruct_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  WithEmbedStruct_json_unmarshal(iter, out.(*WithEmbedStruct))
}
func (json WithEmbedStruct_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  WithEmbedStruct_json_marshal(stream, val.(WithEmbedStruct))
}
func WithEmbedStruct_json_unmarshal(iter *jsoniter.Iterator, out *WithEmbedStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !WithEmbedStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func WithEmbedStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithEmbedStruct) bool {
  var val0 WithEmbedStructBase1
  if WithEmbedStructBase1_json_unmarshal_field(iter, field, &val0) {
    out.WithEmbedStructBase1 = new(WithEmbedStructBase1)
    *out.WithEmbedStructBase1 = val0
    return true
  }
  if WithEmbedStructBase2_json_unmarshal_field(iter, field, &out.WithEmbedStructBase2) { return true }
  switch {
  case field == `Field3`:
    iter.ReadString(&(*out).Field3)
    return true
  }
  return false
}
func WithEmbedStruct_json_marshal(stream *jsoniter.Stream, val WithEmbedStruct) {
    stream.WriteObjectHead()
    WithEmbedStruct_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func WithEmbedStruct_json_marshal_field(stream *jsoniter.Stream, val WithEmbedStruct) {
    stream.WriteObjectField(`Field3`)
    stream.WriteString(val.Field3)
    stream.WriteMore()
}
