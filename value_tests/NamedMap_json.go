package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func NamedMap_json_unmarshal(iter *jsoniter.Iterator, out *NamedMap) {
  more := iter.ReadObjectHead()
  if *out == nil && iter.Error == nil {
    *out = make(map[string]int)
  }
  for more {
    field := iter.ReadObjectField()
    var value int
    iter.ReadInt(&value)
    (*out)[field] = value
    more = iter.ReadObjectMore()
  }
}
type NamedMap_json struct {
}
func (json NamedMap_json) Type() interface{} {
  var val NamedMap
  return &val
}
func (json NamedMap_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  NamedMap_json_unmarshal(iter, val.(*NamedMap))
}
