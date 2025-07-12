package rand

import (
	"testing"
)

func TestByte(t *testing.T) {
	const count = 10
	var b byte

	for i := 0; i < count; i++ {
		b = Byte(0, 255)
		t.Logf("Byte(0, 255) = 0x%X | %d ", b, b)
	}

	for i := 0; i < count; i++ {
		b = ByteAsciiAlphabetLower()
		t.Logf("ByteAsciiAlphabetLower = %v | 0x%X | %d", string(b), b, b)
	}

	for i := 0; i < count; i++ {
		b = ByteAsciiAlphabetUpper()
		t.Logf("ByteAsciiAlphabetUpper = %v | 0x%X | %d", string(b), b, b)
	}

	for i := 0; i < count; i++ {
		b = ByteAsciiNumeric()
		t.Logf("ByteAsciiNumeric = %v | 0x%X | %d", string(b), b, b)
	}

	for i := 0; i < count; i++ {
		b = ByteAsciiVisible()
		t.Logf("ByteAsciiVisible = %v | 0x%X | %d", string(b), b, b)
	}
}
