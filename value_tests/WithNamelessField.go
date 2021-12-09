package value_tests

//go:generate go run github.com/json-iterator/tinygo/gen
type WithNamelessField_f1 float64

//go:generate go run github.com/json-iterator/tinygo/gen
type WithNamelessField_f2 bool

//go:generate go run github.com/json-iterator/tinygo/gen
type WithNamelessField struct {
	WithNamelessField_f1
	*WithNamelessField_f2
	string // lowercase will not be exported
	int    // lowercase will not be exported
}
