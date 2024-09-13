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
