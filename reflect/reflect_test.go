package reflect

import (
	"math"
	"testing"
)

type TestNoArg struct {
}

func (t *TestNoArg) Name() string {
	return "TestNoArg"
}

func (t *TestNoArg) Byte() byte {
	return math.MaxInt8 >> 1
}

func (t *TestNoArg) Int() int {
	return math.MaxInt
}

func (t *TestNoArg) UInt() uint {
	return math.MaxUint
}

func (t *TestNoArg) Int8() int8 {
	return math.MaxInt8
}

func (t *TestNoArg) UInt8() uint8 {
	return math.MaxUint8
}

func (t *TestNoArg) Int16() int16 {
	return math.MaxInt16
}

func (t *TestNoArg) UInt16() uint16 {
	return math.MaxUint16
}

func (t *TestNoArg) Int32() int32 {
	return math.MaxInt32
}

func (t *TestNoArg) UInt32() uint32 {
	return math.MaxUint32
}

func (t *TestNoArg) Int64() int64 {
	return math.MaxInt64
}

func (t *TestNoArg) UInt64() uint64 {
	return math.MaxUint64
}

func (t *TestNoArg) Float32() float32 {
	return math.MaxFloat32
}

func (t *TestNoArg) Float64() float64 {
	return math.MaxFloat64
}

func (t *TestNoArg) Complex64() complex64 {
	return 1 + 1i
}

func (t *TestNoArg) Complex128() complex128 {
	return 2 + 2i
}

func TestMethod(t *testing.T) {
	var r TestNoArg
	methods := Methods(&r)
	for _, m := range methods {
		name := m.Name
		v := CallResultByName(&r, name)
		t.Logf("%v %+v", name, v)
	}
}
