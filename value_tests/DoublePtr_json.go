package value_tests

import jsoniter "github.com/json-iterator/tinygo"

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
type DoublePtr_json struct {
}
func (json DoublePtr_json) Type() interface{} {
  var val DoublePtr
  return &val
}
func (json DoublePtr_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  DoublePtr_json_unmarshal(iter, val.(*DoublePtr))
}
