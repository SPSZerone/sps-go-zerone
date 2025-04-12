package sample

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func (p *Page) GridAndItem(app *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.Grid.Layout(gtx, 100, func(gtx layout.Context, index int) layout.Dimensions {
				i := p.GetItem(index)
				highlight := index == p.SelectItem
				dimensions, clicked := i.Layout(app, gtx, highlight, func(gtx layout.Context, layoutCtx spsitem.LayoutContext) layout.Dimensions {
					baseInfo := material.Body1(app.Theme, fmt.Sprintf("%v", i.Data))
					baseInfo.Font.Style = font.Italic
					baseInfo.Font.Weight = font.Bold
					return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
				})
				if clicked {
					p.SelectItem = index
				}
				return dimensions
			})
		}),
	)
}
