package pref

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

func (p *Page) LayoutPref(win *spswin.Window, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutPrefTableStyle(win, gtx)
		}),
	}
}

func (p *Page) LayoutPrefTableStyle(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	key := material.Body1(win.Theme, win.Pref.Pref.TableStyle.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.TableStyle(win, gtx)
	}
	var aWidget, bWidget layout.Widget
	var ratio float32
	if win.Pref.Settings.ValueInFront.Value {
		aWidget, bWidget = value, key
		ratio = RatioValue
	} else {
		aWidget, bWidget = key, value
		ratio = RatioKey
	}
	return spslayout.FlexInset{Ratio: ratio}.LayoutABWidget(gtx, aWidget, bWidget)
}

func (p *Page) TableStyle(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	return win.Pref.Pref.TableStyle.LayoutSwitch(
		win.Window, win.Theme, gtx,
	)
}
