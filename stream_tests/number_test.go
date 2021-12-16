package stream_tests

import (
	"fmt"
	"strconv"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_write_uint8(t *testing.T) {
	vals := []uint8{0, 1, 11, 111, 255}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteUint8(val)
			if strconv.FormatUint(uint64(val), 10) != string(stream.Buffer()) {
				t.Fatal()
			}
		})
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInterface(val)
			if strconv.FormatUint(uint64(val), 10) != string(stream.Buffer()) {
				t.Fatal()
			}
		})
	}
}

func Test_write_int8(t *testing.T) {
	vals := []int8{0, 1, -1, 99, 0x7f, -0x80}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInt8(val)
			if strconv.FormatInt(int64(val), 10) != string(stream.Buffer()) {
				t.Fatal(string(stream.Buffer()))
			}
		})
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInterface(val)
			if strconv.FormatInt(int64(val), 10) != string(stream.Buffer()) {
				t.Fatal(string(stream.Buffer()))
			}
		})
	}
}

func Test_write_uint16(t *testing.T) {
	vals := []uint16{0, 1, 11, 111, 255, 0xfff, 0xffff}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteUint16(val)
			if strconv.FormatUint(uint64(val), 10) != string(stream.Buffer()) {
				t.Fatal()
			}
		})
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInterface(val)
			if strconv.FormatUint(uint64(val), 10) != string(stream.Buffer()) {
				t.Fatal()
			}
		})
	}
}

func Test_write_int16(t *testing.T) {
	vals := []int16{0, 1, 11, 111, 255, 0xfff, 0x7fff, -0x8000}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInt16(val)
			if strconv.FormatInt(int64(val), 10) != string(stream.Buffer()) {
				t.Fatal(string(stream.Buffer()))
			}
		})
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInterface(val)
			if strconv.FormatInt(int64(val), 10) != string(stream.Buffer()) {
				t.Fatal(string(stream.Buffer()))
			}
		})
	}
}
