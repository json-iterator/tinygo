package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func WithNamelessField_f2_json_unmarshal(iter *jsoniter.Iterator, out *WithNamelessField_f2) {
    iter.ReadBool((*bool)(out))
}
func WithNamelessField_f2_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithNamelessField_f2) bool {
  if field == "WithNamelessField_f2" {
    WithNamelessField_f2_json_unmarshal(iter, out)
    return true
  }
  return false
}
type WithNamelessField_f2_json struct {
}
func (json WithNamelessField_f2_json) Type() interface{} {
  var val WithNamelessField_f2
  return &val
}
func (json WithNamelessField_f2_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  WithNamelessField_f2_json_unmarshal(iter, val.(*WithNamelessField_f2))
}
