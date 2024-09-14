package ansi

// https://en.wikipedia.org/wiki/ANSI_escape_code

// C0 control codes

const (
	C0CCiBEL = 0x07 // Bell
	C0CCiBS  = 0x08 // Backspace
	C0CCiHT  = 0x09 // Tab
	C0CCiLF  = 0x0A // Line Feed
	C0CCiFF  = 0x0C // Form Feed
	C0CCiCR  = 0x0D // Carriage Return
	C0CCiESC = 0x1B // Escape
)

const (
	C0CCsBEL = "\x07" // Bell
	C0CCsBS  = "\x08" // Backspace
	C0CCsHT  = "\x09" // Tab
	C0CCsLF  = "\x0A" // Line Feed
	C0CCsFF  = "\x0C" // Form Feed
	C0CCsCR  = "\x0D" // Carriage Return
	C0CCsESC = "\x1B" // Escape
)

// ==> Fe Escape sequences

const (
	FeEscSEQiSS2 = 0x8E // ESC N | Single Shift Two
	FeEscSEQiSS3 = 0x8F // ESC O | Single Shift Three
	FeEscSEQiDCS = 0x90 // ESC P | Device Control String
	FeEscSEQiCSI = 0x9B // ESC [ | Control Sequence Introducer
	FeEscSEQiST  = 0x9C // ESC \ | String Terminator
	FeEscSEQiOSC = 0x9D // ESC ] | Operating System Command
	FeEscSEQiSOS = 0x98 // ESC X | Start of String
	FeEscSEQiPM  = 0x9E // ESC ^ | Privacy Message
	FeEscSEQiAPC = 0x9F // ESC _ | Application Program Command
)

const (
	FeEscSEQsSS2 = "\x8E" // ESC N | Single Shift Two
	FeEscSEQsSS3 = "\x8F" // ESC O | Single Shift Three
	FeEscSEQsDCS = "\x90" // ESC P | Device Control String
	FeEscSEQsCSI = "\x9B" // ESC [ | Control Sequence Introducer
	FeEscSEQsST  = "\x9C" // ESC \ | String Terminator
	FeEscSEQsOSC = "\x9D" // ESC ] | Operating System Command
	FeEscSEQsSOS = "\x98" // ESC X | Start of String
	FeEscSEQsPM  = "\x9E" // ESC ^ | Privacy Message
	FeEscSEQsAPC = "\x9F" // ESC _ | Application Program Command
)

// ==> CSI (Control Sequence Introducer) sequences
// 	Code	Abbr	Name	Effect
// 	CSI n m	SGR	Select Graphic Rendition	Sets colors and style of the characters following this code

const (
	CSI = C0CCsESC + "["
)
