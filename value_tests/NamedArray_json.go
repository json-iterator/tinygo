package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func NamedArray_json_unmarshal(iter *jsoniter.Iterator, out *NamedArray) {
  if iter.Error != nil { return }
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
