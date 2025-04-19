package pref

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
	spstable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/table"
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

func New(pages *spswin.Pages) *Page {
	p := &Page{
		Pages: pages,

		Tabs:  spstab.NewTabsByNames([]string{TabNameSettings, TabNamePreferences}),
		Table: spstable.NewTable(),
	}
	return p
}

var _ spswin.Page = (*Page)(nil)

type Page struct {
	*spswin.Pages

	Tabs  spstab.Tabs
	Table spstable.Table
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
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

func (p *Page) OnEventPre(win *spswin.Window, evt event.Event, param any) {

}

func (p *Page) OnEventPost(win *spswin.Window, evt event.Event, param any) {

}

func (p *Page) Layout(win *spswin.Window, gtx layout.Context, param any) layout.Dimensions {
	p.Tabs.Opts.Axis = win.Pref.Settings.TabAxis.GetAxis()
	return p.Tabs.Layout(win.Theme, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		if win.Pref.Pref.TableStyle.Value {
			return p.LayoutTableStyle(win, gtx, param, selected)
		}
		return p.LayoutDefault(win, gtx, param, selected)
	})
}

func (p *Page) LayoutDefault(win *spswin.Window, gtx layout.Context, param any, selected int) layout.Dimensions {
	switch selected {
	case TabIdxSettings:
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, p.LayoutSettings(win, gtx, param)...)
	case TabIdxPreferences:
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, p.LayoutPref(win, gtx, param)...)
	default:
		return layout.Dimensions{}
	}
}

func (p *Page) LayoutTableStyle(win *spswin.Window, gtx layout.Context, param any, selected int) layout.Dimensions {
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
	return p.Table.Layout(win.Theme, gtx, count, dimensioner, cell)
}

func (p *Page) UpdateTableHeaders(win *spswin.Window) {
	if win.Pref.Settings.ValueInFront.Value {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Value"}, {Text: "Key"}}...))
	} else {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...))
	}
}
