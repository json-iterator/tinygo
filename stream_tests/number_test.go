package stream_tests

import (
	"encoding/json"
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

func Test_write_uint32(t *testing.T) {
	vals := []uint32{0, 1, 11, 111, 255, 999999, 0xfff, 0xffff, 0xfffff, 0xffffff, 0xfffffff, 0xffffffff}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteUint32(val)
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

func Test_write_int32(t *testing.T) {
	vals := []int32{0, 1, 11, 111, 255, 999999, 0xfff, 0xffff, 0xfffff, 0xffffff, 0xfffffff, 0x7fffffff, -0x80000000}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInt32(val)
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

func Test_write_uint64(t *testing.T) {
	vals := []uint64{0, 1, 11, 111, 255, 999999, 0xfff, 0xffff, 0xfffff, 0xffffff, 0xfffffff, 0xffffffff,
		0xfffffffff, 0xffffffffff, 0xfffffffffff, 0xffffffffffff, 0xfffffffffffff, 0xffffffffffffff,
		0xfffffffffffffff, 0xffffffffffffffff}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteUint64(val)
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

func Test_write_int64(t *testing.T) {
	vals := []int64{0, 1, 11, 111, 255, 999999, 0xfff, 0xffff, 0xfffff, 0xffffff, 0xfffffff, 0xffffffff,
		0xfffffffff, 0xffffffffff, 0xfffffffffff, 0xffffffffffff, 0xfffffffffffff, 0xffffffffffffff,
		0xfffffffffffffff, 0x7fffffffffffffff, -0x8000000000000000}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInt64(val)
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

func Test_write_float32(t *testing.T) {
	vals := []float32{0, 1, -1, 99, 0xff, 0xfff, 0xffff, 0xfffff, 0xffffff, 0x4ffffff, 0xfffffff,
		-0x4ffffff, -0xfffffff, 1.2345, 1.23456, 1.234567, 1.001}
	for _, val := range vals {
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteFloat32(val)
			output, err := json.Marshal(val)
			if err != nil {
				t.Fatal()
			}
			if string(output) != string(stream.Buffer()) {
				t.Fatal()
			}
		})
		t.Run(fmt.Sprintf("%v", val), func(t *testing.T) {
			stream := jsoniter.NewStream()
			stream.WriteInterface(val)
			output, err := json.Marshal(val)
			if err != nil {
				t.Fatal()
			}
			if string(output) != string(stream.Buffer()) {
				t.Fatal()
			}
		})
	}
}
