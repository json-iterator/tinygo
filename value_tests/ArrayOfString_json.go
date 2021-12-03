package value_tests

import jsoniter "github.com/json-iterator/tinygo"

func jd_ArrayOfString(iter *jsoniter.Iterator, out *ArrayOfString) {
	if iter.Error != nil {
		return
	}
	i := 0
	val := *out
	more := iter.ReadArrayHead()
	for more {
		if i == len(val) {
			val = append(val, make(ArrayOfString, 4)...)
		}
		iter.ReadString(&val[i])
		i++
		more = iter.ReadArrayMore()
	}
	if i == 0 {
		*out = make(ArrayOfString, 0)
	} else {
		*out = val[:i]
	}
}
