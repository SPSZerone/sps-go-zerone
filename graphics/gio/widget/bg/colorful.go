package bg

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

func NewColorful(color1, color2 color.NRGBA) Colorful {
	return Colorful{
		Color1: color1,
		Color2: color2,
	}
}

type Colorful struct {
	Color1 color.NRGBA
	Color2 color.NRGBA
}

func (c Colorful) LayoutBG(
	theme *material.Theme, gtx layout.Context,
	size image.Point,
) layout.Dimensions {
	offset := image.Pt(
		(gtx.Constraints.Max.X-size.X)>>1,
		(gtx.Constraints.Max.Y-size.Y)>>1,
	)
	rect := image.Rect(
		offset.X, offset.Y,
		offset.X+size.X, offset.Y+size.Y,
	)
	spscolor.FillRect(gtx, rect, c.Color1, c.Color2)
	return layout.Dimensions{
		Size: gtx.Constraints.Max,
	}
}
