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

func ConvertInsetByContext(gtx layout.Context, in layout.Inset) layout.Inset {
	return layout.Inset{
		Left:   unit.Dp(gtx.Dp(in.Left)),
		Right:  unit.Dp(gtx.Dp(in.Right)),
		Top:    unit.Dp(gtx.Dp(in.Top)),
		Bottom: unit.Dp(gtx.Dp(in.Bottom)),
	}
}
