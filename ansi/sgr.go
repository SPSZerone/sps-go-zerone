package ansi

import (
	"fmt"
	"strings"
)

// ==> SGR (Select Graphic Rendition) parameters
// 	The control sequence CSI n m, named Select Graphic Rendition (SGR), sets display attributes.
// 	Several attributes can be set in the same sequence, separated by semicolons.
// 	Each display attribute remains in effect until a following occurrence of SGR resets it.
// 	If no codes are given, CSI m is treated as CSI 0 m (reset / normal).

const (
	SGROptReset      = 0 // Reset or normal	All attributes become turned off
	SGROptBold       = 1 // Bold or increased intensity	As with faint, the color change is a PC (SCO / CGA) invention.[better source needed]
	SGROptFaint      = 2 // Faint, decreased intensity, or dim	May be implemented as a light font weight like bold.
	SGROptItalic     = 3 // Italic	Not widely supported. Sometimes treated as inverse or blink.
	SGROptUnderline  = 4 // Underline	Style extensions exist for Kitty, VTE, mintty, iTerm2 and Konsole.
	SGROptSlowBlink  = 5 // Slow blink	Sets blinking to less than 150 times per minute
	SGROptRapidBlink = 6 // Rapid blink	MS-DOS ANSI.SYS, 150+ per minute; not widely supported
	SGROptReverse    = 7 // Reverse video or invert	Swap foreground and background colors; inconsistent emulation[dubious – discuss]
	SGROptConceal    = 8 // Conceal or hide	Not widely supported.
	SGROptCrossedOut = 9 // Crossed-out, or strike	Characters legible but marked as if for deletion. Not supported in Terminal.app.

	SGROptPrimaryFont          = 10 // Primary (default) font	NULL
	SGROptAlternativeFontStart = 11 // Alternative font	Select alternative font n − 10
	SGROptAlternativeFontEnd   = 19

	SGROptGothic              = 20 // Fraktur (Gothic)	Rarely supported
	SGROptDoublyUnderlined    = 21 // Doubly underlined; or: not bold	Double-underline per ECMA-48, but instead disables bold intensity on several terminals, including in the Linux kernel's console before version 4.17.
	SGROptNormalIntensity     = 22 // Normal intensity	Neither bold nor faint; color changes where intensity is implemented as such.
	SGROptNeitherItalic       = 23 // Neither italic, nor blackletter	NULL
	SGROptNotUnderlined       = 24 // Not underlined	Neither singly nor doubly underlined
	SGROptNotBlinking         = 25 // Not blinking	Turn blinking off
	SGROptProportionalSpacing = 26 // Proportional spacing	ITU T.61 and T.416, not known to be used on terminals
	SGROptNotReversed         = 27 // Not reversed	NULL
	SGROptReveal              = 28 // Reveal	Not concealed
	SGROptNotCrossedOut       = 29 // Not crossed out	NULL

	SGROptFGColorBlack   = 30 // Set foreground color	NULL
	SGROptFGColorRed     = 31
	SGROptFGColorGreen   = 32
	SGROptFGColorYellow  = 33
	SGROptFGColorBlue    = 34
	SGROptFGColorMagenta = 35
	SGROptFGColorCyan    = 36
	SGROptFGColorWhite   = 37
	// SGROptFGColorCustom https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
	SGROptFGColorCustom      = 38 // Set foreground color	Next arguments are 5;n or 2;r;g;b
	SGROptFGColorCustomIndex = "38;5"
	SGROptFGColorCustomRGB   = "38;2"
	SGROptFGColorDefault     = 39 // Default foreground color	Implementation defined (according to standard)

	SGROptBGColorBlack   = 40 // Set background color	NULL
	SGROptBGColorRed     = 41
	SGROptBGColorGreen   = 42
	SGROptBGColorYellow  = 43
	SGROptBGColorBlue    = 44
	SGROptBGColorMagenta = 45
	SGROptBGColorCyan    = 46
	SGROptBGColorWhite   = 47
	// SGROptBGColorCustom https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
	SGROptBGColorCustom      = 48 // Set background color	Next arguments are 5;n or 2;r;g;b
	SGROptBGColorCustomIndex = "48;5"
	SGROptBGColorCustomRGB   = "48;2"
	SGROptBGColorDefault     = 49 // Default background color	Implementation defined (according to standard)

	SGROptDisableProportionalSpacing = 50 // Disable proportional spacing	T.61 and T.416
	SGROptFramed                     = 51 // Framed	Implemented as "emoji variation selector" in mintty.
	SGROptEncircled                  = 52 // Encircled	DITTO
	SGROptOverline                   = 53 // Overlined	Not supported in Terminal.app
	SGROptNeitherFramed              = 54 // Neither framed nor encircled	NULL
	SGROptNotOverline                = 55 // Not overlined	NULL
	SGROptUnderlineColor             = 58 // Set underline color	Not in standard; implemented in Kitty, VTE, mintty, and iTerm2. Next arguments are 5;n or 2;r;g;b.
	SGROptUnderlineColorDefault      = 59 // Default underline color	Not in standard; implemented in Kitty, VTE, mintty, and iTerm2.

	SGROptIdeogramUnderline       = 60 // Ideogram underline or right side line	Rarely supported
	SGROptIdeogramDoubleUnderline = 61 // Ideogram double underline, or double line on the right side	DITTO
	SGROptIdeogramOverline        = 62 // Ideogram overline or left side line	DITTO
	SGROptIdeogramDoubleOverline  = 63 // Ideogram double overline, or double line on the left side	DITTO
	SGROptIdeogramStressMarking   = 64 // Ideogram stress marking	DITTO
	SGROptIdeogramReset           = 65 // No ideogram attributes	Reset the effects of all of 60–64

	SGROptSuperscript        = 73 // Superscript	Implemented only in mintty
	SGROptSubscript          = 74 // Subscript	DITTO
	SGROptNeitherSuperscript = 75 // Neither superscript nor subscript	DITTO

	SGROptBrightFGColorBlack   = 90 // Set bright foreground color	Not in standard; originally implemented by aixterm
	SGROptBrightFGColorRed     = 91
	SGROptBrightFGColorGreen   = 92
	SGROptBrightFGColorYellow  = 93
	SGROptBrightFGColorBlue    = 94
	SGROptBrightFGColorMagenta = 95
	SGROptBrightFGColorCyan    = 96
	SGROptBrightFGColorWhite   = 97

	SGROptBrightBGColorBlack   = 100 // Set bright background color	DITTO
	SGROptBrightBGColorRed     = 101
	SGROptBrightBGColorGreen   = 102
	SGROptBrightBGColorYellow  = 103
	SGROptBrightBGColorBlue    = 104
	SGROptBrightBGColorMagenta = 105
	SGROptBrightBGColorCyan    = 106
	SGROptBrightBGColorWhite   = 107
)

