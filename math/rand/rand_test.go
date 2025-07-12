package rand

import (
	"testing"
)

func TestRandomByte(t *testing.T) {
	const count = 10
	var b byte

	for i := 0; i < count; i++ {
		b = RandomByte(0, 255)
		t.Logf("RandomByte(0, 255) = 0x%X | %d ", b, b)
	}

	for i := 0; i < count; i++ {
		b = RandomByteAsciiAlphabetLower()
		t.Logf("RandomByteAsciiAlphabetLower = %v | 0x%X | %d", string(b), b, b)
	}

	for i := 0; i < count; i++ {
		b = RandomByteAsciiAlphabetUpper()
		t.Logf("RandomByteAsciiAlphabetUpper = %v | 0x%X | %d", string(b), b, b)
	}

	for i := 0; i < count; i++ {
		b = RandomByteAsciiNumeric()
		t.Logf("RandomByteAsciiNumeric = %v | 0x%X | %d", string(b), b, b)
	}

	for i := 0; i < count; i++ {
		b = RandomByteAsciiVisible()
		t.Logf("RandomByteAsciiVisible = %v | 0x%X | %d", string(b), b, b)
	}
}
