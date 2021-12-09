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
	"unicode"
)

var fset = token.NewFileSet()
var lines = []byte{}
var cwd string
var indent = 0
var anonymousCounter = 1
var anonymousDecoders = []byte{}
var referencedImports = map[string]string{}
var allImports = map[string]string{}
var prefix string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		reportError(err)
		return
	}
}

func main() {
	typeSpec := locateTypeSpec()
	typeName := typeSpec.Name.Name
	prefix = typeName
	switch x := typeSpec.Type.(type) {
	case *ast.ArrayType:
		genArray(typeName, x)
	case *ast.StructType:
		genStruct(typeName, x)
	case *ast.MapType:
		genMap(x)
	case *ast.StarExpr:
		genPtr(x)
	case *ast.Ident:
		genDecodeStmt(x, fmt.Sprintf("(*%s)(out)", x.Name))
	default:
		reportError(fmt.Errorf("not supported type: %s", nodeToString(typeSpec)))
		return
	}
	mainDecoder := lines
	lines = []byte{}
	switch x := typeSpec.Type.(type) {
	case *ast.StructType:
		genStructField(x)
	case *ast.StarExpr:
		genOtherField(typeName)
	case *ast.MapType:
		genOtherField(typeName)
	case *ast.ArrayType:
		genOtherField(typeName)
	case *ast.Ident:
		genOtherField(typeName)
	default:
		reportError(fmt.Errorf("not supported type: %s", nodeToString(typeSpec)))
		return
	}
	fieldsDecoder := lines
	lines = []byte{}
	_l(fmt.Sprintf("package %s", os.Getenv("GOPACKAGE")))
	_n()
	_l(`import jsoniter "github.com/json-iterator/tinygo"`)
	for alias, path := range referencedImports {
		_f(`import %s "%s"`, alias, path)
	}
	_n()
	_f("func %s_json_unmarshal(iter *jsoniter.Iterator, out *%s) {", prefix, typeName)
	lines = append(lines, mainDecoder...)
	_l("}")
	_f("func %s_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *%s) bool {", prefix, typeName)
	lines = append(lines, fieldsDecoder...)
	_l("  return false")
	_l("}")
	lines = append(lines, anonymousDecoders...)
	_f("type %s_json struct {", typeName)
	_l("}")
	_f("func (json %s_json) Type() interface{} {", typeName)
	_f("  var val %s", typeName)
	_l("  return &val")
	_l("}")
	_f("func (json %s_json) Unmarshal(iter *jsoniter.Iterator, val interface{}) {", typeName)
	_f("  %s_json_unmarshal(iter, val.(*%s))", prefix, typeName)
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

func genAnonymousArray(arrayType *ast.ArrayType) string {
	decoderName := fmt.Sprintf(`%s_array%d`, prefix, anonymousCounter)
	typeName := nodeToString(arrayType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, typeName)
	genArray(typeName, arrayType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func genAnonymousStruct(structType *ast.StructType) string {
	decoderName := fmt.Sprintf(`%s_struct%d`, prefix, anonymousCounter)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal_field (iter *jsoniter.Iterator, field string, out *%s) bool {", decoderName, nodeToString(structType))
	genStructField(structType)
	_l("  return false")
	_l("}")
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, nodeToString(structType))
	genStruct(decoderName, structType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func genAnonymousMap(mapType *ast.MapType) string {
	decoderName := fmt.Sprintf(`%s_map%d`, prefix, anonymousCounter)
	typeName := nodeToString(mapType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, typeName)
	genMap(mapType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func genAnonymousPtr(ptrType *ast.StarExpr) string {
	decoderName := fmt.Sprintf(`%s_ptr%d`, prefix, anonymousCounter)
	typeName := nodeToString(ptrType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, typeName)
	genPtr(ptrType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func genPtr(ptrType *ast.StarExpr) {
	_f("    var val %s", nodeToString(ptrType.X))
	genDecodeStmt(ptrType.X, "&val")
	_l("    if iter.Error == nil {")
	_l("      *out = &val")
	_l("    }")
}

func genMap(mapType *ast.MapType) {
	_l("  more := iter.ReadObjectHead()")
	_l("  if *out == nil && iter.Error == nil {")
	_f("    *out = make(%s)", nodeToString(mapType))
	_l("  }")
	_l("  for more {")
	_l("    field := iter.ReadObjectField()")
	_f("    var value %s", nodeToString(mapType.Value))
	_f("    var key %s", nodeToString(mapType.Key))
	_l("    var err error")
	switch x := mapType.Key.(type) {
	case *ast.Ident:
		keyTypeName := x.Name
		switch keyTypeName {
		case "string":
			_l("    key = field")
		case "int":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadInt(&key)")
		case "uint":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadUint(&key)")
		case "int64":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadInt64(&key)")
		case "uint64":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadUint64(&key)")
		case "int32":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadInt32(&key)")
		case "uint32":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadUint32(&key)")
		case "int16":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadInt16(&key)")
		case "uint16":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadUint16(&key)")
		case "int8":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadInt8(&key)")
		case "uint8":
			_l("    err = jsoniter.ParseBytes([]byte(field)).ReadUint8(&key)")
		default:
			reportError(fmt.Errorf("unsupported map key type: %s", nodeToString(mapType.Key)))
			return
		}
	default:
		reportError(fmt.Errorf("unsupported map key type: %s", nodeToString(mapType.Key)))
		return
	}
	genDecodeStmt(mapType.Value, "&value")
	_l("    if err != nil {")
	_l(`      iter.ReportError("read map key", err.Error())`)
	_l("    } else {")
	_l("      (*out)[key] = value")
	_l("    }")
	_l("    more = iter.ReadObjectMore()")
	_l("  }")
}

func genOtherField(expectedFieldName string) {
	_f(`  if field == "%s" {`, expectedFieldName)
	_f("    %s_json_unmarshal(iter, out)", prefix)
	_l("    return true")
	_l("  }")
}

func genStructField(structType *ast.StructType) {
	for i, field := range structType.Fields.List {
		if len(field.Names) != 0 {
			continue
		}
		switch x := field.Type.(type) {
		case *ast.StarExpr: // embed pointer
			switch y := x.X.(type) {
			case *ast.Ident:
				_f("  var val%d %s", i, y.Name)
				_f("  if %s_json_unmarshal_field(iter, field, &val%d) {", y.Name, i)
				_f("    out.%s = new(%s)", y.Name, y.Name)
				_f("    *out.%s = val%d", y.Name, i)
				_l("    return true")
				_l("  }")
			case *ast.SelectorExpr:
				alias := nodeToString(y.X)
				if path, ok := allImports[alias]; ok {
					referencedImports[alias] = path
				} else {
					reportError(fmt.Errorf("unknown import: %s", alias))
					return
				}
				if y.Sel.Name == "Number" {
					_l(`  if field == "Number" {`)
					_l("    var val jsoniter.Number")
					_l("    iter.ReadNumber(&val)")
					_f("    out.Number = new(%s)", nodeToString(y))
					_f("    *out.Number = %s(val)", nodeToString(y))
					_l("    return true")
					_l("  }")
				} else {
					typeName := nodeToString(y)
					_f("  var val%d %s", i, typeName)
					_f("  if %s_json_unmarshal_field(iter, field, &val%d) {", typeName, i)
					_f("    out.%s = new(%s)", y.Sel.Name, typeName)
					_f("    *out.%s = val%d", y.Sel.Name, i)
					_l("    return true")
					_l("  }")
				}
			default:
				reportError(fmt.Errorf("unknown embed field type: %s", nodeToString(field.Type)))
				return
			}
		case *ast.SelectorExpr:
			if x.Sel.Name == "Number" {
				_l(`  if field == "Number" { iter.ReadNumber((*jsoniter.Number)(&out.Number)); return true }`)
			} else {
				alias := nodeToString(x.X)
				if path, ok := allImports[alias]; ok {
					referencedImports[alias] = path
				} else {
					reportError(fmt.Errorf("unknown import: %s", alias))
					return
				}
				_f("  if %s_json_unmarshal_field(iter, field, &out.%s) { return true }", nodeToString(x), x.Sel.Name)
			}
		case *ast.Ident: // embed value
			if x.Name == "string" ||
				x.Name == "bool" ||
				x.Name == "int" ||
				x.Name == "uint" ||
				x.Name == "int64" ||
				x.Name == "uint64" ||
				x.Name == "int32" ||
				x.Name == "uint32" ||
				x.Name == "int16" ||
				x.Name == "uint16" ||
				x.Name == "int8" ||
				x.Name == "uint8" ||
				x.Name == "float64" ||
				x.Name == "float32" {
				continue
			}
			_f("  if %s_json_unmarshal_field(iter, field, &out.%s) { return true }", x.Name, x.Name)
		default:
			reportError(fmt.Errorf("unknown embed field type: %s", field.Type))
			return
		}
	}
	oldLines := lines
	isEmptyStruct := true
	_l("  switch {")
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		fieldName := field.Names[0].Name
		encodedFieldName := fieldName
		if field.Tag != nil {
			t, _ := strconv.Unquote(field.Tag.Value)
			tags, err := parseStructTag(t)
			if err != nil {
				reportError(fmt.Errorf("%s: %w", t, err))
				return
			}
			jsonTag := tags["json"]
			if jsonTag != nil && len(jsonTag) > 0 {
				if jsonTag[0] == "-" {
					continue
				}
				encodedFieldName = jsonTag[0]
			}
		}
		isNotExported := unicode.IsLower(rune(fieldName[0])) || fieldName[0] == '_'
		if isNotExported {
			continue
		}
		isEmptyStruct = false
		ptr := fmt.Sprintf("&(*out).%s", fieldName)
		_f("  case field == `%s`:", encodedFieldName)
		genDecodeStmt(field.Type, ptr)
		_l("    return true")
	}
	_l("  }")
	if isEmptyStruct {
		lines = oldLines
		return
	}
}

func genStruct(typeName string, structType *ast.StructType) {
	_l("  more := iter.ReadObjectHead()")
	_l("  for more {")
	_l("    field := iter.ReadObjectField()")
	_f("    if !%s_json_unmarshal_field(iter, field, out) {", typeName)
	_l("      iter.Skip()")
	_l("    }")
	_l("    more = iter.ReadObjectMore()")
	_l("  }")
}

func genArray(typeName string, arrayType *ast.ArrayType) {
	_l("  i := 0")
	_l("  val := *out")
	_l("  more := iter.ReadArrayHead()")
	_l("  for more {")
	if arrayType.Len == nil {
		// slice
		_l(`    if i == len(val) {`)
		_f(`      val = append(val, make(%s, 4)...)`, typeName)
		_l(`    }`)
		ptr := "&val[i]"
		genDecodeStmt(arrayType.Elt, ptr)
	} else {
		// fixed size array
		_f(`    if i < %s {`, nodeToString(arrayType.Len))
		ptr := "&val[i]"
		genDecodeStmt(arrayType.Elt, ptr)
		_l(`    } else {`)
		_l(`      iter.Skip()`)
		_l(`    }`)
	}
	_l(`    i++`)
	_l(`    more = iter.ReadArrayMore()`)
	_l("  }")
	if arrayType.Len == nil {
		// slice
		_l("  if i == 0 {")
		_f("    *out = %s{}", typeName)
		_l("  } else {")
		_l("    *out = val[:i]")
		_l("  }")
	}
}

func genDecodeStmt(node ast.Node, ptr string) {
	switch x := node.(type) {
	case *ast.Ident:
		_f("    "+getDecoder(x.Name), ptr)
	case *ast.ArrayType:
		_f("    "+genAnonymousArray(x), ptr)
	case *ast.StructType:
		_f("    "+genAnonymousStruct(x), ptr)
	case *ast.MapType:
		_f("    "+genAnonymousMap(x), ptr)
	case *ast.StarExpr:
		_f("    "+genAnonymousPtr(x), ptr)
	case *ast.InterfaceType:
		if nodeToString(node) == "interface{}" {
			_f("    iter.ReadInterface(%s)", ptr)
		} else {
			reportError(fmt.Errorf("unknown type: %s", nodeToString(node)))
			return
		}
	case *ast.SelectorExpr:
		if x.Sel.Name == "Number" {
			_f("    iter.ReadNumber((*jsoniter.Number)(%s))", ptr)
		} else {
			alias := nodeToString(x.X)
			if path, ok := allImports[alias]; ok {
				referencedImports[alias] = path
			} else {
				reportError(fmt.Errorf("unknown import: %s", alias))
				return
			}
			_f("    %s_json_unmarshal(iter, %s)", nodeToString(node), ptr)
		}
	default:
		reportError(fmt.Errorf("unknown type: %s", nodeToString(node)))
		return
	}
}

func getDecoder(typeName string) string {
	switch {
	case typeName == "bool":
		return "iter.ReadBool(%s)"
	case typeName == "string":
		return "iter.ReadString(%s)"
	case typeName == "int":
		return "iter.ReadInt(%s)"
	case typeName == "uint":
		return "iter.ReadUint(%s)"
	case typeName == "int64":
		return "iter.ReadInt64(%s)"
	case typeName == "uint64":
		return "iter.ReadUint64(%s)"
	case typeName == "int32":
		return "iter.ReadInt32(%s)"
	case typeName == "uint32":
		return "iter.ReadUint32(%s)"
	case typeName == "int16":
		return "iter.ReadInt16(%s)"
	case typeName == "uint16":
		return "iter.ReadUint16(%s)"
	case typeName == "int8":
		return "iter.ReadInt8(%s)"
	case typeName == "uint8":
		return "iter.ReadUint8(%s)"
	case typeName == "float64":
		return "iter.ReadFloat64(%s)"
	case typeName == "float32":
		return "iter.ReadFloat32(%s)"
	default:
		return typeName + "_json_unmarshal(iter, %s)"
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
	for _, importSpec := range f.Imports {
		path, _ := strconv.Unquote(importSpec.Path.Value)
		if importSpec.Name == nil {
			parts := strings.Split(path, "/")
			allImports[parts[len(parts)-1]] = path
		} else {
			allImports[importSpec.Name.Name] = path
		}
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
