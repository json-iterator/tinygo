package iter_tests

import (
	"testing"

	jsoniter "github.com/json-iterator/tinygo"
)

func Test_skip(t *testing.T) {
	inputs := []string{
		"-0.12",
		`"hello"`,
		`"\t"`,
		`"\""`,
		`"\\\""`,
		"null",
		"true",
		"false",
		"[1, [2, [3], 4]]",
		"[ ]",
		`{"a" : [{"stream": "c"}], "d": 102 }`,
		`{abc}`, // we do not invalidate the skipped value
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.Skip()
		iter.ReadArrayMore()
		if iter.ReadInt() != 100 {
			t.Fatal(input)
		}
	}
}
