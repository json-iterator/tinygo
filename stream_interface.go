package jsoniter

func (stream *Stream) WriteInterface(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		stream.WriteBool(x)
		return true
	case uint8:
		stream.WriteUint8(x)
		return true
	case int8:
		stream.WriteInt8(x)
		return true
	case uint16:
		stream.WriteUint16(x)
		return true
	case int16:
		stream.WriteInt16(x)
		return true
	case uint32:
		stream.WriteUint32(x)
		return true
	}
	return false
}
