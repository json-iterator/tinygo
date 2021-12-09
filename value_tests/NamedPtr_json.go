package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func NamedPtr_json_unmarshal(iter *jsoniter.Iterator, out *NamedPtr) {
    var val string
    iter.ReadString(&val)
    if iter.Error == nil {
      *out = &val
    }
}
func NamedPtr_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *NamedPtr) bool {
  return false
}
type NamedPtr_json struct {
}
func (json NamedPtr_json) Type() interface{} {
  var val NamedPtr
  return &val
}
func (json NamedPtr_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  NamedPtr_json_unmarshal(iter, val.(*NamedPtr))
}
