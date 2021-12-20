package sub

import jsoniter "github.com/json-iterator/tinygo"
import value_tests "github.com/json-iterator/tinygo/value_tests"

type EmbedOtherPkgNamedArray_json struct {
}
func (json EmbedOtherPkgNamedArray_json) Type() interface{} {
  var val EmbedOtherPkgNamedArray
  return val
}
func (json EmbedOtherPkgNamedArray_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  EmbedOtherPkgNamedArray_json_unmarshal(iter, out.(*EmbedOtherPkgNamedArray))
}
func (json EmbedOtherPkgNamedArray_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  EmbedOtherPkgNamedArray_json_marshal(stream, val.(EmbedOtherPkgNamedArray))
}
func EmbedOtherPkgNamedArray_json_unmarshal(iter *jsoniter.Iterator, out *EmbedOtherPkgNamedArray) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !EmbedOtherPkgNamedArray_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func EmbedOtherPkgNamedArray_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *EmbedOtherPkgNamedArray) bool {
  if value_tests.NamedArray_json_unmarshal_field(iter, field, &out.NamedArray) { return true }
  return false
}
func EmbedOtherPkgNamedArray_json_marshal(stream *jsoniter.Stream, val EmbedOtherPkgNamedArray) {
    stream.WriteObjectHead()
    EmbedOtherPkgNamedArray_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func EmbedOtherPkgNamedArray_json_marshal_field(stream *jsoniter.Stream, val EmbedOtherPkgNamedArray) {
}
