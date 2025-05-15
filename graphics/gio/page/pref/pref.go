package pref

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

func (p *Page) LayoutPref(win spsgio.Window, gtx layout.Context, param any) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutPrefTableStyle(win, gtx)
		}),
	)
}

func (p *Page) LayoutPrefTableStyle(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	theme := win.GetTheme()
	pref := win.GetPref()

	key := material.Body1(theme, pref.Pref.TableStyle.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.TableStyle(win, gtx)
	}
	var aWidget, bWidget layout.Widget
	var ratio float32
	if pref.Settings.ValueInFront.Value {
		aWidget, bWidget = value, key
		ratio = RatioValue
	} else {
		aWidget, bWidget = key, value
		ratio = RatioKey
	}
	return spslayout.FlexInset{}.LayoutFlexedWidgetAB(gtx, ratio, aWidget, bWidget)
}

func (p *Page) TableStyle(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	return pref.Pref.TableStyle.LayoutSwitch(
		window, theme, gtx,
	)
}
