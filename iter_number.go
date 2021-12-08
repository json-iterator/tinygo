package jsoniter

import (
	"math"
	"math/big"
	"strconv"
	"unsafe"
)

var intDigits []int8

const invalidCharForNumber = int8(-1)

func init() {
	intDigits = make([]int8, 256)
	for i := 0; i < len(intDigits); i++ {
		intDigits[i] = invalidCharForNumber
	}
	for i := int8('0'); i <= int8('9'); i++ {
		intDigits[i] = i - int8('0')
	}
}

func (iter *Iterator) ReadNumber(out *Number) error {
	str := iter.readNumberAsString()
	if len(str) == 0 {
		if iter.head < len(iter.buf) && iter.buf[iter.head] == 'n' {
			return iter.skipFourBytes('n', 'u', 'l', 'l') // null
		}
		err := iter.ReportError("ReadNumber", "number not found")
		iter.Skip()
		return err
	}
	*out = Number(str)
	return nil
}

func (iter *Iterator) readNumberAsString() (ret []byte) {
	iter.skipWhitespaces()
	for i := iter.head; i < len(iter.buf); i++ {
		switch iter.buf[i] {
		case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			str := iter.buf[iter.head:i]
			iter.head = i
			return str
		}
	}
	str := iter.buf[iter.head:]
	iter.head = len(iter.buf)
	return str
}

// ReadBigFloat read big.Float
func (iter *Iterator) ReadBigFloat(out *big.Float) error {
	str := iter.readNumberAsString()
	if len(str) == 0 {
		if iter.head < len(iter.buf) && iter.buf[iter.head] == 'n' {
			return iter.skipFourBytes('n', 'u', 'l', 'l') // null
		}
		err := iter.ReportError("ReadBigFloat", "number not found")
		iter.Skip()
		return err
	}
	_, success := (*out).SetString(string(str))
	if !success {
		return iter.ReportError("ReadBigFloat", "invalid big float")
	}
	return nil
}

// ReadBigInt read big.Int
func (iter *Iterator) ReadBigInt(out *big.Int) error {
	str := iter.readNumberAsString()
	if len(str) == 0 {
		if iter.head < len(iter.buf) && iter.buf[iter.head] == 'n' {
			return iter.skipFourBytes('n', 'u', 'l', 'l') // null
		}
		err := iter.ReportError("ReadBigInt", "number not found")
		iter.Skip()
		return err
	}
	_, success := (*out).SetString(string(str), 10)
	if !success {
		return iter.ReportError("ReadBigInt", "invalid big int")
	}
	return nil
}

//ReadFloat32 read float32
func (iter *Iterator) ReadFloat32(out *float32) error {
	str := iter.readNumberAsString()
	if len(str) == 0 {
		if iter.head < len(iter.buf) && iter.buf[iter.head] == 'n' {
			return iter.skipFourBytes('n', 'u', 'l', 'l') // null
		}
		err := iter.ReportError("ReadFloat32", "number not found")
		iter.Skip()
		return err
	}
	val, err := strconv.ParseFloat(string(str), 32)
	if err != nil {
		iter.reportError(err)
		return err
	}
	*out = float32(val)
	return nil
}

// ReadFloat64 read float64
func (iter *Iterator) ReadFloat64(out *float64) error {
	str := iter.readNumberAsString()
	if len(str) == 0 {
		if iter.head < len(iter.buf) && iter.buf[iter.head] == 'n' {
			return iter.skipFourBytes('n', 'u', 'l', 'l') // null
		}
		err := iter.ReportError("ReadFloat64", "number not found")
		iter.Skip()
		return err
	}
	val, err := strconv.ParseFloat(string(str), 64)
	if err != nil {
		iter.reportError(err)
		return err
	}
	*out = val
	return nil
}

// ReadUint read uint
func (iter *Iterator) ReadUint(out *uint) (err error) {
	if strconv.IntSize == 32 {
		err = iter.ReadUint32((*uint32)(unsafe.Pointer(out)))
	} else {
		err = iter.ReadUint64((*uint64)(unsafe.Pointer(out)))
	}
	iter.assertInteger(&err)
	return
}

// ReadInt read int
func (iter *Iterator) ReadInt(out *int) (err error) {
	if strconv.IntSize == 32 {
		err = iter.ReadInt32((*int32)(unsafe.Pointer(out)))
	} else {
		err = iter.ReadInt64((*int64)(unsafe.Pointer(out)))
	}
	iter.assertInteger(&err)
	return
}

