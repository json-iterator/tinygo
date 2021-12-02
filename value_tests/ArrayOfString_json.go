package main

import jsoniter "github.com/json-iterator/tinygo"

func jd_array_string(iter *jsoniter.Iterator) []string {
    var val = []string{}
    if iter.Error != nil { return val }
    for iter.ReadArray() {
        val = append(val, iter.ReadString())
    }
    return val
}
