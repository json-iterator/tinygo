package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func WithStructTag_json_unmarshal(iter *jsoniter.Iterator, out *WithStructTag) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !WithStructTag_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func WithStructTag_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *WithStructTag) bool {
  switch {
  case field == `field1`:
    iter.ReadString(&(*out).Field1)
    return true
  }
  return false
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
