package main

import jsoniter "github.com/json-iterator/tinygo"

func jd_array_string(iter *jsoniter.Iterator, out *[]string) {
  if iter.Error != nil { return }
  i := 0
  val := *out
  for iter.ReadArray() {
    if iter.AssertIsString() && !iter.SkipNull() {
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
  }
  if i == 0 {
    *out = []string{}
  } else {
    *out = val[:i]
  }
}