func (iter *Iterator) assertInteger(err *error) {
	if iter.head < len(iter.buf) && iter.buf[iter.head] == '.' && *err == nil {
		*err = iter.ReportError("assertInteger", "found float instead of integer")
	}
}

// ReadInt8 read int8
func (iter *Iterator) ReadInt8(out *int8) (ret error) {
	c := iter.nextToken()
	if c == '-' {
		var val uint32
		err := iter.readUint32(&val, iter.readByte())
		if err != nil {
			return err
		}
		if val > math.MaxInt8+1 {
			return iter.ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(val), 10))
		}
		*out = -int8(val)
		iter.assertInteger(&ret)
		return
	}
	switch c {
	case '"':
		ret = iter.ReportError("ReadInt8", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadInt8", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadInt8", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadInt8", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadInt8", "unexpected object")
		iter.skipObject()
		return
	}
	var val uint32
	err := iter.readUint32(&val, c)
	if err != nil {
		return err
	}
	if val > math.MaxInt8 {
		return iter.ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(val), 10))
	}
	*out = int8(val)
	iter.assertInteger(&ret)
	return
}

// ReadUint8 read uint8
func (iter *Iterator) ReadUint8(out *uint8) (ret error) {
	var val uint32
	c := iter.nextToken()
	switch c {
	case '"':
		ret = iter.ReportError("ReadUint8", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadUint8", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadUint8", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadUint8", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadUint8", "unexpected object")
		iter.skipObject()
		return
	}
	err := iter.readUint32(&val, c)
	if err != nil {
		return err
	}
	if val > math.MaxUint8 {
		return iter.ReportError("ReadUint8", "overflow: "+strconv.FormatInt(int64(val), 10))
	}
	*out = uint8(val)
	iter.assertInteger(&ret)
	return
}

// ReadInt16 read int16
func (iter *Iterator) ReadInt16(out *int16) (ret error) {
	c := iter.nextToken()
	if c == '-' {
		var val uint32
		err := iter.readUint32(&val, iter.readByte())
		if err != nil {
			return err
		}
		if val > math.MaxInt16+1 {
			return iter.ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(val), 10))
		}
		*out = -int16(val)
		iter.assertInteger(&ret)
		return
	}
	switch c {
	case '"':
		ret = iter.ReportError("ReadInt16", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadInt16", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadInt16", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadInt16", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadInt16", "unexpected object")
		iter.skipObject()
		return
	}
	var val uint32
	err := iter.readUint32(&val, c)
	if err != nil {
		return err
	}
	if val > math.MaxInt16 {
		return iter.ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(val), 10))
	}
	*out = int16(val)
	iter.assertInteger(&ret)
	return
}

// ReadUint16 read uint16
func (iter *Iterator) ReadUint16(out *uint16) (ret error) {
	var val uint32
	c := iter.nextToken()
	switch c {
	case '"':
		ret = iter.ReportError("ReadUint16", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadUint16", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadUint16", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadUint16", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadUint16", "unexpected object")
		iter.skipObject()
		return
	}
	err := iter.readUint32(&val, c)
	if err != nil {
		return err
	}
	if val > math.MaxUint16 {
		return iter.ReportError("ReadUint16", "overflow: "+strconv.FormatInt(int64(val), 10))
	}
	*out = uint16(val)
	iter.assertInteger(&ret)
	return
}

// ReadInt32 read int32
func (iter *Iterator) ReadInt32(out *int32) (ret error) {
	c := iter.nextToken()
	if c == '-' {
		var val uint32
		err := iter.readUint32(&val, iter.readByte())
		if err != nil {
			return err
		}
		if val > math.MaxInt32+1 {
			return iter.ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(val), 10))
		}
		*out = -int32(val)
		iter.assertInteger(&ret)
		return
	}
	switch c {
	case '"':
		ret = iter.ReportError("ReadInt32", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadInt32", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadInt32", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadInt32", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadInt32", "unexpected object")
		iter.skipObject()
		return
	}
	var val uint32
	err := iter.readUint32(&val, c)
	if err != nil {
		return err
	}
	if val > math.MaxInt32 {
		return iter.ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(val), 10))
	}
	*out = int32(val)
	iter.assertInteger(&ret)
	return
}

