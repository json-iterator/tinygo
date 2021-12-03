package iter_tests

import (
	"fmt"
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
			if float32(expected) != iter.ReadFloat32() {
				t.Fatal()
			}
		})
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			iter := jsoniter.ParseBytes([]byte(input))
			expected, err := strconv.ParseFloat(input, 64)
			if err != nil {
				t.Fatal(err)
			}
			if expected != iter.ReadFloat64() {
				t.Fatal()
			}
		})
	}
}
