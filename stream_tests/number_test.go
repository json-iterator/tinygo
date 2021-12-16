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
