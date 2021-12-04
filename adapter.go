package jsoniter

import (
	"fmt"
	"reflect"
)

type TypeAdapter interface {
	Type() interface{}
	Unmarshal(iter *Iterator, out interface{})
}

type JsonAdapter struct {
	Unmarshal func(bytes []byte, out interface{}) error
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
			adapter := adapterMap[t]
			if adapter == nil {
				return fmt.Errorf("unknown type: %s", t)
			}
			iter := ParseBytes(bytes)
			adapter.Unmarshal(iter, out)
			return iter.Error
		},
	}
}
