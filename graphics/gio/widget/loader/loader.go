package loader

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func New() Loader {
	return Loader{
		Inset: layout.Inset{
			Top:    unit.Dp(4),
			Bottom: unit.Dp(4),
			Left:   unit.Dp(4),
			Right:  unit.Dp(4),
		},
		Width:  32,
		Height: 32,
	}
}

type Loader struct {
	layout.Inset
	Width  int
	Height int
}

func (l Loader) Layout(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return l.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Max.X = gtx.Dp(unit.Dp(l.Width))
		gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(l.Height))
		return material.Loader(theme).Layout(gtx)
	})
}
