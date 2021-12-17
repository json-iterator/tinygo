package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type WithStructTag_json struct {
}
func (json WithStructTag_json) Type() interface{} {
  var val WithStructTag
  return val
}
func (json WithStructTag_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  WithStructTag_json_unmarshal(iter, out.(*WithStructTag))
}
func (json WithStructTag_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  WithStructTag_json_marshal(stream, val.(WithStructTag))
}
func WithStructTag_json_unmarshal(iter *jsoniter.Iterator, out *WithStructTag) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !WithStructTag_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func WithStructTag_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithStructTag) bool {
  switch {
  case field == `field1`:
    iter.ReadString(&(*out).Field1)
    return true
  }
  return false
}
func WithStructTag_json_marshal(stream *jsoniter.Stream, val WithStructTag) {
    stream.WriteObjectHead()
    stream.WriteObjectField(`field1`)
    stream.WriteString(val.Field1)
    stream.WriteObjectTail()
}
