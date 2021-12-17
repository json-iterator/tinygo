package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type NamedPtr_json struct {
}
func (json NamedPtr_json) Type() interface{} {
  var val NamedPtr
  return val
}
func (json NamedPtr_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  NamedPtr_json_unmarshal(iter, out.(*NamedPtr))
}
func (json NamedPtr_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  NamedPtr_json_marshal(stream, val.(NamedPtr))
}
func NamedPtr_json_unmarshal(iter *jsoniter.Iterator, out *NamedPtr) {
    var val string
    iter.ReadString(&val)
    if iter.Error == nil {
      *out = &val
    }
}
func NamedPtr_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *NamedPtr) bool {
  if field == "NamedPtr" {
    NamedPtr_json_unmarshal(iter, out)
    return true
  }
  return false
}
func NamedPtr_json_marshal(stream *jsoniter.Stream, val NamedPtr) {
    stream.WriteString(*val)
}
