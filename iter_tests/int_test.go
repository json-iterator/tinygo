package main

import (
	"fmt"
	"strconv"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_int8(t *testing.T) {
	inputs := []string{`127`, `-128`}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseInt(input, 10, 8)
			if err != nil {
				panic(err)
			}
			if int8(expected) != iter.ReadInt8() {
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
			if int16(expected) != iter.ReadInt16() {
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
			if int32(expected) != iter.ReadInt32() {
				t.Fail()
			}
		})
	}
}

func Test_read_int_overflow(t *testing.T) {
	inputArr := []string{"123451", "-123451"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt8()
		if iter.Error == nil {
			t.Fail()
		}
	}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint8()
		if iter.Error == nil {
			t.Fail()
		}
	}
	inputArr = []string{"12345678912", "-12345678912"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt16()
		if iter.Error == nil {
			t.Fail()
		}
	}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint16()
		if iter.Error == nil {
			t.Fail()
		}
	}
	inputArr = []string{"3111111111", "-3111111111", "1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt32()
		if iter.Error == nil {
			t.Fatal("ReadInt32", input)
		}
	}
	inputArr = []string{"1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint32()
		if iter.Error == nil {
			t.Fatal("ReadUint32", input)
		}
	}
	inputArr = []string{"9223372036854775811", "-9523372036854775807", "1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadInt64()
		if iter.Error == nil {
			t.Fatal("ReadInt64", input)
		}
	}
	inputArr = []string{"1234232323232323235678912", "-1234567892323232323212"}
	for _, input := range inputArr {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadUint64()
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
			if expected != iter.ReadInt64() {
				t.Fatal()
			}
		})
	}
}
