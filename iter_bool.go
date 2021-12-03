package jsoniter

// ReadBool will assign bool to out if found, otherwise the value will be skipped
func (iter *Iterator) ReadBool(out *bool) error {
	c := iter.nextToken()
	if c == 't' {
		err := iter.skipThreeBytes('r', 'u', 'e')
		if err == nil {
			*out = true
		}
		return err
	}
	if c == 'f' {
		err := iter.skipFourBytes('a', 'l', 's', 'e')
		if err == nil {
			*out = false
		}
		return err
	}
	err := iter.ReportError("ReadBool", "expect t or f, but found "+string([]byte{c}))
	iter.unreadByte()
	iter.Skip()
	return err
}
