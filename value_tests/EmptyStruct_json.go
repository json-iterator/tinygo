package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type EmptyStruct_json struct {
}
func (json EmptyStruct_json) Type() interface{} {
  var val EmptyStruct
  return val
}
func (json EmptyStruct_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  EmptyStruct_json_unmarshal(iter, out.(*EmptyStruct))
}
func (json EmptyStruct_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  EmptyStruct_json_marshal(stream, val.(EmptyStruct))
}
func EmptyStruct_json_unmarshal(iter *jsoniter.Iterator, out *EmptyStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !EmptyStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func EmptyStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *EmptyStruct) bool {
  return false
}
func EmptyStruct_json_marshal(stream *jsoniter.Stream, val EmptyStruct) {
    stream.WriteObjectHead()
    EmptyStruct_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func EmptyStruct_json_marshal_field(stream *jsoniter.Stream, val EmptyStruct) {
}
