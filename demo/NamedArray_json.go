package main

import jsoniter "github.com/json-iterator/tinygo"

func NamedArray_json_unmarshal(iter *jsoniter.Iterator, out *NamedArray) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i == len(val) {
      val = append(val, make(NamedArray, 4)...)
    }
    iter.ReadString(&val[i])
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = NamedArray{}
  } else {
    *out = val[:i]
  }
}
func NamedArray_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *NamedArray) bool {
  if field == "NamedArray" {
    NamedArray_json_unmarshal(iter, out)
    return true
  }
  return false
}
type NamedArray_json struct {
}
func (json NamedArray_json) Type() interface{} {
  var val NamedArray
  return &val
}
func (json NamedArray_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  NamedArray_json_unmarshal(iter, val.(*NamedArray))
}
