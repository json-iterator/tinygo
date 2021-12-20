package sub

import jsoniter "github.com/json-iterator/tinygo"
import value_tests "github.com/json-iterator/tinygo/value_tests"

type RefOtherPkgNamedArray_json struct {
}
func (json RefOtherPkgNamedArray_json) Type() interface{} {
  var val RefOtherPkgNamedArray
  return val
}
func (json RefOtherPkgNamedArray_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  RefOtherPkgNamedArray_json_unmarshal(iter, out.(*RefOtherPkgNamedArray))
}
func (json RefOtherPkgNamedArray_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  RefOtherPkgNamedArray_json_marshal(stream, val.(RefOtherPkgNamedArray))
}
func RefOtherPkgNamedArray_json_unmarshal(iter *jsoniter.Iterator, out *RefOtherPkgNamedArray) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !RefOtherPkgNamedArray_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func RefOtherPkgNamedArray_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *RefOtherPkgNamedArray) bool {
  switch {
  case field == `Value`:
    value_tests.NamedArray_json_unmarshal(iter, &(*out).Value)
    return true
  }
  return false
}
func RefOtherPkgNamedArray_json_marshal(stream *jsoniter.Stream, val RefOtherPkgNamedArray) {
    stream.WriteObjectHead()
    RefOtherPkgNamedArray_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func RefOtherPkgNamedArray_json_marshal_field(stream *jsoniter.Stream, val RefOtherPkgNamedArray) {
    stream.WriteObjectField(`Value`)
    stream.WriteMore()
}
