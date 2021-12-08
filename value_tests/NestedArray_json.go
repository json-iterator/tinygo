package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func NestedArray_json_unmarshal(iter *jsoniter.Iterator, out *NestedArray) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i == len(val) {
      val = append(val, make(NestedArray, 4)...)
    }
    NestedArray_array1_json_unmarshal(iter, &val[i])
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = NestedArray{}
  } else {
    *out = val[:i]
  }
}
func NestedArray_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *NestedArray) bool {
  return false
}
func NestedArray_array1_json_unmarshal (iter *jsoniter.Iterator, out *[2]float64) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i < 2 {
    iter.ReadFloat64(&val[i])
    } else {
      iter.Skip()
    }
    i++
    more = iter.ReadArrayMore()
  }
}
type NestedArray_json struct {
}
func (json NestedArray_json) Type() interface{} {
  var val NestedArray
  return &val
}
func (json NestedArray_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  NestedArray_json_unmarshal(iter, val.(*NestedArray))
}
