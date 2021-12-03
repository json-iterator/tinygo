package jsoniter

import "fmt"

func (iter *Iterator) skipNumber() {
	for i := iter.head; i < len(iter.buf); i++ {
		c := iter.buf[i]
		switch c {
		case ' ', '\n', '\r', '\t', ',', '}', ']':
			iter.head = i
			return
		}
	}
	iter.head = len(iter.buf)
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

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
func (iter *Iterator) skipString() {
	escaped := false
	for i := iter.head; i < len(iter.buf); i++ {
		c := iter.buf[i]
		if c == '"' {
			if !escaped {
				iter.head = i + 1
				return
			}
			j := i - 1
			for {
				if j < iter.head || iter.buf[j] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					iter.head = i + 1
					return
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
	iter.ReportError("skipString", "incomplete string")
	iter.head = len(iter.buf)
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
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
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

func (iter *Iterator) skipFourBytes(b1, b2, b3, b4 byte) error {
	if iter.readByte() != b1 {
		return iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
	}
	if iter.readByte() != b2 {
		return iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
	}
	if iter.readByte() != b3 {
		return iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
	}
	if iter.readByte() != b4 {
		return iter.ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3, b4})))
	}
	return nil
}

func (iter *Iterator) skipThreeBytes(b1, b2, b3 byte) error {
	if iter.readByte() != b1 {
		return iter.ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3})))
	}
	if iter.readByte() != b2 {
		return iter.ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3})))
	}
	if iter.readByte() != b3 {
		return iter.ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{b1, b2, b3})))
	}
	return nil
}
