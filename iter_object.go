package jsoniter

import (
	"fmt"
)

// call in this sequence AssertIsObject => ReadObjectHead => ReadObjectMore

// AssertIsObject must be called before ReadObjectHead, if the value is not array will be skipped
func (iter *Iterator) AssertIsObject() bool {
	c := iter.nextToken()
	switch c {
	case '{':
		return true
	case '"':
		iter.ReportError("AssertIsObject", "unexpected string")
	case 'n':
		// null is not considered as object, but not a error
		iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		iter.ReportError("AssertIsObject", "unexpected boolean")
		iter.skipThreeBytes('r', 'u', 'e') // true
	case 'f':
		iter.ReportError("AssertIsObject", "unexpected boolean")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		iter.ReportError("AssertIsObject", "unexpected number")
		iter.skipNumber()
	case '[':
		iter.ReportError("AssertIsObject", "unexpected object")
		iter.skipObject()
	default:
		iter.ReportError("AssertIsObject", fmt.Sprintf("unknown data: %v", c))
	}
	return false
}

// ReadObjectHead tells if there is object field to read
func (iter *Iterator) ReadObjectHead() bool {
	c := iter.nextToken()
	if c == '}' {
		return false
	}
	iter.unreadByte()
	return true
}

// ReadObjectMore tells if there is more field to read
func (iter *Iterator) ReadObjectMore() bool {
	c := iter.nextToken()
	switch c {
	case ',':
		return true
	case '}':
		return false
	default:
		iter.ReportError("ReadObjectMore", "expect , or }, but found "+string([]byte{c}))
		return false
	}
}
