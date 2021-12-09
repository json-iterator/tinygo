package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func EmbedNumberStruct_json_unmarshal(iter *jsoniter.Iterator, out *EmbedNumberStruct) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !EmbedNumberStruct_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func EmbedNumberStruct_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *EmbedNumberStruct) bool {
  if field == "Number" {
    out.Number = new(jsoniter.Number)
    iter.ReadNumber(out.Number)
    return true
  }
  return false
}
type EmbedNumberStruct_json struct {
}
func (json EmbedNumberStruct_json) Type() interface{} {
  var val EmbedNumberStruct
  return &val
}
func (json EmbedNumberStruct_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {
  EmbedNumberStruct_json_unmarshal(iter, val.(*EmbedNumberStruct))
}
