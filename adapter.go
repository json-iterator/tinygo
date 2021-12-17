package jsoniter

import (
	"fmt"
	"reflect"
)

type TypeAdapter interface {
	Type() interface{}
	Unmarshal(iter *Iterator, out interface{})
	Marshal(stream *Stream, val interface{})
}

type JsonAdapter struct {
	Unmarshal     func(bytes []byte, out interface{}) error
	Marshal       func(val interface{}) ([]byte, error)
	MarshalIndent func(val interface{}, prefix, indent string) ([]byte, error)
}

func CreateJsonAdapter(adapters ...TypeAdapter) JsonAdapter {
	adapterMap := make(map[reflect.Type]TypeAdapter)
	for _, adapter := range adapters {
		t := reflect.ValueOf(adapter.Type()).Type()
		adapterMap[t] = adapter
	}
	return JsonAdapter{
		Unmarshal: func(bytes []byte, out interface{}) error {
			t := reflect.ValueOf(out).Type()
			adapter := adapterMap[t.Elem()]
			if adapter == nil {
				if t.Kind() != reflect.Ptr {
					return fmt.Errorf("unmarshal expect pointer, actual type is: %s", t)
				}
				return fmt.Errorf("unknown type: %s", t)
			}
			iter := ParseBytes(bytes)
			adapter.Unmarshal(iter, out)
			return iter.Error
		},
		Marshal: func(val interface{}) ([]byte, error) {
			t := reflect.ValueOf(val).Type()
			adapter := adapterMap[t]
			if adapter == nil {
				fmt.Println("!!!", adapterMap)
				fmt.Println("!!!", t)
				return nil, fmt.Errorf("unknown type: %s", t)
			}
			stream := NewStream()
			adapter.Marshal(stream, val)
			return stream.Buffer(), stream.Error
		},
		MarshalIndent: func(val interface{}, prefix, indent string) ([]byte, error) {
			t := reflect.ValueOf(val).Type()
			adapter := adapterMap[t]
			if adapter == nil {
				return nil, fmt.Errorf("unknown type: %s", t)
			}
			stream := NewStream()
			stream.Prefix = prefix
			stream.Indent = indent
			adapter.Marshal(stream, val)
			return stream.Buffer(), stream.Error
		},
	}
}
