package value_tests

//go:generate go run github.com/json-iterator/tinygo/gen
type WithEmbedStructBase1 struct {
	Field1 string
}

//go:generate go run github.com/json-iterator/tinygo/gen
type WithEmbedStructBase2 struct {
	Field2 string
}

//go:generate go run github.com/json-iterator/tinygo/gen
type WithEmbedStruct struct {
	*WithEmbedStructBase1
	WithEmbedStructBase2
	Field3 string
}
