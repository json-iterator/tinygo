package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type NestedArray_json struct {
}
func (json NestedArray_json) Type() interface{} {
  var val NestedArray
  return val
}
func (json NestedArray_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  NestedArray_json_unmarshal(iter, out.(*NestedArray))
}
func (json NestedArray_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  NestedArray_json_marshal(stream, val.(NestedArray))
}
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
  if field == "NestedArray" {
    NestedArray_json_unmarshal(iter, out)
    return true
  }
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
func NestedArray_json_marshal(stream *jsoniter.Stream, val NestedArray) {
  if len(val) == 0 {
    stream.WriteEmptyArray()
  } else {
    stream.WriteArrayHead()
    for i, elem := range val {
      if i != 0 { stream.WriteMore() }
    NestedArray_array2_json_marshal(stream, elem)
    }
    stream.WriteArrayTail()
  }
}
func NestedArray_array2_json_marshal (stream *jsoniter.Stream, val [2]float64) {
    stream.WriteArrayHead()
    stream.WriteFloat64(val[0])
    stream.WriteMore()
    stream.WriteFloat64(val[1])
    stream.WriteArrayTail()
}
