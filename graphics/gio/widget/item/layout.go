package item

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func NewLayoutContext(gtx layout.Context, dimensions *Dimensions, surface *spssurface.Surface) LayoutContext {
	c := LayoutContext{
		InsetOuter: ConvertInsetByContext(gtx, surface.InsetOuter),
		InsetInner: ConvertInsetByContext(gtx, surface.InsetInner),

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
	InsetOuter layout.Inset
	InsetInner layout.Inset

	ContentSize image.Point

	HighlightThickness       int
	HighlightThicknessDouble int
	HighlightThicknessHalf   int
	HighlightRoundness       int

	Padding       int
	PaddingDouble int

	Offset int

	TotalSizeWithoutInset image.Point
	TotalSize             image.Point
}

func (c *LayoutContext) OnInit() {
	c.HighlightThicknessDouble = c.HighlightThickness << 1
	c.HighlightThicknessHalf = c.HighlightThickness >> 1

	c.PaddingDouble = c.Padding << 1

	c.Offset = c.Padding + c.HighlightThickness

	c.TotalSizeWithoutInset = image.Pt(
		c.ContentSize.X+c.HighlightThicknessDouble+c.PaddingDouble,
		c.ContentSize.Y+c.HighlightThicknessDouble+c.PaddingDouble,
	)
}
