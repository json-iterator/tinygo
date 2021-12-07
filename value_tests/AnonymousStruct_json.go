package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func AnonymousStruct_json_unmarshal(iter *jsoniter.Iterator, out *AnonymousStruct) {
struct1_json_unmarshal := func (iter *jsoniter.Iterator, out *struct {
	Name	string
	Price	int
}) {
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
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    switch {
    case field == `Value`:
    struct1_json_unmarshal(iter, &(*out).Value)
    default:
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
type AnonymousStruct_json struct {
}
func (json AnonymousStruct_json) Type() interface{} {
  var val AnonymousStruct
  return &val
}
func (json AnonymousStruct_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  AnonymousStruct_json_unmarshal(iter, val.(*AnonymousStruct))
}
