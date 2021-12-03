package jsoniter

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (iter *Iterator) ReadObject() (ret string) {
	c := iter.nextToken()
	switch c {
	case '{':
		c = iter.nextToken()
		if c == '"' {
			iter.unreadByte()
			field := iter.ReadString()
			c = iter.nextToken()
			if c != ':' {
				iter.ReportError("ReadObject", "expect : after object field, but found "+string([]byte{c}))
			}
			return field
		}
		if c == '}' {
			return "" // end of object
		}
		iter.ReportError("ReadObject", `expect " after {, but found `+string([]byte{c}))
		return
	case ',':
		field := iter.ReadString()
		c = iter.nextToken()
		if c != ':' {
			iter.ReportError("ReadObject", "expect : after object field, but found "+string([]byte{c}))
		}
		return field
	case '}':
		return "" // end of object
	default:
		iter.ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{c})))
		return
	}
}

// CaseInsensitive
func (iter *Iterator) readFieldHash() int64 {
	hash := int64(0x811c9dc5)
	c := iter.nextToken()
	if c != '"' {
		iter.ReportError("readFieldHash", `expect ", but found `+string([]byte{c}))
		return 0
	}
	for {
		for i := iter.head; i < len(iter.buf); i++ {
			// require ascii string and no escape
			b := iter.buf[i]
			if b == '\\' {
				iter.head = i
				for _, b := range iter.ReadString() {
					if 'A' <= b && b <= 'Z' {
						b += 'a' - 'A'
					}
					hash ^= int64(b)
					hash *= 0x1000193
				}
				c = iter.nextToken()
				if c != ':' {
					iter.ReportError("readFieldHash", `expect :, but found `+string([]byte{c}))
					return 0
				}
				return hash
			}
			if b == '"' {
				iter.head = i + 1
				c = iter.nextToken()
				if c != ':' {
					iter.ReportError("readFieldHash", `expect :, but found `+string([]byte{c}))
					return 0
				}
				return hash
			}
			if 'A' <= b && b <= 'Z' {
				b += 'a' - 'A'
			}
			hash ^= int64(b)
			hash *= 0x1000193
		}
		iter.ReportError("readFieldHash", `incomplete field name`)
		return 0
	}
}

func calcHash(str string, caseSensitive bool) int64 {
	if !caseSensitive {
		str = strings.ToLower(str)
	}
	hash := int64(0x811c9dc5)
	for _, b := range []byte(str) {
		hash ^= int64(b)
		hash *= 0x1000193
	}
	return int64(hash)
}

func (iter *Iterator) readObjectStart() bool {
	c := iter.nextToken()
	if c == '{' {
		c = iter.nextToken()
		if c == '}' {
			return false
		}
		iter.unreadByte()
		return true
	} else if c == 'n' {
		iter.skipThreeBytes('u', 'l', 'l')
		return false
	}
	iter.ReportError("readObjectStart", "expect { or n, but found "+string([]byte{c}))
	return false
}
