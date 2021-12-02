package main

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_bad_strings(t *testing.T) {
	badInputs := []string{
		``,
		`"`,
		`"\"`,
		`"\\\"`,
		"\"\n\"",
		`"\U0001f64f"`,
		`"\uD83D\u00"`,
	}
	for i := 0; i < 32; i++ {
		// control characters are invalid
		badInputs = append(badInputs, string([]byte{'"', byte(i), '"'}))
	}
	for _, input := range badInputs {
		iter := jsoniter.ParseBytes([]byte(input))
		iter.ReadString()
		if iter.Error == nil {
			panic("expect ReadString reports error")
		}
	}
}

func Test_good_strings(t *testing.T) {
	goodInputs := []struct {
		input       string
		expectValue string
	}{
		{`""`, ""},
		{`"a"`, "a"},
		{`"Iñtërnâtiônàlizætiøn,💝🐹🌇⛔"`, "Iñtërnâtiônàlizætiøn,💝🐹🌇⛔"},
		{`"\uD83D"`, string([]byte{239, 191, 189})},
		{`"\uD83D\\"`, string([]byte{239, 191, 189, '\\'})},
		{`"\uD83D\ub000"`, string([]byte{239, 191, 189, 235, 128, 128})},
		{`"\uD83D\ude04"`, "😄"},
		{`"\uDEADBEEF"`, string([]byte{239, 191, 189, 66, 69, 69, 70})},
		{`"hel\"lo"`, `hel"lo`},
		{`"hel\\\/lo"`, `hel\/lo`},
		{`"hel\\blo"`, `hel\blo`},
		{`"hel\\\blo"`, "hel\\\blo"},
		{`"hel\\nlo"`, `hel\nlo`},
		{`"hel\\\nlo"`, "hel\\\nlo"},
		{`"hel\\tlo"`, `hel\tlo`},
		{`"hel\\flo"`, `hel\flo`},
		{`"hel\\\flo"`, "hel\\\flo"},
		{`"hel\\\rlo"`, "hel\\\rlo"},
		{`"hel\\\tlo"`, "hel\\\tlo"},
		{`"\u4e2d\u6587"`, "中文"},
		{`"\ud83d\udc4a"`, "\xf0\x9f\x91\x8a"},
	}

	for _, tc := range goodInputs {
		iter := jsoniter.ParseBytes([]byte(tc.input))
		actual := iter.ReadString()
		if iter.Error != nil {
			panic(fmt.Sprintf("expect ReadString not to reports error: %s", tc.input))
		}
		if actual != tc.expectValue {
			panic(fmt.Sprintf("expected %s, actual %s", tc.expectValue, actual))
		}
	}
}
