package item

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
)

func NewDimensions() Dimensions {
	return Dimensions{
		ContentSize: image.Pt(100, 100),

		HighlightThickness: 4,
		HighlightRoundness: 10,

		Padding: 4,
	}
}

type Dimensions struct {
	ContentSize image.Point

	HighlightThickness int
	HighlightRoundness int

	Padding int
}

func ConvertInsetByContext(gtx layout.Context, in layout.Inset) layout.Inset {
	return layout.Inset{
		Left:   unit.Dp(gtx.Dp(in.Left)),
		Right:  unit.Dp(gtx.Dp(in.Right)),
		Top:    unit.Dp(gtx.Dp(in.Top)),
		Bottom: unit.Dp(gtx.Dp(in.Bottom)),
	}
}
