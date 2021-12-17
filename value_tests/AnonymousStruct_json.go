package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type AnonymousStruct_json struct {
}
func (json AnonymousStruct_json) Type() interface{} {
  var val AnonymousStruct
  return val
}
func (json AnonymousStruct_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  AnonymousStruct_json_unmarshal(iter, out.(*AnonymousStruct))
}
func (json AnonymousStruct_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  AnonymousStruct_json_marshal(stream, val.(AnonymousStruct))
}
func AnonymousStruct_json_unmarshal(iter *jsoniter.Iterator, out *AnonymousStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !AnonymousStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func AnonymousStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *AnonymousStruct) bool {
  switch {
  case field == `Value`:
    AnonymousStruct_struct1_json_unmarshal(iter, &(*out).Value)
    return true
  }
  return false
}
func AnonymousStruct_struct1_json_unmarshal_field (iter *jsoniter.Iterator, field string, out *struct {
	Name	string
	Price	int
}) bool {
  switch {
  case field == `Name`:
    iter.ReadString(&(*out).Name)
    return true
  case field == `Price`:
    iter.ReadInt(&(*out).Price)
    return true
  }
  return false
}
func AnonymousStruct_struct1_json_unmarshal (iter *jsoniter.Iterator, out *struct {
	Name	string
	Price	int
}) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !AnonymousStruct_struct1_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func AnonymousStruct_json_marshal(stream *jsoniter.Stream, val AnonymousStruct) {
    stream.WriteObjectHead()
    AnonymousStruct_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func AnonymousStruct_json_marshal_field(stream *jsoniter.Stream, val AnonymousStruct) {
    stream.WriteObjectField(`Value`)
    AnonymousStruct_struct2_json_marshal(stream, val.Value)
    stream.WriteMore()
}
func AnonymousStruct_struct2_json_marshal_field (stream *jsoniter.Stream, val struct {
	Name	string
	Price	int
}) {
    stream.WriteObjectField(`Name`)
    stream.WriteString(val.Name)
    stream.WriteMore()
    stream.WriteObjectField(`Price`)
    stream.WriteInt(val.Price)
    stream.WriteMore()
}
func AnonymousStruct_struct2_json_marshal (stream *jsoniter.Stream, val struct {
	Name	string
	Price	int
}) {
    stream.WriteObjectHead()
    AnonymousStruct_struct2_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
