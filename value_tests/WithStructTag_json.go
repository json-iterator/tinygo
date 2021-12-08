package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func WithStructTag_json_unmarshal(iter *jsoniter.Iterator, out *WithStructTag) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    switch {
    case field == `field1`:
    iter.ReadString(&(*out).Field1)
    default:
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
type WithStructTag_json struct {
}
func (json WithStructTag_json) Type() interface{} {
  var val WithStructTag
  return &val
}
func (json WithStructTag_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  WithStructTag_json_unmarshal(iter, val.(*WithStructTag))
}
