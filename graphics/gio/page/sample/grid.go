package sample

import (
	"gioui.org/layout"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func (p *Page) GridAndItem(app *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.Grid.Layout(app.Theme, gtx, 100, func(gtx layout.Context, index int) layout.Dimensions {
				i := p.GetItem(index)
				highlight := index == p.SelectItem
				dimensions, clicked := i.Layout(app, gtx, highlight)
				if clicked {
					p.SelectItem = index
				}
				return dimensions
			})
		}),
	)
}