const (
	SGRReset = CSI + "0m"
)

type SGR struct {
	Options SGROptions
}

func (s *SGR) SGREnable() string {
	opts := s.Options.Output()
	sgr := fmt.Sprintf("%s%sm", CSI, opts)
	return sgr
}

func (s *SGR) SGRReset() string {
	return SGRReset
}

func (s *SGR) Output(content any, withReset bool) string {
	if withReset {
		return fmt.Sprintf("%s%v%s", s.SGREnable(), content, s.SGRReset())
	}
	return fmt.Sprintf("%s%v", s.SGREnable(), content)
}

type SGROption func(opts *SGROptions)

type SGROptions struct {
	opts []string
}

func (o *SGROptions) Reset() *SGROptions {
	clear(o.opts)
	return o
}

func (o *SGROptions) Opts(opts ...SGROption) *SGROptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *SGROptions) Output() string {
	return strings.Join(o.opts, ";")
}

func SGROpts(opts ...any) SGROption {
	return func(o *SGROptions) {
		for _, opt := range opts {
			o.opts = append(o.opts, fmt.Sprintf("%v", opt))
		}
	}
}

func SGRoBold() SGROption {
	return func(o *SGROptions) {
		o.opts = append(o.opts, fmt.Sprintf("%v", SGROptBold))
	}
}

func SGRoItalic() SGROption {
	return func(o *SGROptions) {
		o.opts = append(o.opts, fmt.Sprintf("%v", SGROptItalic))
	}
}

func SGRoUnderline() SGROption {
	return func(o *SGROptions) {
		o.opts = append(o.opts, fmt.Sprintf("%v", SGROptUnderline))
	}
}

func SGRoColor(value int) SGROption {
	return func(o *SGROptions) {
		o.opts = append(o.opts, fmt.Sprintf("%v", value))
	}
}

// SGRoColorCustomIndex see [DOC]
//
// index:
//
//	    0-7: standard colors (as in ESC [ 30–37 m)
//	   8-15: high intensity colors (as in ESC [ 90–97 m)
//	 16-231: 6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255: grayscale from dark to light in 24 steps
//
// [DOC]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
func SGRoColorCustomIndex(param string, index int) SGROption {
	return func(o *SGROptions) {
		o.opts = append(o.opts, fmt.Sprintf("%v;%v", param, index))
	}
}

// SGRoFGColorCustomIndex see [DOC]
//
// [DOC]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
func SGRoFGColorCustomIndex(index int) SGROption {
	return SGRoColorCustomIndex(SGROptFGColorCustomIndex, index)
}

// SGRoBGColorCustomIndex see [DOC]
//
// [DOC]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
func SGRoBGColorCustomIndex(index int) SGROption {
	return SGRoColorCustomIndex(SGROptBGColorCustomIndex, index)
}

// SGRoColorCustomRGB see [DOC]
//
// RGB:[0-255]
//
// [DOC]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
func SGRoColorCustomRGB(param string, r, g, b int) SGROption {
	return func(o *SGROptions) {
		o.opts = append(o.opts, fmt.Sprintf("%v;%v;%v;%v", param, r, g, b))
	}
}

// SGRoFGColorCustomRGB see [DOC]
//
// [DOC]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
func SGRoFGColorCustomRGB(r, g, b int) SGROption {
	return SGRoColorCustomRGB(SGROptFGColorCustomRGB, r, g, b)
}

// SGRoBGColorCustomRGB see [DOC]
//
// [DOC]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
func SGRoBGColorCustomRGB(r, g, b int) SGROption {
	return SGRoColorCustomRGB(SGROptBGColorCustomRGB, r, g, b)
}
