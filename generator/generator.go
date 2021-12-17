package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

var fset *token.FileSet
var lines = []byte{}
var indent = 0
var anonymousCounter = 1
var referencedImports = map[string]string{}
var allImports = map[string]string{}
var prefix string

func Generate(_fset *token.FileSet, _allImports map[string]string, typeSpec *ast.TypeSpec) []byte {
	fset = _fset
	allImports = _allImports
	typeName := typeSpec.Name.Name
	decoders := generateDecoders(typeSpec)
	encoders := generateEncoders(typeSpec)
	lines = []byte{}
	_l(fmt.Sprintf("package %s", os.Getenv("GOPACKAGE")))
	_n()
	_l(`import jsoniter "github.com/json-iterator/tinygo"`)
	for alias, path := range referencedImports {
		_f(`import %s "%s"`, alias, path)
	}
	_n()
	_f("type %s_json struct {", typeName)
	_l("}")
	_f("func (json %s_json) Type() interface{} {", typeName)
	_f("  var val %s", typeName)
	_l("  return val")
	_l("}")
	_f("func (json %s_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {", typeName)
	_f("  %s_json_unmarshal(iter, out.(*%s))", prefix, typeName)
	_l("}")
	_f("func (json %s_json) Marshal(stream *jsoniter.Stream, val interface{}) {", typeName)
	_f("  %s_json_marshal(stream, val.(%s))", prefix, typeName)
	_l("}")
	lines = append(lines, decoders...)
	lines = append(lines, encoders...)
	return lines
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
}

func _n() {
	_l("")
}

func _f(format string, a ...interface{}) {
	_l(fmt.Sprintf(format, a...))
}

func nodeToString(node ast.Node) string {
	if x, ok := node.(*ast.SelectorExpr); ok {
		alias, ok := x.X.(*ast.Ident)
		if ok {
			if path, ok := allImports[alias.Name]; ok {
				referencedImports[alias.Name] = path
			}
		}
	}
	buf := bytes.NewBuffer([]byte{})
	err := printer.Fprint(buf, fset, node)
	if err != nil {
		reportError(err)
		return ""
	}
	return buf.String()
}
