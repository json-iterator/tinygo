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
	switch x := typeSpec.Type.(type) {
	case *ast.ArrayType:
		genArray(x)
	default:
		reportError(fmt.Errorf("unknown type of TypeSpec"))
		return
	}
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

func genArray(arrayType *ast.ArrayType) {
	typeName := nodeToString(arrayType)
	_f("func jd_%s(iter *jsoniter.Iterator, out *%s) {", escapeTypeName(typeName), typeName)
	_l("  if iter.Error != nil { return }")
	_l("  i := 0")
	_l("  val := *out")
	_l("  for iter.ReadArray() {")
	switch x := arrayType.Elt.(type) {
	case *ast.Ident:
		if x.Name == "string" {
			_l("    elem := iter.ReadString()")
		} else {
			reportError(fmt.Errorf("unknown element type of Array: %s", x.Name))
			return
		}
	default:
		reportError(fmt.Errorf("unknown element type of Array"))
		return
	}
	_l(`    if i < len(val) {`)
	_l(`      if iter.Error == nil {`)
	_l(`        val[i] = elem`)
	_l(`      }`)
	_l(`    } else {`)
	_l(`      val = append(val, elem)`)
	_l(`    }`)
	_l(`    i++`)
	_l("  }")
	_l("  if i == 0 {")
	_f("    *out = %s{}", typeName)
	_l("  } else {")
	_l("    *out = val[:i]")
	_l("  }")
	_l("}")
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
