package slice

import "unsafe"

func IsSame(s1, s2 []byte) bool {
	return unsafe.SliceData(s1) == unsafe.SliceData(s2)
}
