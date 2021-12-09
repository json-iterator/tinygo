package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func WithNamelessField_json_unmarshal(iter *jsoniter.Iterator, out *WithNamelessField) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !WithNamelessField_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func WithNamelessField_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithNamelessField) bool {
  if WithNamelessField_f1_json_unmarshal_field(iter, field, &out.WithNamelessField_f1) { return true }
  var val1 WithNamelessField_f2
  if WithNamelessField_f2_json_unmarshal_field(iter, field, &val1) {
    out.WithNamelessField_f2 = new(WithNamelessField_f2)
    *out.WithNamelessField_f2 = val1
    return true
  }
  return false
}
type WithNamelessField_json struct {
}
func (json WithNamelessField_json) Type() interface{} {
  var val WithNamelessField
  return &val
}
func (json WithNamelessField_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  WithNamelessField_json_unmarshal(iter, val.(*WithNamelessField))
}
