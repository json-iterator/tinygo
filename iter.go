package jsoniter

import (
	"fmt"
	"strconv"
)

type RawMessage []byte

// A Number represents a JSON number literal.
type Number string

// String returns the literal text of the number.
func (n Number) String() string { return string(n) }

// Float64 returns the number as a float64.
func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

// Int64 returns the number as an int64.
func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

const uint32SafeToMultiply10 = uint32(0xffffffff)/10 - 1
const uint64SafeToMultiple10 = uint64(0xffffffffffffffff)/10 - 1
const maxFloat64 = 1<<53 - 1

var pow10 []uint64
var hexDigits []byte

func init() {
	pow10 = []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}
	hexDigits = make([]byte, 256)
	for i := 0; i < len(hexDigits); i++ {
		hexDigits[i] = 255
	}
	for i := '0'; i <= '9'; i++ {
		hexDigits[i] = byte(i - '0')
	}
	for i := 'a'; i <= 'f'; i++ {
		hexDigits[i] = byte((i - 'a') + 10)
	}
	for i := 'A'; i <= 'F'; i++ {
		hexDigits[i] = byte((i - 'A') + 10)
	}
}

// Iterator is a io.Reader like object, with JSON specific read functions.
// Error is not returned as return value, but stored as Error member on this iterator instance.
type Iterator struct {
	buf   []byte
	head  int
	Error error
}

// ParseBytes creates an Iterator instance from byte array
func ParseBytes(input []byte) *Iterator {
	return &Iterator{
		buf:  input,
		head: 0,
	}
}

func (iter *Iterator) skipWhitespaces() {
	for i := iter.head; i < len(iter.buf); i++ {
		c := iter.buf[i]
		switch c {
		case ' ', '\n', '\t', '\r':
			continue
		}
		iter.head = i
		return
	}
	iter.head = len(iter.buf)
}

// ReportError record a error in iterator instance with current position.
func (iter *Iterator) ReportError(operation string, msg string) error {
	if iter.Error != nil {
		return iter.Error
	}
	peekStart := iter.head - 10
	if peekStart < 0 {
		peekStart = 0
	}
	peekEnd := iter.head + 10
	if peekEnd > len(iter.buf) {
		peekEnd = len(iter.buf)
	}
	parsing := string(iter.buf[peekStart:peekEnd])
	contextStart := iter.head - 50
	if contextStart < 0 {
		contextStart = 0
	}
	contextEnd := iter.head + 50
	if contextEnd > len(iter.buf) {
		contextEnd = len(iter.buf)
	}
	context := string(iter.buf[contextStart:contextEnd])
	iter.Error = fmt.Errorf("%s: %s, error found in #%v byte of ...|%s|..., bigger context ...|%s|...",
		operation, msg, iter.head-peekStart, parsing, context)
	return iter.Error
}

func (iter *Iterator) reportError(err error) {
	if iter.Error != nil {
		return
	}
	iter.Error = err
}

// CurrentBuffer gets current buffer as string for debugging purpose
func (iter *Iterator) CurrentBuffer() string {
	peekStart := iter.head - 10
	if peekStart < 0 {
		peekStart = 0
	}
	return fmt.Sprintf("parsing #%v byte, around ...|%s|..., whole buffer ...|%s|...", iter.head,
		string(iter.buf[peekStart:iter.head]), string(iter.buf[0:len(iter.buf)]))
}

func (iter *Iterator) readByte() (ret byte) {
	if iter.head < len(iter.buf) {
		ret = iter.buf[iter.head]
		iter.head++
		return ret
	}
	iter.ReportError("readByte", "EOF")
	return 0
}

func (iter *Iterator) unreadByte() {
	if iter.head > 0 {
		iter.head--
	}
}

func (iter *Iterator) nextToken() byte {
	// a variation of skip whitespaces, returning the next non-whitespace token
	for {
		for i := iter.head; i < len(iter.buf); i++ {
			c := iter.buf[i]
			switch c {
			case ' ', '\n', '\t', '\r':
				continue
			}
			iter.head = i + 1
			return c
		}
		return 0
	}
}
