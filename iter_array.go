package jsoniter

// ReadArray read array element, tells if the array has more element to read.
func (iter *Iterator) ReadArray() (ret bool) {
	c := iter.nextToken()
	switch c {
	case '[':
		c = iter.nextToken()
		if c != ']' {
			iter.unreadByte()
			return true
		}
		return false
	case ']':
		return false
	case ',':
		return true
	default:
		iter.ReportError("ReadArray", "expect [ or , or ], but found "+string([]byte{c}))
		return
	}
}
