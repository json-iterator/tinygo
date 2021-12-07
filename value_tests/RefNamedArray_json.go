package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func RefNamedArray_json_unmarshal(iter *jsoniter.Iterator, out *RefNamedArray) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    switch {
    case field == `Value`:
    NamedArray_json_unmarshal(iter, &(*out).Value)
    default:
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
type RefNamedArray_json struct {
}
func (json RefNamedArray_json) Type() interface{} {
  var val RefNamedArray
  return &val
}
func (json RefNamedArray_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  RefNamedArray_json_unmarshal(iter, val.(*RefNamedArray))
}
