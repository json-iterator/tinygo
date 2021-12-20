package generator

import (
	"fmt"
	"go/ast"
	"strconv"
	"unicode"
)

var anonymousDecoders = []byte{}

func generateDecoders(typeSpec *ast.TypeSpec) []byte {
	lines = []byte{}
	typeName := typeSpec.Name.Name
	prefix = typeName
	switch x := typeSpec.Type.(type) {
	case *ast.ArrayType:
		decodeArray(typeName, x)
	case *ast.StructType:
		decodeStruct(typeName, x)
	case *ast.MapType:
		decodeMap(x)
	case *ast.StarExpr:
		decodePtr(x)
	case *ast.Ident:
		genDecodeStmt(x, fmt.Sprintf("(*%s)(out)", x.Name))
	default:
		reportError(fmt.Errorf("not supported type: %s", nodeToString(typeSpec)))
		return nil
	}
	mainDecoder := lines
	lines = []byte{}
	switch x := typeSpec.Type.(type) {
	case *ast.StructType:
		decodeStructField(x)
	case *ast.StarExpr:
		decodeOtherField(typeName)
	case *ast.MapType:
		decodeOtherField(typeName)
	case *ast.ArrayType:
		decodeOtherField(typeName)
	case *ast.Ident:
		decodeOtherField(typeName)
	default:
		reportError(fmt.Errorf("not supported type: %s", nodeToString(typeSpec)))
		return nil
	}
	fieldsDecoder := lines
	lines = []byte{}
	_f("func %s_json_unmarshal(iter *jsoniter.Iterator, out *%s) {", prefix, typeName)
	lines = append(lines, mainDecoder...)
	_l("}")
	_f("func %s_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *%s) bool {", prefix, typeName)
	lines = append(lines, fieldsDecoder...)
	_l("  return false")
	_l("}")
	lines = append(lines, anonymousDecoders...)
	return lines
}

func decodeAnonymousArray(arrayType *ast.ArrayType) string {
	decoderName := fmt.Sprintf(`%s_array%d`, prefix, anonymousCounter)
	typeName := nodeToString(arrayType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, typeName)
	decodeArray(typeName, arrayType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func decodeAnonymousStruct(structType *ast.StructType) string {
	decoderName := fmt.Sprintf(`%s_struct%d`, prefix, anonymousCounter)
	typeName := nodeToString(structType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal_field (iter *jsoniter.Iterator, field string, out *%s) bool {", decoderName, typeName)
	decodeStructField(structType)
	_l("  return false")
	_l("}")
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, typeName)
	decodeStruct(decoderName, structType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func decodeAnonymousMap(mapType *ast.MapType) string {
	decoderName := fmt.Sprintf(`%s_map%d`, prefix, anonymousCounter)
	typeName := nodeToString(mapType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, typeName)
	decodeMap(mapType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func decodeAnonymousPtr(ptrType *ast.StarExpr) string {
	decoderName := fmt.Sprintf(`%s_ptr%d`, prefix, anonymousCounter)
	typeName := nodeToString(ptrType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_unmarshal (iter *jsoniter.Iterator, out *%s) {", decoderName, typeName)
	decodePtr(ptrType)
	_l("}")
	anonymousDecoders = append(anonymousDecoders, lines...)
	lines = oldLines
	return decoderName + "_json_unmarshal(iter, %s)"
}

func decodePtr(ptrType *ast.StarExpr) {
	_f("    var val %s", nodeToString(ptrType.X))
	genDecodeStmt(ptrType.X, "&val")
	_l("    if iter.Error == nil {")
	_l("      *out = &val")
	_l("    }")
}

func decodeMap(mapType *ast.MapType) {
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

func decodeOtherField(expectedFieldName string) {
	_f(`  if field == "%s" {`, expectedFieldName)
	_f("    %s_json_unmarshal(iter, out)", prefix)
	_l("    return true")
	_l("  }")
}

func decodeStructField(structType *ast.StructType) {
	for i, field := range structType.Fields.List {
		if len(field.Names) != 0 {
			continue
		}
		switch x := field.Type.(type) {
		case *ast.StarExpr: // embed pointer
			switch y := x.X.(type) {
			case *ast.Ident:
				isNotExported := unicode.IsLower(rune(y.Name[0]))
				if isNotExported {
					continue
				}
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
				} else if y.Sel.Name == "RawMessage" {
					reportError(fmt.Errorf("embed json.RawMessage is not supported"))
					return
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
			} else if x.Sel.Name == "RawMessage" {
				reportError(fmt.Errorf("embed json.RawMessage is not supported"))
				return
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
			isNotExported := unicode.IsLower(rune(x.Name[0]))
			if isNotExported {
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

func decodeStruct(typeName string, structType *ast.StructType) {
	_l("  more := iter.ReadObjectHead()")
	_l("  for more {")
	_l("    field := iter.ReadObjectField()")
	_f("    if !%s_json_unmarshal_field(iter, field, out) {", typeName)
	_l("      iter.Skip()")
	_l("    }")
	_l("    more = iter.ReadObjectMore()")
	_l("  }")
}

func decodeArray(typeName string, arrayType *ast.ArrayType) {
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
		_f("    "+decodeAnonymousArray(x), ptr)
	case *ast.StructType:
		_f("    "+decodeAnonymousStruct(x), ptr)
	case *ast.MapType:
		_f("    "+decodeAnonymousMap(x), ptr)
	case *ast.StarExpr:
		_f("    "+decodeAnonymousPtr(x), ptr)
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
		} else if x.Sel.Name == "RawMessage" {
			_f("    iter.ReadRawMessage((*jsoniter.RawMessage)(%s))", ptr)
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
