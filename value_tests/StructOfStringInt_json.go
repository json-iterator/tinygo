package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func jd_StructOfStringInt(iter *jsoniter.Iterator, out *StructOfStringInt) {
	if iter.Error != nil {
		return
	}
	more := iter.ReadObjectHead()
	for more {
		field := iter.ReadObjectField()
		if field == "Name" {
			iter.ReadString(&(*out).Name)
		} else if field == "Price" {
			iter.ReadInt(&(*out).Price)
		} else {
			iter.Skip()
		}
		more = iter.ReadObjectMore()
	}
}
