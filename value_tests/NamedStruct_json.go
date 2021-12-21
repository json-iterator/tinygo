package value_tests

import jsoniter "github.com/json-iterator/tinygo"
import json "encoding/json"

type NamedStruct_json struct {
}
func (json NamedStruct_json) Type() interface{} {
  var val NamedStruct
  return val
}
func (json NamedStruct_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  NamedStruct_json_unmarshal(iter, out.(*NamedStruct))
}
func (json NamedStruct_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  NamedStruct_json_marshal(stream, val.(NamedStruct))
}
func NamedStruct_json_unmarshal(iter *jsoniter.Iterator, out *NamedStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !NamedStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func NamedStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *NamedStruct) bool {
  switch {
  case field == `Name`:
    iter.ReadString(&(*out).Name)
    return true
  case field == `Price`:
    NamedStruct_ptr1_json_unmarshal(iter, &(*out).Price)
    return true
  case field == `Raw`:
    iter.ReadRawMessage((*jsoniter.RawMessage)(&(*out).Raw))
    return true
  }
  return false
}
func NamedStruct_ptr1_json_unmarshal (iter *jsoniter.Iterator, out **json.Number) {
    var val json.Number
    iter.ReadNumber((*jsoniter.Number)(&val))
    if iter.Error == nil {
      *out = &val
    }
}
func NamedStruct_json_marshal(stream *jsoniter.Stream, val NamedStruct) {
    stream.WriteObjectHead()
    NamedStruct_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func NamedStruct_json_marshal_field(stream *jsoniter.Stream, val NamedStruct) {
    stream.WriteObjectField(`Name`)
    stream.WriteString(val.Name)
    stream.WriteMore()
    stream.WriteObjectField(`Price`)
    if val.Price == nil {
       stream.WriteNull()
    } else {
    stream.WriteRawOrZero(string(*val.Price))
    }
    stream.WriteMore()
    stream.WriteObjectField(`Raw`)
    stream.WriteRawOrNull(string(val.Raw))
    stream.WriteMore()
}
