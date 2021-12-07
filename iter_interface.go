package jsoniter

import (
	"fmt"
)

func (iter *Iterator) ReadInterface(out *interface{}) (ret error) {
	c := iter.nextToken()
	switch c {
	case '"':
		var val string
		ret = iter.readString(&val)
		if ret == nil {
			*out = val
		}
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.skipThreeBytes('r', 'u', 'e') // true
		if ret == nil {
			*out = true
		}
		return
	case 'f':
		ret = iter.skipFourBytes('a', 'l', 's', 'e') // false
		if ret == nil {
			*out = false
		}
		return
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		iter.unreadByte()
		var val float64
		ret = iter.ReadFloat64(&val)
		if ret == nil {
			*out = val
		}
		return
	case '[':
		var val = make([]interface{}, 0)
		c := iter.nextToken()
		if c == ']' {
			*out = val
			return
		}
		iter.unreadByte()
		more := true
		for more {
			var elem interface{}
			iter.ReadInterface(&elem)
			val = append(val, elem)
			more = iter.ReadArrayMore()
		}
		*out = val
		return iter.Error
	case '{':
		val := make(map[string]interface{}, 0)
		*out = val
		c := iter.nextToken()
		if c == ']' {
			return
		}
		iter.unreadByte()
		more := true
		for more {
			k := iter.ReadObjectField()
			var v interface{}
			iter.ReadInterface(&v)
			val[k] = v
			more = iter.ReadObjectMore()
		}
		return iter.Error
	default:
		return iter.ReportError("ReadInterface", fmt.Sprintf("unexpected character: %v", c))
	}
}
