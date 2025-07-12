package rand

import (
	"github.com/SPSZerone/sps-go-zerone/generic"
)

var (
	integerRandDefault = NewRandDefault()
)

func Int(min, max int) int {
	return Integer(integerRandDefault.IntN, min, max)
}

func Int32(min, max int32) int32 {
	return Integer(integerRandDefault.Int32N, min, max)
}

func Int64(min, max int64) int64 {
	return Integer(integerRandDefault.Int64N, min, max)
}

func UInt(min, max uint) uint {
	return Integer(integerRandDefault.UintN, min, max)
}

func UInt32(min, max uint32) uint32 {
	return Integer(integerRandDefault.Uint32N, min, max)
}

func UInt64(min, max uint64) uint64 {
	return Integer(integerRandDefault.Uint64N, min, max)
}

// Integer
//
//	funRand func(n T) should return [0,n)
func Integer[T generic.Integer](funRand func(n T) T, min T, max T) T {
	diff := max + 1 - min
	randomNum := funRand(diff) + min
	return randomNum
}
