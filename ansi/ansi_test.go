package ansi

import (
	"fmt"
	"testing"
)

func TestSGR(t *testing.T) {
	testLog := func(content any) {
		//fmt.Println("BEFORE", content, "AFTER")
		t.Log("BEFORE", content, "AFTER")
	}

	type testCase struct {
		content string
		opts    []SGROption
	}

	testCases := []testCase{
		{content: "Normal",
			opts: []SGROption{}},
		{content: "Bold",
			opts: []SGROption{SGRoBold()}},
		{content: "Underline",
			opts: []SGROption{SGRoUnderline()}},

		{content: "Bold & Underline & FGColorBlack & BGColorWhite",
			opts: []SGROption{SGRoBold(), SGRoUnderline(),
				SGRoColor(SGROptFGColorBlack), SGRoColor(SGROptBGColorWhite)}},
		{content: "Bold & Underline & BrightFGColorWhite & BrightBGColorBlue",
			opts: []SGROption{SGRoBold(), SGRoUnderline(),
				SGRoColor(SGROptBrightFGColorWhite), SGRoColor(SGROptBrightBGColorBlue)}},

		{content: "Bold & Underline & FGColorCustomIndex & BGColorCustomIndex",
			opts: []SGROption{SGRoBold(), SGRoUnderline(),
				SGRoFGColorCustomIndex(51), SGRoBGColorCustomIndex(201)}},

		{content: "Bold & Underline & FGColorCustomRGB(255, 255, 0) & BGColorCustomRGB(255, 0, 255)",
			opts: []SGROption{SGRoBold(), SGRoUnderline(),
				SGRoFGColorCustomRGB(255, 255, 0), SGRoBGColorCustomRGB(255, 0, 255)}},
	}

	var sgr SGR
	for _, tCase := range testCases {
		content := tCase.content
		sgr.Options.Reset()
		sgr.Options.Opts(tCase.opts...)
		testLog(sgr.Output(fmt.Sprintf("%s WithReset:TRUE ", content), true))
		testLog(sgr.Output(fmt.Sprintf("%s WithReset:FALSE", content), false))
	}
}
