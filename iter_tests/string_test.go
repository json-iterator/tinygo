package main

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_valid_string(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`"hello"`))
	if iter.WhatIsNext() != jsoniter.StringValue {
		t.Fail()
	}
	val := iter.ReadString()
	if val != "hello" {
		t.Fail()
	}
}

func Test_string_missing_starting_quote(t *testing.T) {
	iter := jsoniter.ParseBytes([]byte(`1,2`))
	iter.ReadString()
	if iter.Error == nil {
		t.Fail()
	}
}
