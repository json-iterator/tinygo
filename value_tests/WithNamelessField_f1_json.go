package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func WithNamelessField_f1_json_unmarshal(iter *jsoniter.Iterator, out *WithNamelessField_f1) {
    iter.ReadFloat64((*float64)(out))
}
func WithNamelessField_f1_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithNamelessField_f1) bool {
  if field == "WithNamelessField_f1" {
    WithNamelessField_f1_json_unmarshal(iter, out)
    return true
  }
  return false
}
type WithNamelessField_f1_json struct {
}
func (json WithNamelessField_f1_json) Type() interface{} {
  var val WithNamelessField_f1
  return &val
}
func (json WithNamelessField_f1_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  WithNamelessField_f1_json_unmarshal(iter, val.(*WithNamelessField_f1))
}
