package stream_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_write_string(t *testing.T) {
	stream := jsoniter.NewStream()
	stream.WriteString("hello")
	if string(stream.Buffer()) != `"hello"` {
		t.Fatal()
	}
	stream = jsoniter.NewStream()
	stream.WriteInterface("hello")
	if string(stream.Buffer()) != `"hello"` {
		t.Fatal()
	}
}
