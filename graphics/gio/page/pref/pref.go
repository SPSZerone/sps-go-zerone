package pref

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

func (p *Page) LayoutPref(app *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutPrefTableStyle(app, gtx)
		}),
	}
}

func (p *Page) LayoutPrefTableStyle(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	key := material.Body1(app.Theme, app.Pref.Pref.TableStyle.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.TableStyle(app, gtx)
	}
	var aWidget, bWidget layout.Widget
	var ratio float32
	if app.Pref.Settings.ValueInFront.Value {
		aWidget, bWidget = value, key
		ratio = RatioValue
	} else {
		aWidget, bWidget = key, value
		ratio = RatioKey
	}
	return spslayout.FlexInset{Ratio: ratio}.LayoutABWidget(gtx, aWidget, bWidget)
}

func (p *Page) TableStyle(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	return app.Pref.Pref.TableStyle.LayoutSwitch(
		app.Window, app.Theme, gtx,
	)
}
