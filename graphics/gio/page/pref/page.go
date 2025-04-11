package pref

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/tab"
	spstable "github.com/SPSZerone/sps-go-zerone/graphics/gio/table"
)

const (
	SettingsColIdxKey = iota
	SettingsColIdxValue
)

const (
	SettingsRowIdxNonModalDrawer = iota
	SettingsRowIdxBottomBar
	SettingsRowIdxTableStyle
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
	}
	p.decorated.Widget.Value = app.Pref.Settings.Decorated
	p.nonModalDrawer.Widget.Value = app.Pref.Settings.NonModalDrawer
	p.bottomBar.Widget.Value = app.Pref.Settings.BottomBar
	p.prefTableStyle.Widget.Value = app.Pref.Settings.PrefTableStyle
	return p
}

const (
	TabIdxSettings = iota
	TabIdxDecorated
)

const (
	TabNameSettings  = "Settings"
	TabNameDecorated = "Decorated"
)

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
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Preferences...",
		Icon: spsicon.ActionSettings,
	}
}

func (p *Page) OnEventPre(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) OnEventPost(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) Layout(application *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	p.List.Axis = layout.Vertical
	return material.List(application.Theme, &p.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return p.Tabs.Layout(application, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
			if selected == TabIdxDecorated {
				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Vertical,
				}.Layout(gtx, p.PrefDecorated(application, gtx, param)...)
			}

			if application.Pref.Settings.PrefTableStyle {
				p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...))
				dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
					return settingsDimension(gtx, axis, index, constraint, minSize, height)
				}
				cell := func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
					return settingsCell(p, application, gtx, row, col, labelStyle)
				}
				return p.Table.Layout(application.Theme, gtx, 3, dimensioner, cell)
			}

			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Vertical,
			}.Layout(gtx, p.PrefSettings(application, gtx, param)...)
		})
	})
}

func (p *Page) PrefDecorated(application *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.decorated.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.Decorated(application, gtx)
				})
		}),
	}
}

func (p *Page) PrefSettings(application *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.nonModalDrawer.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.NonModalDrawer(application, gtx)
				})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.bottomBar.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.BottomBar(application, gtx)
				})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.prefTableStyle.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.PrefTableStyle(application, gtx)
				})
		}),
	}
}

func (p *Page) Decorated(application *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.decorated.Widget.Update(gtx) {
		application.Pref.Settings.Decorated = p.decorated.Widget.Value
		application.Window.Option(app.Decorated(p.decorated.Widget.Value))
	}
	return material.Switch(application.Theme, &p.decorated.Widget, p.bottomBar.Desc).Layout(gtx)
}

func (p *Page) NonModalDrawer(application *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.nonModalDrawer.Widget.Update(gtx) {
		application.Pref.Settings.NonModalDrawer = p.nonModalDrawer.Widget.Value
		if p.nonModalDrawer.Widget.Value {
			p.Pages.NavAnim.Appear(gtx.Now)
		} else {
			p.Pages.NavAnim.Disappear(gtx.Now)
		}
	}
	return material.Switch(application.Theme, &p.nonModalDrawer.Widget, p.bottomBar.Desc).Layout(gtx)
}

func (p *Page) BottomBar(application *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.bottomBar.Widget.Update(gtx) {
		if p.bottomBar.Widget.Value {
			p.Pages.ModalNavDrawer.Anchor = component.Bottom
			p.Pages.AppBar.Anchor = component.Bottom
		} else {
			p.Pages.ModalNavDrawer.Anchor = component.Top
			p.Pages.AppBar.Anchor = component.Top
		}
		application.Pref.Settings.BottomBar = p.bottomBar.Widget.Value
	}
	return material.Switch(application.Theme, &p.bottomBar.Widget, p.bottomBar.Desc).Layout(gtx)
}

func (p *Page) PrefTableStyle(application *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.prefTableStyle.Widget.Update(gtx) {
		application.Pref.Settings.PrefTableStyle = p.prefTableStyle.Widget.Value
	}
	return material.Switch(application.Theme, &p.prefTableStyle.Widget, p.prefTableStyle.Desc).Layout(gtx)
}

func settingsDimension(gtx layout.Context, axis layout.Axis, index, constraint, minSize, height int) int {
	switch axis {
	case layout.Horizontal:
		var widthUnit int
		switch index {
		case SettingsColIdxKey:
			widthUnit = gtx.Dp(unit.Dp(220))
		case SettingsColIdxValue:
			widthUnit = gtx.Dp(unit.Dp(64))
		}
		return widthUnit
	default:
		return height
	}
}

func settingsCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	switch row {
	case SettingsRowIdxNonModalDrawer:
		switch col {
		case SettingsColIdxValue:
			return p.NonModalDrawer(app, gtx)
		default:
			labelStyle.Text = p.nonModalDrawer.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxBottomBar:
		switch col {
		case SettingsColIdxValue:
			return p.BottomBar(app, gtx)
		default:
			labelStyle.Text = p.bottomBar.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxTableStyle:
		switch col {
		case SettingsColIdxValue:
			return p.PrefTableStyle(app, gtx)
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
