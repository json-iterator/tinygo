package jsoniter

import (
	"fmt"
	"unicode/utf16"
)

// ReadString will assign string to out if found, otherwise the value will be skipped
func (iter *Iterator) ReadString(out *string) error {
	c := iter.nextToken()
	if c != '"' {
		if c == 'n' {
			iter.skipThreeBytes('u', 'l', 'l') // null
			return nil
		}
		err := iter.ReportError("ReadString", `expects ", but found `+string([]byte{c}))
		iter.skip(c)
		return err
	}
	return iter.readString(out)
}

func (iter *Iterator) readString(out *string) error {
	for i := iter.head; i < len(iter.buf); i++ {
		c := iter.buf[i]
		if c == '"' {
			*out = string(iter.buf[iter.head:i])
			iter.head = i + 1
			return nil
		} else if c == '\\' {
			return iter.readStringSlowPath(out)
		} else if c < ' ' {
			return iter.ReportError("ReadString",
				fmt.Sprintf(`invalid control character found: %d`, c))
		}
	}
	return iter.ReportError("ReadString", `missing "`)
}

func (iter *Iterator) readStringSlowPath(out *string) error {
	var str []byte
	var c byte
	for iter.head < len(iter.buf) {
		c = iter.readByte()
		if c == '"' {
			*out = string(str)
			return nil
		}
		if c == '\\' {
			c = iter.readByte()
			str = iter.readEscapedChar(c, str)
		} else {
			str = append(str, c)
		}
	}
	return iter.ReportError("readStringSlowPath", "unexpected end of input")
}

func (iter *Iterator) readEscapedChar(c byte, str []byte) []byte {
	switch c {
	case 'u':
		r := iter.readU4()
		if utf16.IsSurrogate(r) {
			c = iter.readByte()
			if c != '\\' {
				iter.unreadByte()
				str = appendRune(str, r)
				return str
			}
			c = iter.readByte()
			if c != 'u' {
				str = appendRune(str, r)
				return iter.readEscapedChar(c, str)
			}
			r2 := iter.readU4()
			combined := utf16.DecodeRune(r, r2)
			if combined == '\uFFFD' {
				str = appendRune(str, r)
				str = appendRune(str, r2)
			} else {
				str = appendRune(str, combined)
			}
		} else {
			str = appendRune(str, r)
		}
	case '"':
		str = append(str, '"')
	case '\\':
		str = append(str, '\\')
	case '/':
		str = append(str, '/')
	case 'b':
		str = append(str, '\b')
	case 'f':
		str = append(str, '\f')
	case 'n':
		str = append(str, '\n')
	case 'r':
		str = append(str, '\r')
	case 't':
		str = append(str, '\t')
	default:
		iter.ReportError("readEscapedChar",
			`invalid escape char after \`)
		return nil
	}
	return str
}

func (iter *Iterator) readU4() (ret rune) {
	for i := 0; i < 4; i++ {
		c := iter.readByte()
		if c >= '0' && c <= '9' {
			ret = ret*16 + rune(c-'0')
		} else if c >= 'a' && c <= 'f' {
			ret = ret*16 + rune(c-'a'+10)
		} else if c >= 'A' && c <= 'F' {
			ret = ret*16 + rune(c-'A'+10)
		} else {
			iter.ReportError("readU4", "expects 0~9 or a~f, but found "+string([]byte{c}))
			return
		}
	}
	return ret
}

const (
	t1 = 0x00 // 0000 0000
	tx = 0x80 // 1000 0000
	t2 = 0xC0 // 1100 0000
	t3 = 0xE0 // 1110 0000
	t4 = 0xF0 // 1111 0000
	t5 = 0xF8 // 1111 1000

	maskx = 0x3F // 0011 1111
	mask2 = 0x1F // 0001 1111
	mask3 = 0x0F // 0000 1111
	mask4 = 0x07 // 0000 0111

	rune1Max = 1<<7 - 1
	rune2Max = 1<<11 - 1
	rune3Max = 1<<16 - 1

	surrogateMin = 0xD800
	surrogateMax = 0xDFFF

	maxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
	runeError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
)

func appendRune(p []byte, r rune) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch i := uint32(r); {
	case i <= rune1Max:
		p = append(p, byte(r))
		return p
	case i <= rune2Max:
		p = append(p, t2|byte(r>>6))
		p = append(p, tx|byte(r)&maskx)
		return p
	case i > maxRune, surrogateMin <= i && i <= surrogateMax:
		r = runeError
		fallthrough
	case i <= rune3Max:
		p = append(p, t3|byte(r>>12))
		p = append(p, tx|byte(r>>6)&maskx)
		p = append(p, tx|byte(r)&maskx)
		return p
	default:
		p = append(p, t4|byte(r>>18))
		p = append(p, tx|byte(r>>12)&maskx)
		p = append(p, tx|byte(r>>6)&maskx)
		p = append(p, tx|byte(r)&maskx)
		return p
	}
}
