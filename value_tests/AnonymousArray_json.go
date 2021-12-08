package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func AnonymousArray_json_unmarshal(iter *jsoniter.Iterator, out *AnonymousArray) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !AnonymousArray_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func AnonymousArray_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *AnonymousArray) bool {
  switch {
  case field == `Value`:
    AnonymousArray_array1_json_unmarshal(iter, &(*out).Value)
    return true
  }
  return false
}
func AnonymousArray_array1_json_unmarshal (iter *jsoniter.Iterator, out *[]string) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i == len(val) {
      val = append(val, make([]string, 4)...)
    }
    iter.ReadString(&val[i])
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = []string{}
  } else {
    *out = val[:i]
  }
}
type AnonymousArray_json struct {
}
func (json AnonymousArray_json) Type() interface{} {
  var val AnonymousArray
  return &val
}
func (json AnonymousArray_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  AnonymousArray_json_unmarshal(iter, val.(*AnonymousArray))
}
