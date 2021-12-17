package jsoniter

// WriteArrayHead write [ with possible indention
func (stream *Stream) WriteArrayHead() {
	stream.writeByte('[')
	stream.indentCount += 1
	stream.writeIndent()
}

// WriteEmptyArray write []
func (stream *Stream) WriteEmptyArray() {
	stream.writeTwoBytes('[', ']')
}

// WriteArrayTail write ] with possible indention
func (stream *Stream) WriteArrayTail() {
	stream.indentCount -= 1
	stream.writeIndent()
	stream.writeByte(']')
}
