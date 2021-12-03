package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func jd_StructOfStringInt(iter *jsoniter.Iterator, out *StructOfStringInt) {
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
