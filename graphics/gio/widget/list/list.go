package list

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func NewList() List {
	return List{
		List: widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

type List struct {
	widget.List
}

func (p *List) Layout(
	theme *material.Theme, gtx layout.Context,
	length int, element layout.ListElement,
) layout.Dimensions {
	return material.List(theme, &p.List).Layout(gtx, length, element)
}
