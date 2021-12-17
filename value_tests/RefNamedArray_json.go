package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type RefNamedArray_json struct {
}
func (json RefNamedArray_json) Type() interface{} {
  var val RefNamedArray
  return val
}
func (json RefNamedArray_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  RefNamedArray_json_unmarshal(iter, out.(*RefNamedArray))
}
func (json RefNamedArray_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  RefNamedArray_json_marshal(stream, val.(RefNamedArray))
}
func RefNamedArray_json_unmarshal(iter *jsoniter.Iterator, out *RefNamedArray) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !RefNamedArray_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func RefNamedArray_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *RefNamedArray) bool {
  switch {
  case field == `Value`:
    NamedArray_json_unmarshal(iter, &(*out).Value)
    return true
  }
  return false
}
func RefNamedArray_json_marshal(stream *jsoniter.Stream, val RefNamedArray) {
    stream.WriteObjectHead()
    RefNamedArray_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func RefNamedArray_json_marshal_field(stream *jsoniter.Stream, val RefNamedArray) {
    stream.WriteObjectField(`Value`)
    NamedArray_json_marshal(stream, val.Value)
    stream.WriteMore()
}
