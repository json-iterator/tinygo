package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func IntKeyMap_json_unmarshal(iter *jsoniter.Iterator, out *IntKeyMap) {
  more := iter.ReadObjectHead()
  if *out == nil && iter.Error == nil {
    *out = make(map[int]string)
  }
  for more {
    field := iter.ReadObjectField()
    var value string
    var key int
    var err error
    err = jsoniter.ParseBytes([]byte(field)).ReadInt(&key)
    iter.ReadString(&value)
    if err != nil {
      iter.ReportError("read map key", err.Error())
    } else {
      (*out)[key] = value
    }
    more = iter.ReadObjectMore()
  }
}
func IntKeyMap_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *IntKeyMap) bool {
  return false
}
type IntKeyMap_json struct {
}
func (json IntKeyMap_json) Type() interface{} {
  var val IntKeyMap
  return &val
}
func (json IntKeyMap_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  IntKeyMap_json_unmarshal(iter, val.(*IntKeyMap))
}
