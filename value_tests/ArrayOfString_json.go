package main

import jsoniter "github.com/json-iterator/tinygo"

func jd_array_string(iter *jsoniter.Iterator, out *[]string) {
  if iter.Error != nil { return }
  i := 0
  val := *out
  for iter.ReadArray() {
    if iter.WhatIsNext() == jsoniter.StringValue {
      if i == len(val) {
        val = append(val, iter.ReadString())
      } else {
        val[i] = iter.ReadString()
      }
    } else {
      iter.ReportError("decode array", "expect string")
      iter.Skip()
      if i == len(val) {
        val = append(val, "")
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
