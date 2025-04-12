package pref

import (
	"image/color"

	gioapp "gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

func (p *Page) PrefDecorated(app *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.PrefDecoratedDecorated(app, gtx)
		}),
	}
}

func (p *Page) PrefDecoratedDecorated(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	key := material.Body1(app.Theme, p.decorated.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.Decorated(app, gtx)
	}
	var aWidget, bWidget layout.Widget
	var ratio float32
	if app.Pref.Settings.ValueInFront {
		aWidget, bWidget = value, key
		ratio = RatioValue
	} else {
		aWidget, bWidget = key, value
		ratio = RatioKey
	}
	return spslayout.FlexInset{Ratio: ratio}.LayoutABWidget(gtx, aWidget, bWidget)
}

func (p *Page) Decorated(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.decorated.Widget.Update(gtx) {
		app.Pref.Settings.Decorated = p.decorated.Widget.Value
		app.Window.Option(gioapp.Decorated(p.decorated.Widget.Value))
	}
	return material.Switch(app.Theme, &p.decorated.Widget, p.bottomBar.Desc).Layout(gtx)
}

func decoratedCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	colIdx := TableColIdxValue
	if app.Pref.Settings.ValueInFront {
		colIdx = TableColIdxKey
	}
	switch row {
	case DecoratedRowIdxDecorated:
		switch col {
		case colIdx:
			return p.Decorated(app, gtx)
		default:
			labelStyle.Text = p.decorated.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}
