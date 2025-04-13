package divider

import (
	"gioui.org/layout"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

type Divider struct {
	Subheading string
}

func (d Divider) Layout(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if d.Subheading != "" {
		return component.SubheadingDivider(app.Theme, d.Subheading).Layout(gtx)
	}
	return component.Divider(app.Theme).Layout(gtx)
}
