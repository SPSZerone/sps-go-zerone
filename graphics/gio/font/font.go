package font

import (
	"gioui.org/font"
	"gioui.org/font/gofont"

	spsnerdfont "github.com/SPSZerone/sps-go-zerone/graphics/gio/font/nerdfont"
)

func Collection() []font.FontFace {
	goFontCollection := gofont.Collection()
	nerdFontCollection := spsnerdfont.Collection()
	return append(goFontCollection, nerdFontCollection...)
}
