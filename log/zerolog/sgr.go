package zerolog

import (
	"github.com/rs/zerolog"

	"github.com/SPSZerone/sps-go-zerone/ansi"
)

func SGRTimestamp() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoColor(ansi.SGROptFGColorBlack), ansi.SGRoColor(ansi.SGROptBrightBGColorCyan)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRLevel(level zerolog.Level) ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoColor(ansi.SGROptFGColorBlack)}
	sgr.Options.Opts(opts...)

	switch level {
	case zerolog.DebugLevel:
		sgr.Options.Opts(ansi.SGRoColor(ansi.SGROptBrightBGColorCyan))
	case zerolog.InfoLevel:
		sgr.Options.Opts(ansi.SGRoColor(ansi.SGROptBrightBGColorGreen))
	case zerolog.WarnLevel:
		sgr.Options.Opts(ansi.SGRoColor(ansi.SGROptBrightBGColorYellow))
	case zerolog.ErrorLevel:
		sgr.Options.Opts(ansi.SGRoColor(ansi.SGROptBrightBGColorRed))
	case zerolog.FatalLevel:
		sgr.Options.Opts(ansi.SGRoBGColorCustomIndex(88))
	default:
		sgr.Options.Opts(ansi.SGRoBGColorCustomIndex(52))
	}
	return sgr
}

func SGRCaller() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoUnderline(), ansi.SGRoFGColorCustomIndex(165), ansi.SGRoColor(ansi.SGROptBrightBGColorWhite)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRMessage() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoColor(ansi.SGROptFGColorBlack), ansi.SGRoColor(ansi.SGROptBrightBGColorWhite)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRFieldName() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoItalic(), ansi.SGRoFGColorCustomRGB(0xFF, 0xFF, 0xFF), ansi.SGRoColor(ansi.SGROptBrightBGColorGreen)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRFieldSeparator() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoFGColorCustomRGB(0x0, 0x0, 0x0), ansi.SGRoColor(ansi.SGROptBrightBGColorGreen)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRFieldValue() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoItalic(), ansi.SGRoFGColorCustomRGB(0xFF, 0xFF, 0xFF), ansi.SGRoColor(ansi.SGROptBrightBGColorGreen)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRErrFieldName() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoItalic(), ansi.SGRoFGColorCustomRGB(0xFF, 0xFF, 0xFF), ansi.SGRoColor(ansi.SGROptBrightBGColorRed)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRErrFieldSeparator() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoFGColorCustomRGB(0x0, 0x0, 0x0), ansi.SGRoColor(ansi.SGROptBrightBGColorRed)}
	sgr.Options.Opts(opts...)
	return sgr
}

func SGRErrFieldValue() ansi.SGR {
	var sgr ansi.SGR
	opts := []ansi.SGROption{ansi.SGRoBold(), ansi.SGRoItalic(), ansi.SGRoFGColorCustomRGB(0xFF, 0xFF, 0xFF), ansi.SGRoColor(ansi.SGROptBrightBGColorRed)}
	sgr.Options.Opts(opts...)
	return sgr
}
