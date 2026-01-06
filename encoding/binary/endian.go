package binary

import (
	"unsafe"
)

func IsLittleEndian() bool {
	var i int16 = 0x0001
	ptr := unsafe.Pointer(&i)
	firstByte := *(*byte)(ptr)
	return firstByte == 0x01
}
