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
	stream.indentCount -= 1
	stream.writeIndent()
	stream.writeByte('}')
}

// WriteEmptyObject write {}
func (stream *Stream) WriteEmptyObject() {
	stream.writeByte('{')
	stream.writeByte('}')
}
