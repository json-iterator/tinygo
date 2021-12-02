package jsoniter

import (
	"fmt"
	"io"
)

// ValueType the type for JSON element
type ValueType int

type RawMessage []byte

const (
	// InvalidValue invalid JSON element
	InvalidValue ValueType = iota
	// StringValue JSON element "string"
	StringValue
	// NumberValue JSON element 100 or 0.10
	NumberValue
	// NullValue JSON element null
	NullValue
	// BoolValue JSON element true or false
	BoolValue
	// ArrayValue JSON element []
	ArrayValue
	// ObjectValue JSON element {}
	ObjectValue
)
const uint32SafeToMultiply10 = uint32(0xffffffff)/10 - 1
const uint64SafeToMultiple10 = uint64(0xffffffffffffffff)/10 - 1
const maxFloat64 = 1<<53 - 1

var pow10 []uint64
var hexDigits []byte
var valueTypes []ValueType

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
	valueTypes = make([]ValueType, 256)
	for i := 0; i < len(valueTypes); i++ {
		valueTypes[i] = InvalidValue
	}
	valueTypes['"'] = StringValue
	valueTypes['-'] = NumberValue
	valueTypes['0'] = NumberValue
	valueTypes['1'] = NumberValue
	valueTypes['2'] = NumberValue
	valueTypes['3'] = NumberValue
	valueTypes['4'] = NumberValue
	valueTypes['5'] = NumberValue
	valueTypes['6'] = NumberValue
	valueTypes['7'] = NumberValue
	valueTypes['8'] = NumberValue
	valueTypes['9'] = NumberValue
	valueTypes['t'] = BoolValue
	valueTypes['f'] = BoolValue
	valueTypes['n'] = NullValue
	valueTypes['['] = ArrayValue
	valueTypes['{'] = ObjectValue
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

// WhatIsNext gets ValueType of relatively next json element
func (iter *Iterator) WhatIsNext() ValueType {
	valueType := valueTypes[iter.nextToken()]
	iter.unreadByte()
	return valueType
}

func (iter *Iterator) skipWhitespacesWithoutLoadMore() bool {
	for i := iter.head; i < len(iter.buf); i++ {
		c := iter.buf[i]
		switch c {
		case ' ', '\n', '\t', '\r':
			continue
		}
		iter.head = i
		return false
	}
	return true
}

func (iter *Iterator) isObjectEnd() bool {
	c := iter.nextToken()
	if c == ',' {
		return false
	}
	if c == '}' {
		return true
	}
	iter.ReportError("isObjectEnd", "object ended prematurely, unexpected char "+string([]byte{c}))
	return true
}

// ReportError record a error in iterator instance with current position.
func (iter *Iterator) ReportError(operation string, msg string) {
	if iter.Error != nil {
		if iter.Error != io.EOF {
			return
		}
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
	return 0
}

func (iter *Iterator) unreadByte() {
	iter.head--
	return
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
