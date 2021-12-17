package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type DoublePtr_json struct {
}
func (json DoublePtr_json) Type() interface{} {
  var val DoublePtr
  return val
}
func (json DoublePtr_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  DoublePtr_json_unmarshal(iter, out.(*DoublePtr))
}
func (json DoublePtr_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  DoublePtr_json_marshal(stream, val.(DoublePtr))
}
func DoublePtr_json_unmarshal(iter *jsoniter.Iterator, out *DoublePtr) {
    var val *string
    DoublePtr_ptr1_json_unmarshal(iter, &val)
    if iter.Error == nil {
      *out = &val
    }
}
func DoublePtr_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *DoublePtr) bool {
  if field == "DoublePtr" {
    DoublePtr_json_unmarshal(iter, out)
    return true
  }
  return false
}
func DoublePtr_ptr1_json_unmarshal (iter *jsoniter.Iterator, out **string) {
    var val string
    iter.ReadString(&val)
    if iter.Error == nil {
      *out = &val
    }
}
func DoublePtr_json_marshal(stream *jsoniter.Stream, val DoublePtr) {
    if *val == nil {
       stream.WriteNull()
    } else {
    stream.WriteString(**val)
    }
}
func DoublePtr_json_marshal_field(stream *jsoniter.Stream, val DoublePtr) {
    stream.WriteObjectField("DoublePtr")
    DoublePtr_json_marshal(stream, val)
    stream.WriteMore()
}
