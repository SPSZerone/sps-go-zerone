package rand

import (
	"math/rand/v2"
	"time"
)

func NewRandDefault() *rand.Rand {
	timeNow := time.Now()
	return NewRandPCG(uint64(timeNow.Unix()), uint64(timeNow.UnixNano()))
}

func NewRandPCG(seed1, seed2 uint64) *rand.Rand {
	return rand.New(rand.NewPCG(seed1, seed2))
}
