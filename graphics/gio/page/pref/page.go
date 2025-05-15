package pref

import (
	"image/color"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spsdiscloser "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/discloser"
	spslist "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/list"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
	spstable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/table"
)

func New(pages spsgio.Pages, app spsgio.App, newWindow spsgio.NewWindow) *Page {
	tabs := spstab.NewByNames(
		[]string{TabNameSettings, TabNamePreferences},
		spstab.OptAxisSetting(&pages.GetWindow().GetPref().Settings.TabAxis),
	)
	p := &Page{
		Pages: pages,

		Tabs:  tabs,
		Table: spstable.New(),

		app:       app,
		newWindow: newWindow,

		List:                   spslist.New(),
		SettingsDiscloserMajor: spsdiscloser.New(true),
		SettingsDiscloserMinor: spsdiscloser.New(true),
	}
	return p
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	spsgio.Pages

	Tabs  spstab.Tabs
	Table spstable.Table

	app          spsgio.App
	newWindow    spsgio.NewWindow
	NewWindowBtn widget.Clickable

	List                   spslist.List
	SettingsDiscloserMajor spsdiscloser.Discloser
	SettingsDiscloserMinor spsdiscloser.Discloser
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Name: "New Window",
				Tag:  &p.NewWindowBtn,
			},
			Layout: func(gtx layout.Context, bg, fg color.NRGBA) layout.Dimensions {
				if p.NewWindowBtn.Clicked(gtx) {
					p.RunWindow()
				}
				btn := component.SimpleIconButton(bg, fg, &p.NewWindowBtn, spsicon.ContentAdd)
				return btn.Layout(gtx)
			},
		},
	}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Preferences",
		Icon: spsicon.ActionSettings,
	}
}

func (p *Page) OnEventPre(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) OnEventPost(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) Layout(win spsgio.Window, gtx layout.Context, param any) layout.Dimensions {
	theme := win.GetTheme()
	pref := win.GetPref()
	return p.Tabs.Layout(theme, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		if pref.Pref.TableStyle.Value {
			return p.LayoutTableStyle(win, gtx, param, selected)
		}
		return p.LayoutDefault(win, gtx, param, selected)
	})
}

func (p *Page) LayoutDefault(win spsgio.Window, gtx layout.Context, param any, selected int) layout.Dimensions {
	return p.List.Layout(win.GetTheme(), gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		switch selected {
		case TabIdxSettings:
			return p.LayoutSettings(win, gtx, param)
		case TabIdxPreferences:
			return p.LayoutPref(win, gtx, param)
		default:
			return layout.Dimensions{}
		}
	})
}

func (p *Page) LayoutTableStyle(win spsgio.Window, gtx layout.Context, param any, selected int) layout.Dimensions {
	theme := win.GetTheme()

	var count int
	var cell spstable.Cell
	switch selected {
	case TabIdxPreferences:
		count = PreferencesRowCount
		cell = func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
			return prefCell(p, win, gtx, row, col, labelStyle)
		}
	case TabIdxSettings:
		count = SettingsRowCount
		cell = func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
			return settingsCell(p, win, gtx, row, col, labelStyle)
		}
	default:
		return layout.Dimensions{}
	}

	p.UpdateTableHeaders(win)
	dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
		return tableDimension(win, gtx, axis, index, constraint, minSize, height)
	}
	return p.Table.Layout(theme, gtx, count, dimensioner, cell)
}

func (p *Page) UpdateTableHeaders(win spsgio.Window) {
	if win.GetPref().Settings.ValueInFront.Value {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Value"}, {Text: "Key"}}...))
	} else {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...))
	}
}

func (p *Page) RunWindow() {
	if p.app == nil || p.newWindow == nil {
		return
	}
	win := p.newWindow(p.app, p.Pages.GetWindow())
	if win == nil {
		return
	}
	p.app.RunWindow(win)
}
