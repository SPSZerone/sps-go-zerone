package rand

const (
	AsciiVisibleStart = byte(0x20)
	AsciiVisibleEnd   = byte(0x7E)
	AsciiLowerCaseA   = byte('a')
	AsciiLowerCaseZ   = byte('z')
	AsciiUpperCaseA   = byte('A')
	AsciiUpperCaseZ   = byte('Z')
	AsciiNumeric0     = byte('0')
	AsciiNumeric9     = byte('9')
)

func RandomByteAsciiVisible() byte {
	return RandomByte(AsciiVisibleStart, AsciiVisibleEnd)
}

func RandomByteAsciiAlphabetLower() byte {
	return RandomByte(AsciiLowerCaseA, AsciiLowerCaseZ)
}

func RandomByteAsciiAlphabetUpper() byte {
	return RandomByte(AsciiUpperCaseA, AsciiUpperCaseZ)
}

func RandomByteAsciiNumeric() byte {
	return RandomByte(AsciiNumeric0, AsciiNumeric9)
}

func RandomByte(min, max byte) byte {
	return byte(RandomInt(int(min), int(max)))
}
