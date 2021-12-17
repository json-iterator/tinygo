package value_tests

import jsoniter "github.com/json-iterator/tinygo"

type NamedMap_json struct {
}
func (json NamedMap_json) Type() interface{} {
  var val NamedMap
  return val
}
func (json NamedMap_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  NamedMap_json_unmarshal(iter, out.(*NamedMap))
}
func (json NamedMap_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  NamedMap_json_marshal(stream, val.(NamedMap))
}
func NamedMap_json_unmarshal(iter *jsoniter.Iterator, out *NamedMap) {
  more := iter.ReadObjectHead()
  if *out == nil && iter.Error == nil {
    *out = make(map[string]int)
  }
  for more {
    field := iter.ReadObjectField()
    var value int
    var key string
    var err error
    key = field
    iter.ReadInt(&value)
    if err != nil {
      iter.ReportError("read map key", err.Error())
    } else {
      (*out)[key] = value
    }
    more = iter.ReadObjectMore()
  }
}
func NamedMap_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *NamedMap) bool {
  if field == "NamedMap" {
    NamedMap_json_unmarshal(iter, out)
    return true
  }
  return false
}
func NamedMap_json_marshal(stream *jsoniter.Stream, val NamedMap) {
  if len(val) == 0 {
    stream.WriteEmptyObject()
  } else {
    isFirst := true
    stream.WriteObjectHead()
    for k, v := range val {
      if isFirst {
        isFirst = false
      } else {
        stream.WriteMore()
      }
      stream.WriteObjectField(k)
    stream.WriteInt(v)
    }
    stream.WriteObjectTail()
  }
}
