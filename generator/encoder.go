package generator

import (
	"fmt"
	"go/ast"
	"strconv"
	"unicode"
)

var anonymousEncoders = []byte{}

func generateEncoders(typeSpec *ast.TypeSpec) []byte {
	typeName := typeSpec.Name.Name
	lines = []byte{}
	encodeType(typeName, typeSpec.Type)
	mainEncoder := lines
	lines = []byte{}
	switch x := typeSpec.Type.(type) {
	case *ast.StructType:
		encodeStructField(x)
	case *ast.StarExpr:
		encodeOtherField(typeName)
	case *ast.MapType:
		encodeOtherField(typeName)
	case *ast.ArrayType:
		encodeOtherField(typeName)
	case *ast.Ident:
		encodeOtherField(typeName)
	default:
		reportError(fmt.Errorf("not supported type: %s", nodeToString(typeSpec)))
		return nil
	}
	fieldsEncoder := lines
	lines = []byte{}
	_f("func %s_json_marshal(stream *jsoniter.Stream, val %s) {", prefix, typeName)
	lines = append(lines, mainEncoder...)
	_l("}")
	_f("func %s_json_marshal_field(stream *jsoniter.Stream, val %s) {", prefix, typeName)
	lines = append(lines, fieldsEncoder...)
	_l("}")
	lines = append(lines, anonymousEncoders...)
	return lines
}

func encodeType(typeName string, t ast.Expr) {
	switch x := t.(type) {
	case *ast.ArrayType:
		encodeArray(x)
	case *ast.StructType:
		encodeStruct(typeName, x)
	case *ast.MapType:
		encodeMap(x)
	case *ast.StarExpr:
		genEncodeStmt(x.X, "*val")
	case *ast.Ident:
		genEncodeStmt(x, fmt.Sprintf("(%s)(val)", nodeToString(t)))
	default:
		reportError(fmt.Errorf("not supported type: %s", nodeToString(t)))
		return
	}
}

