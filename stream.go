package jsoniter

type Stream struct {
	buf         []byte
	indentCount int
	Error       error
	Prefix      string
	Indent      string
}

// NewStream create new stream instance.
func NewStream() *Stream {
	return &Stream{
		buf:         make([]byte, 0, 16),
		indentCount: 0,
	}
}

// Buffer if writer is nil, use this method to take the result
func (stream *Stream) Buffer() []byte {
	return stream.buf
}

// Write writes the contents of p into the buffer.
// It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining
// why the write is short.
func (stream *Stream) Write(p []byte) (nn int, err error) {
	stream.buf = append(stream.buf, p...)
	return len(p), nil
}

// WriteByte writes a single byte.
func (stream *Stream) writeByte(c byte) {
	stream.buf = append(stream.buf, c)
}

func (stream *Stream) writeTwoBytes(c1 byte, c2 byte) {
	stream.buf = append(stream.buf, c1, c2)
}

func (stream *Stream) writeThreeBytes(c1 byte, c2 byte, c3 byte) {
	stream.buf = append(stream.buf, c1, c2, c3)
}

func (stream *Stream) writeFourBytes(c1 byte, c2 byte, c3 byte, c4 byte) {
	stream.buf = append(stream.buf, c1, c2, c3, c4)
}

func (stream *Stream) writeFiveBytes(c1 byte, c2 byte, c3 byte, c4 byte, c5 byte) {
	stream.buf = append(stream.buf, c1, c2, c3, c4, c5)
}

// WriteRaw write string out without quotes, just like []byte
func (stream *Stream) WriteRaw(s string) {
	stream.buf = append(stream.buf, s...)
}

func (stream *Stream) WriteRawOrNull(s string) {
	if s == "" {
		stream.WriteNull()
	} else {
		stream.WriteRaw(s)
	}
}

func (stream *Stream) WriteRawOrZero(s string) {
	if s == "" {
		stream.WriteRaw("0")
	} else {
		stream.WriteRaw(s)
	}
}

// WriteNull write null to stream
func (stream *Stream) WriteNull() {
	stream.writeFourBytes('n', 'u', 'l', 'l')
}

// WriteMore write , with possible indention
func (stream *Stream) WriteMore() {
	stream.writeByte(',')
	stream.writeIndent()
}

func (stream *Stream) withIndent() bool {
	return len(stream.Prefix) > 0 || len(stream.Indent) > 0
}

func (stream *Stream) writeIndent() {
	if stream.withIndent() {
		stream.writeByte('\n')
	}
	stream.buf = append(stream.buf, stream.Prefix...)
	for i := 0; i < stream.indentCount; i++ {
		stream.buf = append(stream.buf, stream.Indent...)
	}
}

func (stream *Stream) reportError(err error) error {
	if stream.Error == nil {
		stream.Error = err
	}
	return err
}
