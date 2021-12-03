package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var fset = token.NewFileSet()
var lines = []byte{}
var cwd string
var indent = 0

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		reportError(err)
		return
	}
}

func main() {
	_l(fmt.Sprintf("package %s", os.Getenv("GOPACKAGE")))
	_n()
	_l(`import jsoniter "github.com/json-iterator/tinygo"`)
	_n()
	typeSpec := locateTypeSpec()
	typeName := typeSpec.Name.Name
	_f("func jd_%s(iter *jsoniter.Iterator, out *%s) {", escapeTypeName(typeName), typeName)
	switch x := typeSpec.Type.(type) {
	case *ast.ArrayType:
		genArray(typeName, x)
	case *ast.StructType:
		genStruct(x)
	default:
		reportError(fmt.Errorf("unknown type of TypeSpec"))
		return
	}
	_l("}")
	outputPath := filepath.Join(cwd, fmt.Sprintf("%s_json.go", typeSpec.Name.Name))
	os.WriteFile(outputPath, lines, 0644)
}

func reportError(err error) {
	panic(err)
}

func _l(line string) {
	for i := 0; i < indent; i++ {
		lines = append(lines, ' ')
		print(" ")
	}
	lines = append(lines, line...)
	lines = append(lines, '\n')
	println(line)
}

func _n() {
	_l("")
}

func _f(format string, a ...interface{}) {
	_l(fmt.Sprintf(format, a...))
}

func nodeToString(node ast.Node) string {
	buf := bytes.NewBuffer([]byte{})
	err := printer.Fprint(buf, fset, node)
	if err != nil {
		reportError(err)
		return ""
	}
	return buf.String()
}

func escapeTypeName(typeName string) string {
	return strings.ReplaceAll(typeName, "[]", "array_")
}

func genStruct(structType *ast.StructType) {
	_l("  if iter.Error != nil { return }")
	_l("  more := iter.ReadObjectHead()")
	_l("  for more {")
	_l("    field := iter.ReadObjectField()")
	_l("    switch {")
	for _, field := range structType.Fields.List {
		fieldName := field.Names[0].Name
		ptr := fmt.Sprintf("&(*out).%s", fieldName)
		_f("    case field == `%s`:", fieldName)
		switch x := field.Type.(type) {
		case *ast.Ident:
			_f("      "+getTypeDecoder(x.Name), ptr)
		default:
			reportError(fmt.Errorf("unknown field type of struct"))
			return
		}
	}
	_l("    default:")
	_l("      iter.Skip()")
	_l("    }")
	_l("    more = iter.ReadObjectMore()")
	_l("  }")
}

func genArray(typeName string, arrayType *ast.ArrayType) {
	_l("  if iter.Error != nil { return }")
	_l("  i := 0")
	_l("  val := *out")
	_l("  more := iter.ReadArrayHead()")
	_l("  for more {")
	_l(`    if i == len(val) {`)
	_f(`      val = append(val, make(%s, 4)...)`, typeName)
	_l(`    }`)
	switch x := arrayType.Elt.(type) {
	case *ast.Ident:
		_f("    "+getTypeDecoder(x.Name), "&val[i]")
	default:
		reportError(fmt.Errorf("unknown element type of array"))
		return
	}
	_l(`    i++`)
	_l(`    more = iter.ReadArrayMore()`)
	_l("  }")
	_l("  if i == 0 {")
	_f("    *out = %s{}", typeName)
	_l("  } else {")
	_l("    *out = val[:i]")
	_l("  }")
}

func getTypeDecoder(typeName string) string {
	switch {
	case typeName == "string":
		return "iter.ReadString(%s)"
	case typeName == "int":
		return "iter.ReadInt(%s)"
	default:
		reportError(fmt.Errorf("unknown type: %s", typeName))
		return ""
	}
}

func locateTypeSpec() *ast.TypeSpec {
	goline, err := strconv.Atoi(os.Getenv("GOLINE"))
	if err != nil {
		reportError(err)
		return nil
	}
	f, err := parser.ParseFile(fset, os.Getenv("GOFILE"), nil, parser.ParseComments)
	if err != nil {
		reportError(err)
		return nil
	}
	var located *ast.TypeSpec
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if goline+1 == fset.Position(x.Pos()).Line {
				located = x
			}
			return false
		}
		return true
	})
	if located == nil {
		reportError(fmt.Errorf("%s:%s go generate should be marked just before type definition",
			os.Getenv("GOFILE"), os.Getenv("GOLINE")))
		return nil
	}
	return located
}
