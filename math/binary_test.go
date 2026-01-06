package math

import (
	"testing"
)

func TestIsPowerOf2(t *testing.T) {
	{
		n := int(-1)
		t.Logf("   int %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := uint(0)
		t.Logf("  uint %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := int8(1)
		t.Logf("  int8 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := uint8(2)
		t.Logf(" uint8 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := int16(3)
		t.Logf(" int16 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := uint16(4)
		t.Logf("uint16 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := int32(5)
		t.Logf(" int32 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := uint32(6)
		t.Logf("uint32 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := int64(7)
		t.Logf(" int64 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
	{
		n := uint64(8)
		t.Logf("uint64 %2d is power of 2:%v", n, IsPowerOf2(n))
	}
}
