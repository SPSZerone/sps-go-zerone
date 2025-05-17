package item

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
)

func NewLayoutContext(gtx layout.Context, dimensions Dimensions) LayoutContext {
	c := LayoutContext{
		InsetOuter: ConvertInsetByContext(gtx, dimensions.InsetOuter),
		InsetInner: ConvertInsetByContext(gtx, dimensions.InsetInner),
		ContentSize: image.Pt(
			gtx.Dp(unit.Dp(dimensions.ContentSize.X)),
			gtx.Dp(unit.Dp(dimensions.ContentSize.Y)),
		),

		HighlightThickness: gtx.Dp(unit.Dp(dimensions.HighlightThickness)),
		HighlightRoundness: gtx.Dp(unit.Dp(dimensions.HighlightRoundness)),

		Padding: gtx.Dp(unit.Dp(dimensions.Padding)),
	}
	c.OnInit()
	return c
}

type LayoutContext struct {
	InsetOuter  layout.Inset
	InsetInner  layout.Inset
	ContentSize image.Point

	HighlightThickness       int
	HighlightThicknessDouble int
	HighlightThicknessHalf   int

	Padding       int
	PaddingDouble int

	HighlightRoundness int

	Offset int
	Size   image.Point
}

func (c *LayoutContext) OnInit() {
	c.HighlightThicknessDouble = c.HighlightThickness << 1
	c.HighlightThicknessHalf = c.HighlightThickness >> 1

	c.PaddingDouble = c.Padding << 1

	c.Offset = c.Padding + c.HighlightThickness
	c.Size = image.Pt(
		c.ContentSize.X+c.HighlightThicknessDouble+c.PaddingDouble,
		c.ContentSize.Y+c.HighlightThicknessDouble+c.PaddingDouble,
	)
}