// ReadUint32 read uint32
func (iter *Iterator) ReadUint32(out *uint32) (ret error) {
	c := iter.nextToken()
	switch c {
	case '"':
		ret = iter.ReportError("ReadUint32", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadUint32", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadUint32", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadUint32", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadUint32", "unexpected object")
		iter.skipObject()
		return
	}
	ret = iter.readUint32(out, c)
	iter.assertInteger(&ret)
	return
}

func (iter *Iterator) readUint32(out *uint32, c byte) error {
	ind := intDigits[c]
	if ind == 0 {
		*out = 0
		return nil // single zero
	}
	if ind == invalidCharForNumber {
		return iter.ReportError("readUint32", "unexpected character: "+string([]byte{byte(ind)}))
	}
	value := uint32(ind)
	if len(iter.buf)-iter.head > 10 {
		i := iter.head
		ind2 := intDigits[iter.buf[i]]
		if ind2 == invalidCharForNumber {
			iter.head = i
			*out = value
			return nil
		}
		i++
		ind3 := intDigits[iter.buf[i]]
		if ind3 == invalidCharForNumber {
			iter.head = i
			*out = value*10 + uint32(ind2)
			return nil
		}
		//iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		i++
		ind4 := intDigits[iter.buf[i]]
		if ind4 == invalidCharForNumber {
			iter.head = i
			*out = value*100 + uint32(ind2)*10 + uint32(ind3)
			return nil
		}
		i++
		ind5 := intDigits[iter.buf[i]]
		if ind5 == invalidCharForNumber {
			iter.head = i
			*out = value*1000 + uint32(ind2)*100 + uint32(ind3)*10 + uint32(ind4)
			return nil
		}
		i++
		ind6 := intDigits[iter.buf[i]]
		if ind6 == invalidCharForNumber {
			iter.head = i
			*out = value*10000 + uint32(ind2)*1000 + uint32(ind3)*100 + uint32(ind4)*10 + uint32(ind5)
			return nil
		}
		i++
		ind7 := intDigits[iter.buf[i]]
		if ind7 == invalidCharForNumber {
			iter.head = i
			*out = value*100000 + uint32(ind2)*10000 + uint32(ind3)*1000 + uint32(ind4)*100 + uint32(ind5)*10 + uint32(ind6)
			return nil
		}
		i++
		ind8 := intDigits[iter.buf[i]]
		if ind8 == invalidCharForNumber {
			iter.head = i
			*out = value*1000000 + uint32(ind2)*100000 + uint32(ind3)*10000 + uint32(ind4)*1000 + uint32(ind5)*100 + uint32(ind6)*10 + uint32(ind7)
			return nil
		}
		i++
		ind9 := intDigits[iter.buf[i]]
		value = value*10000000 + uint32(ind2)*1000000 + uint32(ind3)*100000 + uint32(ind4)*10000 + uint32(ind5)*1000 + uint32(ind6)*100 + uint32(ind7)*10 + uint32(ind8)
		iter.head = i
		if ind9 == invalidCharForNumber {
			*out = value
			return nil
		}
	}
	for i := iter.head; i < len(iter.buf); i++ {
		ind = intDigits[iter.buf[i]]
		if ind == invalidCharForNumber {
			iter.head = i
			*out = value
			return nil
		}
		if value > uint32SafeToMultiply10 {
			value2 := (value << 3) + (value << 1) + uint32(ind)
			if value2 < value {
				return iter.ReportError("readUint32", "overflow")
			}
			value = value2
			continue
		}
		value = (value << 3) + (value << 1) + uint32(ind)
	}
	*out = value
	return nil
}

// ReadInt64 read int64
func (iter *Iterator) ReadInt64(out *int64) (ret error) {
	c := iter.nextToken()
	if c == '-' {
		var val uint64
		err := iter.readUint64(&val, iter.readByte())
		if err != nil {
			return err
		}
		if val > math.MaxInt64+1 {
			return iter.ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(val), 10))
		}
		*out = -int64(val)
		iter.assertInteger(&ret)
		return
	}
	switch c {
	case '"':
		ret = iter.ReportError("ReadInt64", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadInt64", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadInt64", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadInt64", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadInt64", "unexpected object")
		iter.skipObject()
		return
	}
	var val uint64
	err := iter.readUint64(&val, c)
	if err != nil {
		return err
	}
	if val > math.MaxInt64 {
		return iter.ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(val), 10))
	}
	*out = int64(val)
	iter.assertInteger(&ret)
	return
}

