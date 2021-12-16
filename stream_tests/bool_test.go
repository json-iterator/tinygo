package stream_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_write_true(t *testing.T) {
	stream := jsoniter.NewStream()
	stream.WriteBool(true)
	if string(stream.Buffer()) != "true" {
		t.Fatal()
	}
	stream = jsoniter.NewStream()
	stream.WriteInterface(true)
	if string(stream.Buffer()) != "true" {
		t.Fatal()
	}
}

func Test_write_false(t *testing.T) {
	stream := jsoniter.NewStream()
	stream.WriteBool(false)
	if string(stream.Buffer()) != "false" {
		t.Fatal()
	}
	stream = jsoniter.NewStream()
	stream.WriteInterface(false)
	if string(stream.Buffer()) != "false" {
		t.Fatal()
	}
}
