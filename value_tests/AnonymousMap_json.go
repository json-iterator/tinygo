package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func AnonymousMap_json_unmarshal(iter *jsoniter.Iterator, out *AnonymousMap) {
map1_json_unmarshal := func (iter *jsoniter.Iterator, out *map[string]string) {
  more := iter.ReadObjectHead()
  if *out == nil && iter.Error == nil {
    *out = make(map[string]string)
  }
  for more {
    field := iter.ReadObjectField()
    var value string
    var key string
    var err error
    key = field
    iter.ReadString(&value)
    if err != nil {
      iter.ReportError("read map key", err.Error())
    } else {
      (*out)[key] = value
    }
    more = iter.ReadObjectMore()
  }
}
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    switch {
    case field == `Value`:
    map1_json_unmarshal(iter, &(*out).Value)
    default:
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
type AnonymousMap_json struct {
}
func (json AnonymousMap_json) Type() interface{} {
  var val AnonymousMap
  return &val
}
func (json AnonymousMap_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  AnonymousMap_json_unmarshal(iter, val.(*AnonymousMap))
}
