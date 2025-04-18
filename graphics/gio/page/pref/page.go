package pref

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
	spstable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/table"
)

func New(pages *spsgio.Pages) *Page {
	p := &Page{
		Pages: pages,

		Tabs:  spstab.NewTabsByNames([]string{TabNameSettings, TabNamePreferences}),
		Table: spstable.NewTable(),
	}
	return p
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	*spsgio.Pages

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

func (p *Page) OnEventPre(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) OnEventPost(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) Layout(app *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	p.Tabs.Opts.Axis = app.Pref.Settings.TabAxis.GetAxis()
	return p.Tabs.Layout(app.Theme, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		if app.Pref.Pref.TableStyle.Value {
			return p.LayoutTableStyle(app, gtx, param, selected)
		}
		return p.LayoutDefault(app, gtx, param, selected)
	})
}

func (p *Page) LayoutDefault(app *spsgio.Application, gtx layout.Context, param any, selected int) layout.Dimensions {
	switch selected {
	case TabIdxSettings:
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, p.LayoutSettings(app, gtx, param)...)
	case TabIdxPreferences:
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, p.LayoutPref(app, gtx, param)...)
	default:
		return layout.Dimensions{}
	}
}

func (p *Page) LayoutTableStyle(app *spsgio.Application, gtx layout.Context, param any, selected int) layout.Dimensions {
	var count int
	var cell spstable.Cell
	switch selected {
	case TabIdxPreferences:
		count = PreferencesRowCount
		cell = func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
			return prefCell(p, app, gtx, row, col, labelStyle)
		}
	case TabIdxSettings:
		count = SettingsRowCount
		cell = func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
			return settingsCell(p, app, gtx, row, col, labelStyle)
		}
	default:
		return layout.Dimensions{}
	}

	p.UpdateTableHeaders(app)
	dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
		return tableDimension(app, gtx, axis, index, constraint, minSize, height)
	}
	return p.Table.Layout(app.Theme, gtx, count, dimensioner, cell)
}

func (p *Page) UpdateTableHeaders(app *spsgio.Application) {
	if app.Pref.Settings.ValueInFront.Value {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Value"}, {Text: "Key"}}...))
	} else {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...))
	}
}
