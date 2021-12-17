package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type NamedArray_json struct {
}
func (json NamedArray_json) Type() interface{} {
  var val NamedArray
  return val
}
func (json NamedArray_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  NamedArray_json_unmarshal(iter, out.(*NamedArray))
}
func (json NamedArray_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  NamedArray_json_marshal(stream, val.(NamedArray))
}
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
func NamedArray_json_marshal(stream *jsoniter.Stream, val NamedArray) {
  if len(val) == 0 {
    stream.WriteEmptyArray()
  } else {
    stream.WriteArrayHead()
    for i, elem := range val {
      if i != 0 { stream.WriteMore() }
    stream.WriteString(elem)
    }
    stream.WriteArrayTail()
  }
}
