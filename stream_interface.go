package jsoniter

func (stream *Stream) WriteInterface(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		stream.WriteBool(x)
		return true
	case uint:
		stream.WriteUint(x)
		return true
	case int:
		stream.WriteInt(x)
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
	case int32:
		stream.WriteInt32(x)
		return true
	case uint64:
		stream.WriteUint64(x)
		return true
	case int64:
		stream.WriteInt64(x)
		return true
	case float32:
		stream.WriteFloat32(x)
		return true
	case float64:
		stream.WriteFloat64(x)
		return true
	}
	return false
}
