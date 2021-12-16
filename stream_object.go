package jsoniter

// WriteObjectStart write { with possible indention
func (stream *Stream) WriteObjectStart() {
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

// WriteObjectEnd write } with possible indention
func (stream *Stream) WriteObjectEnd() {
	stream.indentCount -= 1
	stream.writeIndent()
	stream.writeByte('}')
}

// WriteEmptyObject write {}
func (stream *Stream) WriteEmptyObject() {
	stream.writeByte('{')
	stream.writeByte('}')
}
