package surface

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func NewDefault() Surface {
	return NewUniformInset(8, 8)
}

func NewUniformInset(outer, inner unit.Dp) Surface {
	return New(
		layout.UniformInset(outer),
		layout.UniformInset(inner),
	)
}

func New(outer, inner layout.Inset) Surface {
	return Surface{
		InsetOuter:  outer,
		InsetInner:  inner,
		BGColor:     color.NRGBA{R: 0xD3, G: 0xD3, B: 0xD3, A: 0xFF},
		ShadowStyle: component.Shadow(unit.Dp(10), unit.Dp(5)),
	}
}

type Style func(surface *Surface, style *component.SurfaceStyle)

type Surface struct {
	InsetOuter layout.Inset
	InsetInner layout.Inset

	BGColor     color.NRGBA
	ShadowStyle component.ShadowStyle
}

func (s Surface) LayoutDefault(
	theme *material.Theme, gtx layout.Context,
	widget layout.Widget,
) layout.Dimensions {
	return s.Layout(theme, gtx, widget, nil)
}

func (s Surface) Layout(
	theme *material.Theme, gtx layout.Context,
	widget layout.Widget,
	style Style,
) layout.Dimensions {
	return s.InsetOuter.Layout(
		gtx,
		func(gtx layout.Context) layout.Dimensions {
			surfaceStyle := component.Surface(theme)
			surfaceStyle.Fill = s.BGColor
			surfaceStyle.ShadowStyle = s.ShadowStyle

			if style != nil {
				style(&s, &surfaceStyle)
			}

			return surfaceStyle.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return s.InsetInner.Layout(gtx, widget)
			})
		},
	)
}
