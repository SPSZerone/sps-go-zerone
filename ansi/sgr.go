package ansi

// ==> SGR (Select Graphic Rendition) parameters
// 	The control sequence CSI n m, named Select Graphic Rendition (SGR), sets display attributes.
// 	Several attributes can be set in the same sequence, separated by semicolons.
// 	Each display attribute remains in effect until a following occurrence of SGR resets it.
// 	If no codes are given, CSI m is treated as CSI 0 m (reset / normal).

const (
	SGRReset      = 0 // Reset or normal	All attributes become turned off
	SGRBold       = 1 // Bold or increased intensity	As with faint, the color change is a PC (SCO / CGA) invention.[better source needed]
	SGRFaint      = 2 // Faint, decreased intensity, or dim	May be implemented as a light font weight like bold.
	SGRItalic     = 3 // Italic	Not widely supported. Sometimes treated as inverse or blink.
	SGRUnderline  = 4 // Underline	Style extensions exist for Kitty, VTE, mintty, iTerm2 and Konsole.
	SGRSlowBlink  = 5 // Slow blink	Sets blinking to less than 150 times per minute
	SGRRapidBlink = 6 // Rapid blink	MS-DOS ANSI.SYS, 150+ per minute; not widely supported
	SGRReverse    = 7 // Reverse video or invert	Swap foreground and background colors; inconsistent emulation[dubious – discuss]
	SGRConceal    = 8 // Conceal or hide	Not widely supported.
	SGRCrossedOut = 9 // Crossed-out, or strike	Characters legible but marked as if for deletion. Not supported in Terminal.app.

	SGRPrimaryFont          = 10 // Primary (default) font	NULL
	SGRAlternativeFontStart = 11 // Alternative font	Select alternative font n − 10
	SGRAlternativeFontEnd   = 19

	SGRGothic              = 20 // Fraktur (Gothic)	Rarely supported
	SGRDoublyUnderlined    = 21 // Doubly underlined; or: not bold	Double-underline per ECMA-48, but instead disables bold intensity on several terminals, including in the Linux kernel's console before version 4.17.
	SGRNormalIntensity     = 22 // Normal intensity	Neither bold nor faint; color changes where intensity is implemented as such.
	SGRNeitherItalic       = 23 // Neither italic, nor blackletter	NULL
	SGRNotUnderlined       = 24 // Not underlined	Neither singly nor doubly underlined
	SGRNotBlinking         = 25 // Not blinking	Turn blinking off
	SGRProportionalSpacing = 26 // Proportional spacing	ITU T.61 and T.416, not known to be used on terminals
	SGRNotReversed         = 27 // Not reversed	NULL
	SGRReveal              = 28 // Reveal	Not concealed
	SGRNotCrossedOut       = 29 // Not crossed out	NULL

	SGRFGColorStart   = 30 // Set foreground color	NULL
	SGRFGColorEnd     = 37
	SGRFGColorSet     = 38 // Set foreground color	Next arguments are 5;n or 2;r;g;b
	SGRFGColorDefault = 39 // Default foreground color	Implementation defined (according to standard)

	SGRBGColorStart   = 40 // Set background color	NULL
	SGRBGColorEnd     = 47
	SGRBGColorSet     = 48 // Set background color	Next arguments are 5;n or 2;r;g;b
	SGRBGColorDefault = 49 // Default background color	Implementation defined (according to standard)

	SGRDisableProportionalSpacing = 50 // Disable proportional spacing	T.61 and T.416
	SGRFramed                     = 51 // Framed	Implemented as "emoji variation selector" in mintty.
	SGREncircled                  = 52 // Encircled	DITTO
	SGROverline                   = 53 // Overlined	Not supported in Terminal.app
	SGRNeitherFramed              = 54 // Neither framed nor encircled	NULL
	SGRNotOverline                = 55 // Not overlined	NULL
	SGRUnderlineColor             = 58 // Set underline color	Not in standard; implemented in Kitty, VTE, mintty, and iTerm2. Next arguments are 5;n or 2;r;g;b.
	SGRUnderlineColorDefault      = 59 // Default underline color	Not in standard; implemented in Kitty, VTE, mintty, and iTerm2.

	SGRIdeogramUnderline       = 60 // Ideogram underline or right side line	Rarely supported
	SGRIdeogramDoubleUnderline = 61 // Ideogram double underline, or double line on the right side	DITTO
	SGRIdeogramOverline        = 62 // Ideogram overline or left side line	DITTO
	SGRIdeogramDoubleOverline  = 63 // Ideogram double overline, or double line on the left side	DITTO
	SGRIdeogramStressMarking   = 64 // Ideogram stress marking	DITTO
	SGRIdeogramReset           = 65 // No ideogram attributes	Reset the effects of all of 60–64

	SGRSuperscript        = 73 // Superscript	Implemented only in mintty
	SGRSubscript          = 74 // Subscript	DITTO
	SGRNeitherSuperscript = 75 // Neither superscript nor subscript	DITTO

	SGRBrightFGColorStart = 90 // Set bright foreground color	Not in standard; originally implemented by aixterm
	SGRBrightFGColorEnd   = 97
	SGRBrightBGColorStart = 100 // Set bright background color	DITTO
	SGRBrightBGColorEnd   = 107
)
