package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func AnonymousArray_json_unmarshal(iter *jsoniter.Iterator, out *AnonymousArray) {
array1_json_unmarshal := func (iter *jsoniter.Iterator, out *[]string) {
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
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    switch {
    case field == `Value`:
      array1_json_unmarshal(iter, &(*out).Value)
    default:
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}