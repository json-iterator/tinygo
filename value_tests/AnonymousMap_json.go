package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type AnonymousMap_json struct {
}
func (json AnonymousMap_json) Type() interface{} {
  var val AnonymousMap
  return val
}
func (json AnonymousMap_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  AnonymousMap_json_unmarshal(iter, out.(*AnonymousMap))
}
func (json AnonymousMap_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  AnonymousMap_json_marshal(stream, val.(AnonymousMap))
}
func AnonymousMap_json_unmarshal(iter *jsoniter.Iterator, out *AnonymousMap) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !AnonymousMap_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func AnonymousMap_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *AnonymousMap) bool {
  switch {
  case field == `Value`:
    AnonymousMap_map1_json_unmarshal(iter, &(*out).Value)
    return true
  }
  return false
}
func AnonymousMap_map1_json_unmarshal (iter *jsoniter.Iterator, out *map[string]string) {
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
func AnonymousMap_json_marshal(stream *jsoniter.Stream, val AnonymousMap) {
    stream.WriteObjectHead()
    AnonymousMap_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func AnonymousMap_json_marshal_field(stream *jsoniter.Stream, val AnonymousMap) {
    stream.WriteObjectField(`Value`)
    AnonymousMap_map2_json_marshal(stream, val.Value)
    stream.WriteMore()
}
func AnonymousMap_map2_json_marshal (stream *jsoniter.Stream, val map[string]string) {
  stream.WriteObjectHead()
  for k, v := range val {
    stream.WriteObjectField(k)
    stream.WriteString(v)
    stream.WriteMore()
  }
  stream.WriteObjectTail()
}
