package jsoniter

import (
	"math/big"
	"strconv"
)

func (iter *Iterator) readNumberAsString() (ret string) {
	for i := iter.head; i < len(iter.buf); i++ {
		switch iter.buf[i] {
		case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			str := string(iter.buf[iter.head:i])
			iter.head = i
			return str
		}
	}
	str := string(iter.buf[iter.head:])
	iter.head = len(iter.buf)
	return str
}

// ReadBigFloat read big.Float
func (iter *Iterator) ReadBigFloat() (ret *big.Float) {
	str := iter.readNumberAsString()
	prec := 64
	if len(str) > prec {
		prec = len(str)
	}
	val, _, err := big.ParseFloat(str, 10, uint(prec), big.ToZero)
	if err != nil {
		iter.reportError(err)
		return nil
	}
	return val
}

// ReadBigInt read big.Int
func (iter *Iterator) ReadBigInt() (ret *big.Int) {
	str := iter.readNumberAsString()
	ret = big.NewInt(0)
	var success bool
	ret, success = ret.SetString(str, 10)
	if !success {
		iter.ReportError("ReadBigInt", "invalid big int")
		return nil
	}
	return ret
}

//ReadFloat32 read float32
func (iter *Iterator) ReadFloat32() (ret float32) {
	str := iter.readNumberAsString()
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		iter.reportError(err)
		return
	}
	return float32(val)
}

// ReadFloat64 read float64
func (iter *Iterator) ReadFloat64() (ret float64) {
	str := iter.readNumberAsString()
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		iter.reportError(err)
		return
	}
	return val
}
