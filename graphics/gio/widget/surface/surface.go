package surface

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func NewDefault() Surface {
	return NewUniformInset(8, 8)
}

func NewUniformInset(inner, outer unit.Dp) Surface {
	return New(
		layout.UniformInset(inner),
		layout.UniformInset(outer),
	)
}

func New(inner, outer layout.Inset) Surface {
	return Surface{
		InnerInset: inner,
		OuterInset: outer,
	}
}

type Style func(surface *Surface, style *component.SurfaceStyle)

type Surface struct {
	InnerInset layout.Inset
	OuterInset layout.Inset
}

func (s *Surface) LayoutDefault(
	theme *material.Theme, gtx layout.Context,
	widget layout.Widget,
) layout.Dimensions {
	return s.Layout(theme, gtx, widget, nil)
}

func (s *Surface) Layout(
	theme *material.Theme, gtx layout.Context,
	widget layout.Widget,
	style Style,
) layout.Dimensions {
	return s.OuterInset.Layout(
		gtx,
		func(gtx layout.Context) layout.Dimensions {
			surfaceStyle := component.Surface(theme)
			if style != nil {
				style(s, &surfaceStyle)
			}
			return surfaceStyle.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return s.InnerInset.Layout(gtx, widget)
			})
		},
	)
}
