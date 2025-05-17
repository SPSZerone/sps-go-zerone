package item

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
)

func NewDimensions() Dimensions {
	return Dimensions{
		InsetInner:         layout.UniformInset(unit.Dp(4)),
		InsetOuter:         layout.UniformInset(unit.Dp(4)),
		ContentSize:        image.Pt(100, 100),
		Padding:            4,
		HighlightThickness: 4,
		HighlightRoundness: 10,
	}
}

type Dimensions struct {
	InsetInner         layout.Inset
	InsetOuter         layout.Inset
	ContentSize        image.Point
	Padding            int
	HighlightThickness int
	HighlightRoundness int
}
