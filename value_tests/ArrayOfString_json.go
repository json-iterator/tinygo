package main

import jsoniter "github.com/json-iterator/tinygo"

func jd_array_string(iter *jsoniter.Iterator, out *[]string) {
  if iter.Error != nil { return }
  i := 0
  val := *out
  for iter.ReadArray() {
    elem := iter.ReadString()
    if i < len(val) {
      if iter.Error == nil {
        val[i] = elem
      }
    } else {
      val = append(val, elem)
    }
    i++
  }
  if i == 0 {
    *out = []string{}
  } else {
    *out = val[:i]
  }
}
