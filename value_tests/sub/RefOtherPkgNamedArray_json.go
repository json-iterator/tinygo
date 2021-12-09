package sub

import jsoniter "github.com/json-iterator/tinygo"
import value_tests "github.com/json-iterator/tinygo/value_tests"

func RefOtherPkgNamedArray_json_unmarshal(iter *jsoniter.Iterator, out *RefOtherPkgNamedArray) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !RefOtherPkgNamedArray_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func RefOtherPkgNamedArray_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *RefOtherPkgNamedArray) bool {
  switch {
  case field == `Value`:
    value_tests.NamedArray_json_unmarshal(iter, &(*out).Value)
    return true
  }
  return false
}
type RefOtherPkgNamedArray_json struct {
}
func (json RefOtherPkgNamedArray_json) Type() interface{} {
  var val RefOtherPkgNamedArray
  return &val
}
func (json RefOtherPkgNamedArray_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  RefOtherPkgNamedArray_json_unmarshal(iter, val.(*RefOtherPkgNamedArray))
}