// ReadUint64 read uint64
func (iter *Iterator) ReadUint64(out *uint64) (ret error) {
	c := iter.nextToken()
	switch c {
	case '"':
		ret = iter.ReportError("ReadUint64", "unexpected string")
		iter.skipString()
		return
	case 'n':
		return iter.skipThreeBytes('u', 'l', 'l') // null
	case 't':
		ret = iter.ReportError("ReadUint64", "unexpected true")
		iter.skipThreeBytes('r', 'u', 'e') // true
		return
	case 'f':
		ret = iter.ReportError("ReadUint64", "unexpected false")
		iter.skipFourBytes('a', 'l', 's', 'e') // false
		return
	case '[':
		ret = iter.ReportError("ReadUint64", "unexpected array")
		iter.skipArray()
		return
	case '{':
		ret = iter.ReportError("ReadUint64", "unexpected object")
		iter.skipObject()
		return
	}
	ret = iter.readUint64(out, c)
	iter.assertInteger(&ret)
	return
}

func (iter *Iterator) readUint64(out *uint64, c byte) error {
	ind := intDigits[c]
	if ind == 0 {
		*out = 0 // single zero
		return nil
	}
	if ind == invalidCharForNumber {
		return iter.ReportError("readUint64", "unexpected character: "+string([]byte{byte(ind)}))
	}
	value := uint64(ind)
	if len(iter.buf)-iter.head > 10 {
		i := iter.head
		ind2 := intDigits[iter.buf[i]]
		if ind2 == invalidCharForNumber {
			iter.head = i
			*out = value
			return nil
		}
		i++
		ind3 := intDigits[iter.buf[i]]
		if ind3 == invalidCharForNumber {
			iter.head = i
			*out = value*10 + uint64(ind2)
			return nil
		}
		i++
		ind4 := intDigits[iter.buf[i]]
		if ind4 == invalidCharForNumber {
			iter.head = i
			*out = value*100 + uint64(ind2)*10 + uint64(ind3)
			return nil
		}
		i++
		ind5 := intDigits[iter.buf[i]]
		if ind5 == invalidCharForNumber {
			iter.head = i
			*out = value*1000 + uint64(ind2)*100 + uint64(ind3)*10 + uint64(ind4)
			return nil
		}
		i++
		ind6 := intDigits[iter.buf[i]]
		if ind6 == invalidCharForNumber {
			iter.head = i
			*out = value*10000 + uint64(ind2)*1000 + uint64(ind3)*100 + uint64(ind4)*10 + uint64(ind5)
			return nil
		}
		i++
		ind7 := intDigits[iter.buf[i]]
		if ind7 == invalidCharForNumber {
			iter.head = i
			*out = value*100000 + uint64(ind2)*10000 + uint64(ind3)*1000 + uint64(ind4)*100 + uint64(ind5)*10 + uint64(ind6)
			return nil
		}
		i++
		ind8 := intDigits[iter.buf[i]]
		if ind8 == invalidCharForNumber {
			iter.head = i
			*out = value*1000000 + uint64(ind2)*100000 + uint64(ind3)*10000 + uint64(ind4)*1000 + uint64(ind5)*100 + uint64(ind6)*10 + uint64(ind7)
			return nil
		}
		i++
		ind9 := intDigits[iter.buf[i]]
		value = value*10000000 + uint64(ind2)*1000000 + uint64(ind3)*100000 + uint64(ind4)*10000 + uint64(ind5)*1000 + uint64(ind6)*100 + uint64(ind7)*10 + uint64(ind8)
		iter.head = i
		if ind9 == invalidCharForNumber {
			*out = value
			return nil
		}
	}
	for i := iter.head; i < len(iter.buf); i++ {
		ind = intDigits[iter.buf[i]]
		if ind == invalidCharForNumber {
			iter.head = i
			*out = value
			return nil
		}
		if value > uint64SafeToMultiple10 {
			value2 := (value << 3) + (value << 1) + uint64(ind)
			if value2 < value {
				return iter.ReportError("readUint64", "overflow")
			}
			value = value2
			continue
		}
		value = (value << 3) + (value << 1) + uint64(ind)
	}
	*out = value
	return nil
}
