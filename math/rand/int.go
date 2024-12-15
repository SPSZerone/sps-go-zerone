package rand

import (
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Random
//
//	funRand func(n T) should return [0,n)
func Random[T int | int32 | int64](funRand func(n T) T, min T, max T) T {
	diff := max + 1 - min
	randomNum := funRand(diff) + min
	return randomNum
}

func RandomInt(min int, max int) int {
	return Random(r.Intn, min, max)
}

func RandomInt32(min int32, max int32) int32 {
	return Random(r.Int31n, min, max)
}

func RandomInt64(min int64, max int64) int64 {
	return Random(r.Int63n, min, max)
}
