package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type EmbedNumberStruct_json struct {
}
func (json EmbedNumberStruct_json) Type() interface{} {
  var val EmbedNumberStruct
  return val
}
func (json EmbedNumberStruct_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  EmbedNumberStruct_json_unmarshal(iter, out.(*EmbedNumberStruct))
}
func (json EmbedNumberStruct_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  EmbedNumberStruct_json_marshal(stream, val.(EmbedNumberStruct))
}
func EmbedNumberStruct_json_unmarshal(iter *jsoniter.Iterator, out *EmbedNumberStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !EmbedNumberStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func EmbedNumberStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *EmbedNumberStruct) bool {
  if field == "Number" { iter.ReadNumber((*jsoniter.Number)(&out.Number)); return true }
  return false
}
func EmbedNumberStruct_json_marshal(stream *jsoniter.Stream, val EmbedNumberStruct) {
    stream.WriteObjectHead()
    stream.WriteObjectTail()
}
