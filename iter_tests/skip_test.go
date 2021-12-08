package iter_tests

import (
	"math/big"
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
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_read_raw_message(t *testing.T) {
	var msg jsoniter.RawMessage
	iter := jsoniter.ParseBytes([]byte(`"hello",100`))
	iter.ReadRawMessage(&msg)
	if string(msg) != `"hello"` {
		t.Fatal()
	}
}

func Test_skip_null(t *testing.T) {
	parseNull := func() *jsoniter.Iterator {
		return jsoniter.ParseBytes([]byte(`null`))
	}
	if parseNull().ReadFloat64(new(float64)) != nil {
		t.Fatal()
	}
	if parseNull().ReadFloat32(new(float32)) != nil {
		t.Fatal()
	}
	if parseNull().ReadBigInt(new(big.Int)) != nil {
		t.Fatal()
	}
	if parseNull().ReadBigFloat(new(big.Float)) != nil {
		t.Fatal()
	}
	if parseNull().ReadInt(new(int)) != nil {
		t.Fatal()
	}
	if parseNull().ReadInt64(new(int64)) != nil {
		t.Fatal()
	}
	if parseNull().ReadInt32(new(int32)) != nil {
		t.Fatal()
	}
	if parseNull().ReadUint(new(uint)) != nil {
		t.Fatal()
	}
	if parseNull().ReadUint64(new(uint64)) != nil {
		t.Fatal()
	}
	if parseNull().ReadUint32(new(uint32)) != nil {
		t.Fatal()
	}
	if parseNull().ReadUint16(new(uint16)) != nil {
		t.Fatal()
	}
	if parseNull().ReadInt16(new(int16)) != nil {
		t.Fatal()
	}
	if parseNull().ReadUint8(new(uint8)) != nil {
		t.Fatal()
	}
	if parseNull().ReadInt8(new(int8)) != nil {
		t.Fatal()
	}
	if parseNull().ReadString(new(string)) != nil {
		t.Fatal()
	}
	if parseNull().ReadBool(new(bool)) != nil {
		t.Fatal()
	}
	if parseNull().ReadNumber(new(jsoniter.Number)) != nil {
		t.Fatal()
	}
	iter := parseNull()
	if iter.ReadObjectHead() != false {
		t.Fatal()
	}
	if iter.Error != nil {
		t.Fatal()
	}
	iter = parseNull()
	if iter.ReadArrayHead() != false {
		t.Fatal()
	}
	if iter.Error != nil {
		t.Fatal()
	}
}

func Test_skip_invalid_int(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadInt(new(int))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_int64(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadInt64(new(int64))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_int32(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadInt32(new(int32))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_uint(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadUint(new(uint))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_uint64(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadUint64(new(uint64))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_uint32(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadUint32(new(uint32))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_uint16(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadUint16(new(uint16))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_int16(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadInt16(new(int16))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_uint8(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadUint8(new(uint8))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_int8(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadInt8(new(int8))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_float64(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadFloat64(new(float64))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_float32(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadFloat32(new(float32))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_big_int(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadBigInt(new(big.Int))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_big_float(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadBigFloat(new(big.Float))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_string(t *testing.T) {
	inputs := []string{
		`100`,
		`true`,
		`false`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadString(new(string))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_bool(t *testing.T) {
	inputs := []string{
		`100`,
		`"hello"`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadBool(new(bool))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}

func Test_skip_invalid_Number(t *testing.T) {
	inputs := []string{
		`"hello"`,
		`[]`,
		`{}`,
	}
	for _, input := range inputs {
		iter := jsoniter.ParseBytes([]byte(input + `,100`))
		iter.ReadNumber(new(jsoniter.Number))
		if iter.Error == nil {
			t.Fatal()
		}
		iter.ReadArrayMore()
		var val int
		iter.ReadInt(&val)
		if val != 100 {
			t.Fatal(input)
		}
	}
}
