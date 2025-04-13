package divider

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type Divider struct {
	Subheading string
}

func (d Divider) Layout(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	if d.Subheading != "" {
		return component.SubheadingDivider(theme, d.Subheading).Layout(gtx)
	}
	return component.Divider(theme).Layout(gtx)
}
