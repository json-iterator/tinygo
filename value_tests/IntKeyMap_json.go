package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type IntKeyMap_json struct {
}
func (json IntKeyMap_json) Type() interface{} {
  var val IntKeyMap
  return val
}
func (json IntKeyMap_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  IntKeyMap_json_unmarshal(iter, out.(*IntKeyMap))
}
func (json IntKeyMap_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  IntKeyMap_json_marshal(stream, val.(IntKeyMap))
}
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
  if field == "IntKeyMap" {
    IntKeyMap_json_unmarshal(iter, out)
    return true
  }
  return false
}
func IntKeyMap_json_marshal(stream *jsoniter.Stream, val IntKeyMap) {
  stream.WriteObjectHead()
  for k, v := range val {
      stream.WriteRaw("\"")
    stream.WriteInt(k)
      stream.WriteRaw("\": ")
    stream.WriteString(v)
    stream.WriteMore()
  }
  stream.WriteObjectTail()
}
func IntKeyMap_json_marshal_field(stream *jsoniter.Stream, val IntKeyMap) {
    stream.WriteObjectField("IntKeyMap")
    IntKeyMap_json_marshal(stream, val)
    stream.WriteMore()
}
