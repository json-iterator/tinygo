package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func NamedStruct_json_unmarshal(iter *jsoniter.Iterator, out *NamedStruct) {
  if iter.Error != nil { return }
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
