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

func ByteAsciiVisible() byte {
	return Byte(AsciiVisibleStart, AsciiVisibleEnd)
}

func ByteAsciiAlphabetLower() byte {
	return Byte(AsciiLowerCaseA, AsciiLowerCaseZ)
}

func ByteAsciiAlphabetUpper() byte {
	return Byte(AsciiUpperCaseA, AsciiUpperCaseZ)
}

func ByteAsciiNumeric() byte {
	return Byte(AsciiNumeric0, AsciiNumeric9)
}

func Byte(min, max byte) byte {
	return byte(Int(int(min), int(max)))
}
