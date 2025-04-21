package clickable

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Clickable struct {
	widget.Clickable
}

func (c *Clickable) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return material.Clickable(gtx, &c.Clickable, widget)
}
