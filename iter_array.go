package jsoniter

// ReadArrayHead tells if there is array element to read, if the value is not array will be skipped
func (iter *Iterator) ReadArrayHead() bool {
	c := iter.nextToken()
	if c != '[' {
		if c == 'n' {
			iter.skipThreeBytes('u', 'l', 'l') // null
			return false
		}
		iter.ReportError("ReadArrayHead", "expect [, but found "+string([]byte{c}))
		iter.skip(c)
		return false
	}
	c = iter.nextToken()
	if c == ']' {
		// [] is empty array
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
