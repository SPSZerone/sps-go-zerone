package pref

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

func (p *Page) PrefSettings(app *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.PrefSettingsNonModalDrawer(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.PrefSettingsValueInFront(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.PrefSettingsBottomBar(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.PrefSettingsPrefTableStyle(app, gtx)
		}),
	}
}

func (p *Page) PrefSettingsNonModalDrawer(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	key := material.Body1(app.Theme, p.nonModalDrawer.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.NonModalDrawer(app, gtx)
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

func (p *Page) PrefSettingsValueInFront(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	key := material.Body1(app.Theme, p.valueInFront.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.ValueInFront(app, gtx)
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

func (p *Page) PrefSettingsBottomBar(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	key := material.Body1(app.Theme, p.bottomBar.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.BottomBar(app, gtx)
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

func (p *Page) PrefSettingsPrefTableStyle(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	key := material.Body1(app.Theme, p.prefTableStyle.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return p.TableStyle(app, gtx)
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

func (p *Page) NonModalDrawer(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.nonModalDrawer.Widget.Update(gtx) {
		app.Pref.Settings.NonModalDrawer = p.nonModalDrawer.Widget.Value
		if p.nonModalDrawer.Widget.Value {
			p.Pages.NavAnim.Appear(gtx.Now)
		} else {
			p.Pages.NavAnim.Disappear(gtx.Now)
		}
	}
	return material.Switch(app.Theme, &p.nonModalDrawer.Widget, p.bottomBar.Desc).Layout(gtx)
}

func (p *Page) ValueInFront(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.valueInFront.Widget.Update(gtx) {
		app.Pref.Settings.ValueInFront = p.valueInFront.Widget.Value
	}
	return material.Switch(app.Theme, &p.valueInFront.Widget, p.valueInFront.Desc).Layout(gtx)
}

func (p *Page) BottomBar(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.bottomBar.Widget.Update(gtx) {
		if p.bottomBar.Widget.Value {
			p.Pages.ModalNavDrawer.Anchor = component.Bottom
			p.Pages.AppBar.Anchor = component.Bottom
		} else {
			p.Pages.ModalNavDrawer.Anchor = component.Top
			p.Pages.AppBar.Anchor = component.Top
		}
		app.Pref.Settings.BottomBar = p.bottomBar.Widget.Value
	}
	return material.Switch(app.Theme, &p.bottomBar.Widget, p.bottomBar.Desc).Layout(gtx)
}

func (p *Page) TableStyle(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.prefTableStyle.Widget.Update(gtx) {
		app.Pref.Settings.PrefTableStyle = p.prefTableStyle.Widget.Value
	}
	return material.Switch(app.Theme, &p.prefTableStyle.Widget, p.prefTableStyle.Desc).Layout(gtx)
}

func settingsCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	colIdx := TableColIdxValue
	if app.Pref.Settings.ValueInFront {
		colIdx = TableColIdxKey
	}
	switch row {
	case SettingsRowIdxNonModalDrawer:
		switch col {
		case colIdx:
			return p.NonModalDrawer(app, gtx)
		default:
			labelStyle.Text = p.nonModalDrawer.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxValueInFront:
		switch col {
		case colIdx:
			return p.ValueInFront(app, gtx)
		default:
			labelStyle.Text = p.valueInFront.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxBottomBar:
		switch col {
		case colIdx:
			return p.BottomBar(app, gtx)
		default:
			labelStyle.Text = p.bottomBar.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxTableStyle:
		switch col {
		case colIdx:
			return p.TableStyle(app, gtx)
		default:
			labelStyle.Text = p.prefTableStyle.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}
