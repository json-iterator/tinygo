package jsoniter

import "fmt"

func (iter *Iterator) skipNumber() {
	for {
		for i := iter.head; i < len(iter.buf); i++ {
			c := iter.buf[i]
			switch c {
			case ' ', '\n', '\r', '\t', ',', '}', ']':
				iter.head = i
				return
			}
		}
	}
}

func (iter *Iterator) skipArray() {
	level := 1
	for {
		for i := iter.head; i < len(iter.buf); i++ {
			switch iter.buf[i] {
			case '"': // If inside string, skip it
				iter.head = i + 1
				iter.skipString()
				i = iter.head - 1 // it will be i++ soon
			case '[': // If open symbol, increase level
				level++
			case ']': // If close symbol, increase level
				level--
				// If we have returned to the original level, we're done
				if level == 0 {
					iter.head = i + 1
					return
				}
			}
		}
	}
}

func (iter *Iterator) skipObject() {
	level := 1
	for {
		for i := iter.head; i < len(iter.buf); i++ {
			switch iter.buf[i] {
			case '"': // If inside string, skip it
				iter.head = i + 1
				iter.skipString()
				i = iter.head - 1 // it will be i++ soon
			case '{': // If open symbol, increase level
				level++
			case '}': // If close symbol, increase level
				level--
				// If we have returned to the original level, we're done
				if level == 0 {
					iter.head = i + 1
					return
				}
			}
		}
	}
}

func (iter *Iterator) skipString() {
	end, _ := iter.findStringEnd()
	if end == -1 {
		iter.ReportError("skipString", "incomplete string")
		return
	} else {
		iter.head = end
		return
	}
}

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
// Tries to find the end of string
// Support if string contains escaped quote symbols.
func (iter *Iterator) findStringEnd() (int, bool) {
	escaped := false
	for i := iter.head; i < len(iter.buf); i++ {
		c := iter.buf[i]
		if c == '"' {
			if !escaped {
				return i + 1, false
			}
			j := i - 1
			for {
				if j < iter.head || iter.buf[j] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					return i + 1, true
				}
				j--
				if j < iter.head || iter.buf[j] != '\\' {
					// odd number of backslashes
					// it is \" or \\\"
					break
				}
				j--
			}
		} else if c == '\\' {
			escaped = true
		}
	}
	j := len(iter.buf) - 1
	for {
		if j < iter.head || iter.buf[j] != '\\' {
			// even number of backslashes
			// either end of buffer, or " found
			return -1, false // do not end with \
		}
		j--
		if j < iter.head || iter.buf[j] != '\\' {
			// odd number of backslashes
			// it is \" or \\\"
			break
		}
		j--

	}
	return -1, true // end with \
}

// ReadBool reads a json object as BoolValue
func (iter *Iterator) ReadBool() (ret bool) {
	c := iter.readByte()
	if c == 't' {
		iter.skipThreeBytes('r', 'u', 'e')
		return true
	}
	if c == 'f' {
		iter.skipFourBytes('a', 'l', 's', 'e')
		return false
	}
	iter.ReportError("ReadBool", "expect t or f, but found "+string([]byte{c}))
	return
}

// Skip skips a json object and positions to relatively the next json object
func (iter *Iterator) Skip() {
	c := iter.readByte()
	switch c {
	case '"':
		iter.skipString()
	case 'n':
		iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		iter.skipThreeBytes('r', 'u', 'e') // true
	case 'f':
		iter.skipFourBytes('a', 'l', 's', 'e') // false
	case '0':
		iter.unreadByte()
		iter.ReadFloat32()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		iter.skipNumber()
	case '[':
		iter.skipArray()
	case '{':
		iter.skipObject()
	default:
		iter.ReportError("Skip", fmt.Sprintf("do not know how to skip: %v", c))
		return
	}
}

func (iter *Iterator) skipFourBytes(b1, b2, b3, b4 byte) {
	if iter.readByte() != b1 {
		iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
		return
	}
	if iter.readByte() != b2 {
		iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
		return
	}
	if iter.readByte() != b3 {
		iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
		return
	}
	if iter.readByte() != b4 {
		iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
		return
	}
}

func (iter *Iterator) skipThreeBytes(b1, b2, b3 byte) {
	if iter.readByte() != b1 {
		iter.ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3})))
		return
	}
	if iter.readByte() != b2 {
		iter.ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3})))
		return
	}
	if iter.readByte() != b3 {
		iter.ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3})))
		return
	}
}
