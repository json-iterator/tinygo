package jsoniter

// WriteArrayStart write [ with possible indention
func (stream *Stream) WriteArrayStart() {
	stream.writeByte('[')
	stream.indentCount += 1
	stream.writeIndent()
}

// WriteEmptyArray write []
func (stream *Stream) WriteEmptyArray() {
	stream.writeTwoBytes('[', ']')
}

// WriteArrayEnd write ] with possible indention
func (stream *Stream) WriteArrayEnd() {
	stream.indentCount -= 1
	stream.writeIndent()
	stream.writeByte(']')
}
