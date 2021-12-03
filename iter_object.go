package jsoniter

// ReadObjectHead tells if there is object field to read
func (iter *Iterator) ReadObjectHead() bool {
	c := iter.nextToken()
	if c != '{' {
		if c == 'n' {
			iter.skipThreeBytes('u', 'l', 'l') // null
			return false
		}
		iter.ReportError("ReadArrayHead", "expect {, but found "+string([]byte{c}))
		iter.unreadByte()
		iter.Skip()
		return false
	}
	c = iter.nextToken()
	if c == '}' {
		// {} is empty object
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

// ReadObjectField will return field name or empty string if not found
func (iter *Iterator) ReadObjectField() string {
	var field string
	if iter.ReadString(&field) != nil {
		return ""
	}
	c := iter.nextToken()
	if c != ':' {
		iter.ReportError("ReadObjectField", "expect :, but found "+string([]byte{c}))
		return ""
	}
	return field
}
