package pref

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
	spstable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/table"
)

func New(app *spsgio.Application) *Page {
	p := &Page{
		Pages: &app.Pages,

		Tabs:  spstab.NewTabsByNames(TabNameSettings, TabNameDecorated),
		Table: spstable.NewTable(),

		decorated:      NewDecorated(),
		nonModalDrawer: NewNonModalDrawer(),
		bottomBar:      NewBottomBar(),
		prefTableStyle: NewPrefTableStyle(),
		valueInFront:   NewValueInFront(),
	}
	p.decorated.Widget.Value = app.Pref.Settings.Decorated
	p.nonModalDrawer.Widget.Value = app.Pref.Settings.NonModalDrawer
	p.bottomBar.Widget.Value = app.Pref.Settings.BottomBar
	p.prefTableStyle.Widget.Value = app.Pref.Settings.PrefTableStyle
	p.valueInFront.Widget.Value = app.Pref.Settings.ValueInFront
	return p
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	widget.List
	*spsgio.Pages

	Tabs  spstab.Tabs
	Table spstable.Table

	decorated      SettingBool
	nonModalDrawer SettingBool
	bottomBar      SettingBool
	prefTableStyle SettingBool
	valueInFront   SettingBool
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
	p.List.Axis = layout.Vertical
	return material.List(app.Theme, &p.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return p.Tabs.Layout(app, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
			switch selected {
			case TabIdxDecorated:
				if app.Pref.Settings.PrefTableStyle {
					p.UpdateTableHeaders(app)
					dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
						return dimension(app, gtx, axis, index, constraint, minSize, height)
					}
					cell := func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
						return decoratedCell(p, app, gtx, row, col, labelStyle)
					}
					return p.Table.Layout(app, gtx, DecoratedRowCount, dimensioner, cell)
				}

				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Vertical,
				}.Layout(gtx, p.PrefDecorated(app, gtx, param)...)

			case TabIdxSettings:
				if app.Pref.Settings.PrefTableStyle {
					p.UpdateTableHeaders(app)
					dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
						return dimension(app, gtx, axis, index, constraint, minSize, height)
					}
					cell := func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
						return settingsCell(p, app, gtx, row, col, labelStyle)
					}
					return p.Table.Layout(app, gtx, SettingsRowCount, dimensioner, cell)
				}

				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Vertical,
				}.Layout(gtx, p.PrefSettings(app, gtx, param)...)
			default:
				return layout.Dimensions{}
			}
		})
	})
}

func (p *Page) UpdateTableHeaders(app *spsgio.Application) {
	if app.Pref.Settings.ValueInFront {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Value"}, {Text: "Key"}}...))
	} else {
		p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...))
	}
}
