package sub

import jsoniter "github.com/json-iterator/tinygo"
import value_tests "github.com/json-iterator/tinygo/value_tests"

type EmbedViaPtr_json struct {
}
func (json EmbedViaPtr_json) Type() interface{} {
  var val EmbedViaPtr
  return val
}
func (json EmbedViaPtr_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  EmbedViaPtr_json_unmarshal(iter, out.(*EmbedViaPtr))
}
func (json EmbedViaPtr_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  EmbedViaPtr_json_marshal(stream, val.(EmbedViaPtr))
}
func EmbedViaPtr_json_unmarshal(iter *jsoniter.Iterator, out *EmbedViaPtr) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !EmbedViaPtr_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func EmbedViaPtr_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *EmbedViaPtr) bool {
  var val0 value_tests.NamedArray
  if value_tests.NamedArray_json_unmarshal_field(iter, field, &val0) {
    out.NamedArray = new(value_tests.NamedArray)
    *out.NamedArray = val0
    return true
  }
  return false
}
func EmbedViaPtr_json_marshal(stream *jsoniter.Stream, val EmbedViaPtr) {
    stream.WriteObjectHead()
    EmbedViaPtr_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func EmbedViaPtr_json_marshal_field(stream *jsoniter.Stream, val EmbedViaPtr) {
}
