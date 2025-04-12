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
	if app.Pref.Settings.ValueInFront {
		return []layout.FlexChild{
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{
					Ratio: RatioValue,
				}.LayoutABWidget(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return p.Decorated(app, gtx)
					},
					material.Body1(app.Theme, p.decorated.Name).Layout,
				)
			}),
		}
	}

	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{
				Ratio: RatioKey,
			}.LayoutABWidget(gtx,
				material.Body1(app.Theme, p.decorated.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.Decorated(app, gtx)
				},
			)
		}),
	}
}

func (p *Page) Decorated(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.decorated.Widget.Update(gtx) {
		app.Pref.Settings.Decorated = p.decorated.Widget.Value
		app.Window.Option(gioapp.Decorated(p.decorated.Widget.Value))
	}
	return material.Switch(app.Theme, &p.decorated.Widget, p.bottomBar.Desc).Layout(gtx)
}

func decoratedCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	switch row {
	case DecoratedRowIdxDecorated:
		switch col {
		case TableColIdxValue:
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
