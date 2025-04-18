package surface

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func NewSurface() Surface {
	return Surface{
		InnerInset: layout.UniformInset(unit.Dp(8)),
		OuterInset: layout.UniformInset(unit.Dp(8)),
	}
}

type Surface struct {
	InnerInset layout.Inset
	OuterInset layout.Inset
}

func (s Surface) Layout(theme *material.Theme, gtx layout.Context, w layout.Widget) layout.Dimensions {
	return s.OuterInset.Layout(
		gtx,
		func(gtx layout.Context) layout.Dimensions {
			return component.Surface(theme).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return s.InnerInset.Layout(gtx, w)
			})
		},
	)
}
