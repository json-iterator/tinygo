package iter_tests

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_read_float(t *testing.T) {
	inputs := []string{
		`1.1`, `1000`, `9223372036854775807`, `12.3`, `-12.3`, `720368.54775807`, `720368.547758075`,
		`1e1`, `1e+1`, `1e-1`, `1E1`, `1E+1`, `1E-1`, `-1e1`, `-1e+1`, `-1e-1`,
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseFloat(input, 32)
			if err != nil {
				t.Fatal(err)
			}
			var val float32
			iter.ReadFloat32(&val)
			if float32(expected) != val {
				t.Fatal()
			}
		})
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseFloat(input, 64)
			if err != nil {
				t.Fatal(err)
			}
			var val float64
			iter.ReadFloat64(&val)
			if expected != val {
				t.Fatal()
			}
		})
	}
}

func Test_int8(t *testing.T) {
	inputs := []string{`127`, `-128`}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseInt(input, 10, 8)
			if err != nil {
				panic(err)
			}
			var val int8
			iter.ReadInt8(&val)
			if int8(expected) != val {
				t.Fail()
			}
		})
	}
}

func Test_read_int16(t *testing.T) {
	inputs := []string{`32767`, `-32768`}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseInt(input, 10, 16)
			if err != nil {
				panic(err)
			}
			var val int16
			iter.ReadInt16(&val)
			if int16(expected) != val {
				t.Fail()
			}
		})
	}
}

func Test_read_int32(t *testing.T) {
	inputs := []string{`1`, `12`, `123`, `1234`, `12345`, `123456`, `2147483647`, `-2147483648`}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				panic(err)
			}
			var val int32
			iter.ReadInt32(&val)
			if int32(expected) != val {
				t.Fail()
			}
		})
	}
}

func Test_read_int_overflow(t *testing.T) {
	inputArr := []string{"123451", "-123451"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt8(new(int8))
		if iter.Error == nil {
			t.Fail()
		}
	}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint8(new(uint8))
		if iter.Error == nil {
			t.Fail()
		}
	}
	inputArr = []string{"12345678912", "-12345678912"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt16(new(int16))
		if iter.Error == nil {
			t.Fail()
		}
	}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint16(new(uint16))
		if iter.Error == nil {
			t.Fail()
		}
	}
	inputArr = []string{"3111111111", "-3111111111", "1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt32(new(int32))
		if iter.Error == nil {
			t.Fatal("ReadInt32", input)
		}
	}
	inputArr = []string{"1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint32(new(uint32))
		if iter.Error == nil {
			t.Fatal("ReadUint32", input)
		}
	}
	inputArr = []string{"9223372036854775811", "-9523372036854775807", "1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt64(new(int64))
		if iter.Error == nil {
			t.Fatal("ReadInt64", input)
		}
	}
	inputArr = []string{"1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint64(new(uint64))
		if iter.Error == nil {
			t.Fatal("ReadUint64", input)
		}
	}
}

func Test_read_int64(t *testing.T) {
	inputs := []string{`1`, `12`, `123`, `1234`, `12345`, `123456`, `9223372036854775807`, `-9223372036854775808`}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			var val int64
			iter.ReadInt64(&val)
			if expected != val {
				t.Fatal()
			}
		})
	}
}

func Test_read_float_as_integer(t *testing.T) {
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadInt(new(int)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadUint(new(uint)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadInt64(new(int64)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadUint64(new(uint64)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadInt32(new(int32)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadUint32(new(uint32)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadInt16(new(int16)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadUint16(new(uint16)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadInt8(new(int8)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadUint8(new(uint8)) == nil {
		t.Fatal()
	}
	if jsoniter.ParseBytes([]byte(`1.1`)).ReadBigInt(new(big.Int)) == nil {
		t.Fatal()
	}
}
