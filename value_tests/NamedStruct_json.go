package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func NamedStruct_json_unmarshal(iter *jsoniter.Iterator, out *NamedStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    switch {
    case field == `Name`:
      iter.ReadString(&(*out).Name)
    case field == `Price`:
      iter.ReadInt(&(*out).Price)
    default:
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
type NamedStruct_json struct {
}
func (json NamedStruct_json) Type() interface{} {
  var val NamedStruct
  return &val
}
func (json NamedStruct_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  NamedStruct_json_unmarshal(iter, val.(*NamedStruct))
}
