package jsoniter

import "fmt"

// call in this sequence AssertIsArray => ReadArrayHead => ReadArrayMore

// AssertIsArray must be called before ReadArrayHead, if the value is not array will be skipped
func (iter *Iterator) AssertIsArray() bool {
	c := iter.nextToken()
	switch c {
	case '[':
		return true
	case '"':
		iter.ReportError("AssertIsArray", "unexpected string")
	case 'n':
		// null is not considered as array, but not a error
		iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		iter.ReportError("AssertIsArray", "unexpected boolean")
		iter.skipThreeBytes('r', 'u', 'e') // true
	case 'f':
		iter.ReportError("AssertIsArray", "unexpected boolean")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		iter.ReportError("AssertIsArray", "unexpected number")
		iter.skipNumber()
	case '{':
		iter.ReportError("AssertIsArray", "unexpected object")
		iter.skipObject()
	default:
		iter.ReportError("AssertIsArray", fmt.Sprintf("unknown data: %v", c))
	}
	return false
}

// ReadArrayHead tells if there is array element to read
func (iter *Iterator) ReadArrayHead() bool {
	c := iter.nextToken()
	if c == ']' {
		return false
	}
	iter.unreadByte()
	return true
}

// ReadArrayMore tells if there is more element to read
func (iter *Iterator) ReadArrayMore() bool {
	c := iter.nextToken()
	switch c {
	case ',':
		return true
	case ']':
		return false
	default:
		iter.ReportError("ReadArrayMore", "expect , or ], but found "+string([]byte{c}))
		return false
	}
}
