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
	appendLine(fmt.Sprintf("package %s", os.Getenv("GOPACKAGE")))
	appendNewLine()
	appendLine(`import jsoniter "github.com/json-iterator/tinygo"`)
	appendNewLine()
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

func appendLine(line string) {
	for i := 0; i < indent; i++ {
		lines = append(lines, ' ')
		print(" ")
	}
	lines = append(lines, line...)
	lines = append(lines, '\n')
	println(line)
}

func appendNewLine() {
	appendLine("")
}

func enter(line string) {
	appendLine(line)
	indent += 4
}

func exit(line string) {
	indent -= 4
	appendLine(line)
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
	enter(fmt.Sprintf("func jd_%s(iter *jsoniter.Iterator) %s {", escapeTypeName(typeName), typeName))
	appendLine(fmt.Sprintf("var val = %s{}", typeName))
	appendLine("if iter.Error != nil { return val }")
	enter("for iter.ReadArray() {")
	switch x := arrayType.Elt.(type) {
	case *ast.Ident:
		if x.Name == "string" {
			appendLine("val = append(val, iter.ReadString())")
		} else {
			reportError(fmt.Errorf("unknown element type of Array: %s", x.Name))
			return
		}
	default:
		reportError(fmt.Errorf("unknown element type of Array"))
		return
	}
	exit("}")
	appendLine("return val")
	exit("}")
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
