package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func jd_ArrayOfString(iter *jsoniter.Iterator, out *ArrayOfString) {
  if iter.Error != nil { return }
  if !iter.AssertIsArray() { return }
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if iter.AssertIsString() {
      if i == len(val) {
        val = append(val, iter.ReadString())
      } else {
        val[i] = iter.ReadString()
      }
    } else if i == len(val) {
      if i == len(val) {
        var empty string
        val = append(val, empty)
      }
    }
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = []string{}
  } else {
    *out = val[:i]
  }
}
