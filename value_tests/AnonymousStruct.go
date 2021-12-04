package value_tests

//go:generate go run github.com/json-iterator/tinygo/gen
type AnonymousStruct struct {
	Value struct {
		Name  string
		Price int
	}
}
