package jsoniter

import "fmt"

// call in this sequence AssertIsBool => ReadBool

// AssertIsBool tell if the value is bool, otherwise report error and skip
func (iter *Iterator) AssertIsBool() bool {
	c := iter.nextToken()
	switch c {
	case '"':
		iter.ReportError("AssertIsBool", "unexpected string")
		iter.skipString()
	case 'n':
		// null is not considered as bool, but not a error
		iter.skipThreeBytes('u', 'l', 'l')
	case 't':
		return true
	case 'f':
		return true
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		iter.ReportError("AssertIsNumber", "unexpected array")
		iter.skipNumber()
	case '[':
		iter.ReportError("AssertIsNumber", "unexpected array")
		iter.skipArray()
	case '{':
		iter.ReportError("AssertIsNumber", "unexpected object")
		iter.skipObject()
	default:
		iter.ReportError("AssertIsNumber", fmt.Sprintf("unknown data: %v", c))
	}
	return false
}

// ReadBool reads a json value as bool
func (iter *Iterator) ReadBool() (ret bool) {
	c := iter.readByte()
	if c == 'r' {
		iter.skipTwoBytes('u', 'e')
		return true
	}
	if c == 'a' {
		iter.skipThreeBytes('l', 's', 'e')
		return false
	}
	iter.ReportError("ReadBool", "expect r or a, but found "+string([]byte{c}))
	return
}
