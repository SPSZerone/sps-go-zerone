package clickable

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func New() Clickable {
	return Clickable{}
}

type Clickable struct {
	widget.Clickable
}

func (c *Clickable) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return material.Clickable(gtx, &c.Clickable, widget)
}

func (c *Clickable) GetClickable() *widget.Clickable {
	return &c.Clickable
}
