package pref

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

var _ spsgio.Tab = (*Tab)(nil)

type Tab struct {
	widget.List
	*spsgio.Tabs

	decorated      widget.Bool
	nonModalDrawer widget.Bool
	bottomBar      widget.Bool
}

func New(tabs *spsgio.Tabs) *Tab {
	return &Tab{
		Tabs: tabs,
	}
}

func (p *Tab) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Tab) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Tab) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Preferences...",
		Icon: spsicon.ActionSettings,
	}
}

func (p *Tab) Layout(application *spsgio.Application, gtx layout.Context) layout.Dimensions {
	p.List.Axis = layout.Vertical
	return material.List(application.Theme, &p.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.DefaultInset.Layout(gtx, material.Body1(application.Theme, `Settings...`).Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(application.Theme, "  * Decorated").Layout,
					func(gtx layout.Context) layout.Dimensions {
						if p.decorated.Update(gtx) {
							application.Pref.Settings.Decorated = p.decorated.Value
							application.Window.Option(app.Decorated(p.decorated.Value))
						}
						return material.Switch(application.Theme, &p.decorated, "Use decorated").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(application.Theme, "  * Use non-modal drawer").Layout,
					func(gtx layout.Context) layout.Dimensions {
						if p.nonModalDrawer.Update(gtx) {
							application.Pref.Settings.NonModalDrawer = p.nonModalDrawer.Value
							if p.nonModalDrawer.Value {
								p.Tabs.NavAnim.Appear(gtx.Now)
							} else {
								p.Tabs.NavAnim.Disappear(gtx.Now)
							}
						}
						return material.Switch(application.Theme, &p.nonModalDrawer, "Use Non-Modal Navigation Drawer").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(application.Theme, "  * Bottom App Bar").Layout,
					func(gtx layout.Context) layout.Dimensions {
						if p.bottomBar.Update(gtx) {
							if p.bottomBar.Value {
								p.Tabs.ModalNavDrawer.Anchor = component.Bottom
								p.Tabs.AppBar.Anchor = component.Bottom
							} else {
								p.Tabs.ModalNavDrawer.Anchor = component.Top
								p.Tabs.AppBar.Anchor = component.Top
							}
							application.Pref.Settings.BottomBar = p.bottomBar.Value
						}
						return material.Switch(application.Theme, &p.bottomBar, "Use Bottom App Bar").Layout(gtx)
					})
			}),
		)
	})
}
