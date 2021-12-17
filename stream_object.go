package jsoniter

// WriteObjectHead write { with possible indention
func (stream *Stream) WriteObjectHead() {
	stream.indentCount += 1
	stream.writeByte('{')
	stream.writeIndent()
}

// WriteObjectField write "field": with possible indention
func (stream *Stream) WriteObjectField(field string) {
	stream.WriteString(field)
	if stream.withIndent() {
		stream.writeTwoBytes(':', ' ')
	} else {
		stream.writeByte(':')
	}
}

// WriteObjectTail write } with possible indention
func (stream *Stream) WriteObjectTail() {
	lookBack := len(stream.buf) - 1 - len(stream.Prefix) - len(stream.Indent)*stream.indentCount
	if stream.withIndent() {
		lookBack -= 1
	}
	lookBackChar := stream.buf[lookBack]
	stream.indentCount -= 1
	if lookBackChar == ',' {
		stream.buf = stream.buf[:lookBack]
	}
	if lookBackChar == '{' {
		stream.buf = stream.buf[:lookBack+1]
	} else {
		stream.writeIndent()
	}
	stream.writeByte('}')
}
