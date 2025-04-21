package caption

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Style func(caption *Caption, style *material.LabelStyle)

type Caption struct {
}

func (c *Caption) Layout(
	theme *material.Theme, gtx layout.Context,
	txt string,
	style Style,
) layout.Dimensions {
	labelStyle := material.Caption(theme, txt)
	if style != nil {
		style(c, &labelStyle)
	}
	return labelStyle.Layout(gtx)
}