func encodeAnonymousArray(arrayType *ast.ArrayType) string {
	encoderName := fmt.Sprintf(`%s_array%d`, prefix, anonymousCounter)
	typeName := nodeToString(arrayType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_marshal (stream *jsoniter.Stream, val %s) {", encoderName, typeName)
	encodeArray(arrayType)
	_l("}")
	anonymousEncoders = append(anonymousEncoders, lines...)
	lines = oldLines
	return encoderName + "_json_marshal(stream, %s)"
}

func encodeAnonymousMap(mapType *ast.MapType) string {
	encoderName := fmt.Sprintf(`%s_map%d`, prefix, anonymousCounter)
	typeName := nodeToString(mapType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_marshal (stream *jsoniter.Stream, val %s) {", encoderName, typeName)
	encodeMap(mapType)
	_l("}")
	anonymousEncoders = append(anonymousEncoders, lines...)
	lines = oldLines
	return encoderName + "_json_marshal(stream, %s)"
}

func encodeAnonymousStruct(structType *ast.StructType) string {
	encoderName := fmt.Sprintf(`%s_struct%d`, prefix, anonymousCounter)
	typeName := nodeToString(structType)
	anonymousCounter++
	oldLines := lines
	lines = []byte{}
	_f("func %s_json_marshal_field (stream *jsoniter.Stream, val %s) {", encoderName, typeName)
	encodeStructField(structType)
	_l("}")
	_f("func %s_json_marshal (stream *jsoniter.Stream, val %s) {", encoderName, typeName)
	encodeStruct(encoderName, structType)
	_l("}")
	anonymousEncoders = append(anonymousEncoders, lines...)
	lines = oldLines
	return encoderName + "_json_marshal(stream, %s)"
}

func encodeArray(arrayType *ast.ArrayType) {
	if arrayType.Len != nil {
		arrayLen, err := strconv.ParseInt(nodeToString(arrayType.Len), 10, 64)
		if err == nil && arrayLen < 10 {
			encodeFixedSizeArray(arrayType, int(arrayLen))
			return
		}
	}
	_l("  if len(val) == 0 {")
	_l("    stream.WriteEmptyArray()")
	_l("  } else {")
	_l("    stream.WriteArrayHead()")
	_l("    for i, elem := range val {")
	_l("      if i != 0 { stream.WriteMore() }")
	genEncodeStmt(arrayType.Elt, "elem")
	_l("    }")
	_l("    stream.WriteArrayTail()")
	_l("  }")
}

func encodeFixedSizeArray(arrayType *ast.ArrayType, length int) {
	_l("    stream.WriteArrayHead()")
	for i := 0; i < length; i++ {
		if i != 0 {
			_l("    stream.WriteMore()")
		}
		genEncodeStmt(arrayType.Elt, fmt.Sprintf("val[%d]", i))
	}
	_l("    stream.WriteArrayTail()")
}

func encodeMap(mapType *ast.MapType) {
	_l("  stream.WriteObjectHead()")
	_l("  for k, v := range val {")
	switch x := mapType.Key.(type) {
	case *ast.Ident:
		switch x.Name {
		case "string":
			_l("    stream.WriteObjectField(k)")
		case "int":
			fallthrough
		case "uint":
			fallthrough
		case "int64":
			fallthrough
		case "uint64":
			fallthrough
		case "int32":
			fallthrough
		case "uint32":
			fallthrough
		case "int16":
			fallthrough
		case "uint16":
			fallthrough
		case "int8":
			fallthrough
		case "uint8":
			_l(`      stream.WriteRaw("\"")`)
			genEncodeStmt(mapType.Key, "k")
			_l(`      stream.WriteRaw("\": ")`)
		default:
			reportError(fmt.Errorf("unsupported map key type: %s", nodeToString(mapType.Key)))
		}
	default:
		reportError(fmt.Errorf("unsupported map key type: %s", nodeToString(mapType.Key)))
	}
	genEncodeStmt(mapType.Value, "v")
	_l("    stream.WriteMore()")
	_l("  }")
	_l("  stream.WriteObjectTail()")
}

func encodeStruct(typeName string, structType *ast.StructType) {
	_l("    stream.WriteObjectHead()")
	_f("    %s_json_marshal_field(stream, val)", typeName)
	_l("    stream.WriteObjectTail()")
}

func encodeStructField(structType *ast.StructType) {
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
		_f("    stream.WriteObjectField(`%s`)", encodedFieldName)
		genEncodeStmt(field.Type, fmt.Sprintf("val.%s", fieldName))
		_l("    stream.WriteMore()")
	}
}

func encodeOtherField(typeName string) {
	_f(`    stream.WriteObjectField("%s")`, typeName)
	_f("    %s_json_marshal(stream, val)", prefix)
	_l("    stream.WriteMore()")
}

func genEncodeStmt(node ast.Node, val string) {
	switch x := node.(type) {
	case *ast.Ident:
		_f("    "+getEncoder(x.Name), val)
	case *ast.ArrayType:
		_f("    "+encodeAnonymousArray(x), val)
	case *ast.StructType:
		_f("    "+encodeAnonymousStruct(x), val)
	case *ast.MapType:
		_f("    "+encodeAnonymousMap(x), val)
	case *ast.StarExpr:
		_f("    if %s == nil {", val)
		_l("       stream.WriteNull()")
		_l("    } else {")
		genEncodeStmt(x.X, "*"+val)
		_l("    }")
	case *ast.SelectorExpr:
		if x.Sel.Name == "Number" || x.Sel.Name == "RawMessage" {
			_f("    stream.WriteRawOrNull(string(%s))", val)
		}
	default:
		reportError(fmt.Errorf("unknown type: %s", nodeToString(node)))
		return
	}
}

func getEncoder(typeName string) string {
	switch {
	case typeName == "bool":
		return "stream.WriteBool(%s)"
	case typeName == "string":
		return "stream.WriteString(%s)"
	case typeName == "int":
		return "stream.WriteInt(%s)"
	case typeName == "uint":
		return "stream.WriteUint(%s)"
	case typeName == "int64":
		return "stream.WriteInt64(%s)"
	case typeName == "uint64":
		return "stream.WriteUint64(%s)"
	case typeName == "int32":
		return "stream.WriteInt32(%s)"
	case typeName == "uint32":
		return "stream.WriteUint32(%s)"
	case typeName == "int16":
		return "stream.WriteInt16(%s)"
	case typeName == "uint16":
		return "stream.WriteUint16(%s)"
	case typeName == "int8":
		return "stream.WriteInt8(%s)"
	case typeName == "uint8":
		return "stream.WriteUint8(%s)"
	case typeName == "float64":
		return "stream.WriteFloat64(%s)"
	case typeName == "float32":
		return "stream.WriteFloat32(%s)"
	default:
		return typeName + "_json_marshal(stream, %s)"
	}
}
