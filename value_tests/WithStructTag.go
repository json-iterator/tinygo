package value_tests

//go:generate go run github.com/json-iterator/tinygo/gen
type WithStructTag struct {
	Field1 string `json:"field1"`
	Field2 string `json:"-"`
}
