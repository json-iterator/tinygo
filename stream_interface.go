package jsoniter

func (stream *Stream) WriteInterface(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		stream.WriteBool(x)
		return true
	case uint8:
		stream.WriteUint8(x)
		return true
	}
	return false
}
