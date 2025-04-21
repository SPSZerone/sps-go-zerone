package label

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func NewLabel() Label {
	return Label{}
}

type Style func(label *Label, style *material.LabelStyle)

type Label struct {
}

func (l *Label) Layout(
	theme *material.Theme, gtx layout.Context,
	size int, txt string,
	style Style,
) layout.Dimensions {
	labelStyle := material.Label(theme, unit.Sp(size), txt)
	if style != nil {
		style(l, &labelStyle)
	}
	return labelStyle.Layout(gtx)
}
