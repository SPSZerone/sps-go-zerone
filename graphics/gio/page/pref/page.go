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
	TableColIdxKey = iota
	TableColIdxValue
)

const (
	SettingsRowIdxNonModalDrawer = iota
	SettingsRowIdxBottomBar
	SettingsRowIdxTableStyle
	SettingsRowIdxSwitchInFront
	SettingsRowCount
)

const (
	DecoratedRowIdxDecorated = iota
	DecoratedRowCount
)

const (
	RatioSwitch = 0.2
	RatioName   = 0.3
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
		switchInFront:  NewSwitchInFront(),
	}
	p.decorated.Widget.Value = app.Pref.Settings.Decorated
	p.nonModalDrawer.Widget.Value = app.Pref.Settings.NonModalDrawer
	p.bottomBar.Widget.Value = app.Pref.Settings.BottomBar
	p.prefTableStyle.Widget.Value = app.Pref.Settings.PrefTableStyle
	p.switchInFront.Widget.Value = app.Pref.Settings.SwitchInFront
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
	switchInFront  SettingBool
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

func (p *Page) Layout(application *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	p.List.Axis = layout.Vertical
	return material.List(application.Theme, &p.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return p.Tabs.Layout(application, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
			switch selected {
			case TabIdxDecorated:
				if application.Pref.Settings.PrefTableStyle {
					p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...))
					dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
						return dimension(gtx, axis, index, constraint, minSize, height)
					}
					cell := func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
						return decoratedCell(p, application, gtx, row, col, labelStyle)
					}
					return p.Table.Layout(application.Theme, gtx, DecoratedRowCount, dimensioner, cell)
				}

				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Vertical,
				}.Layout(gtx, p.PrefDecorated(application, gtx, param)...)

			case TabIdxSettings:
				if application.Pref.Settings.PrefTableStyle {
					p.Table.Update(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...))
					dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
						return dimension(gtx, axis, index, constraint, minSize, height)
					}
					cell := func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
						return settingsCell(p, application, gtx, row, col, labelStyle)
					}
					return p.Table.Layout(application.Theme, gtx, SettingsRowCount, dimensioner, cell)
				}

				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Vertical,
				}.Layout(gtx, p.PrefSettings(application, gtx, param)...)
			default:
				return layout.Dimensions{}
			}
		})
	})
}

func (p *Page) PrefDecorated(application *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	if application.Pref.Settings.SwitchInFront {
		return []layout.FlexChild{
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{
					Ratio: RatioSwitch,
				}.LayoutABWidget(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return p.Decorated(application, gtx)
					},
					material.Body1(application.Theme, p.decorated.Name).Layout,
				)
			}),
		}
	}

	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{
				Ratio: RatioName,
			}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.decorated.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.Decorated(application, gtx)
				},
			)
		}),
	}
}

func (p *Page) PrefSettings(application *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	if application.Pref.Settings.SwitchInFront {
		return []layout.FlexChild{
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{
					Ratio: RatioSwitch,
				}.LayoutABWidget(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return p.NonModalDrawer(application, gtx)
					},
					material.Body1(application.Theme, p.nonModalDrawer.Name).Layout,
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{
					Ratio: RatioSwitch,
				}.LayoutABWidget(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return p.BottomBar(application, gtx)
					},
					material.Body1(application.Theme, p.bottomBar.Name).Layout,
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{
					Ratio: RatioSwitch,
				}.LayoutABWidget(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return p.PrefTableStyle(application, gtx)
					},
					material.Body1(application.Theme, p.prefTableStyle.Name).Layout,
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{
					Ratio: RatioSwitch,
				}.LayoutABWidget(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return p.PrefSwitchInFront(application, gtx)
					},
					material.Body1(application.Theme, p.switchInFront.Name).Layout,
				)
			}),
		}
	}

	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{
				Ratio: RatioName,
			}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.nonModalDrawer.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.NonModalDrawer(application, gtx)
				},
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{
				Ratio: RatioName,
			}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.bottomBar.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.BottomBar(application, gtx)
				},
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{
				Ratio: RatioName,
			}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.prefTableStyle.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.PrefTableStyle(application, gtx)
				},
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spslayout.FlexInset{
				Ratio: RatioName,
			}.LayoutABWidget(gtx,
				material.Body1(application.Theme, p.switchInFront.Name).Layout,
				func(gtx layout.Context) layout.Dimensions {
					return p.PrefSwitchInFront(application, gtx)
				},
			)
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

func (p *Page) PrefSwitchInFront(application *spsgio.Application, gtx layout.Context) layout.Dimensions {
	if p.switchInFront.Widget.Update(gtx) {
		application.Pref.Settings.SwitchInFront = p.switchInFront.Widget.Value
	}
	return material.Switch(application.Theme, &p.switchInFront.Widget, p.switchInFront.Desc).Layout(gtx)
}

func dimension(gtx layout.Context, axis layout.Axis, index, constraint, minSize, height int) int {
	switch axis {
	case layout.Horizontal:
		var widthUnit int
		switch index {
		case TableColIdxKey:
			widthUnit = gtx.Dp(unit.Dp(220))
		case TableColIdxValue:
			widthUnit = gtx.Dp(unit.Dp(64))
		}
		return widthUnit
	default:
		return height
	}
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

func settingsCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	switch row {
	case SettingsRowIdxNonModalDrawer:
		switch col {
		case TableColIdxValue:
			return p.NonModalDrawer(app, gtx)
		default:
			labelStyle.Text = p.nonModalDrawer.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxBottomBar:
		switch col {
		case TableColIdxValue:
			return p.BottomBar(app, gtx)
		default:
			labelStyle.Text = p.bottomBar.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxTableStyle:
		switch col {
		case TableColIdxValue:
			return p.PrefTableStyle(app, gtx)
		default:
			labelStyle.Text = p.prefTableStyle.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxSwitchInFront:
		switch col {
		case TableColIdxValue:
			return p.PrefSwitchInFront(app, gtx)
		default:
			labelStyle.Text = p.switchInFront.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}
