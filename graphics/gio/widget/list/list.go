package list

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func New() List {
	return List{
		List: widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

type ScrollBarMode int

const (
	ScrollBarModeAuto ScrollBarMode = iota
	ScrollBarModeEnable
	ScrollBarModeDisable
)

type List struct {
	widget.List

	ScrollBarMode ScrollBarMode
}

func (p *List) Layout(
	theme *material.Theme, gtx layout.Context,
	length int, element layout.ListElement,
) layout.Dimensions {
	var withScrollbar bool

	switch p.ScrollBarMode {
	case ScrollBarModeAuto:
		if p.List.Axis == layout.Horizontal {
			withScrollbar = true
		}
	case ScrollBarModeEnable:
		withScrollbar = true
	case ScrollBarModeDisable:
	}

	if withScrollbar {
		return material.List(theme, &p.List).Layout(gtx, length, element)
	}
	return p.List.Layout(gtx, length, element)
}
