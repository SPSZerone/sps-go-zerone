package pref

import (
	"gioui.org/app"
	"gioui.org/io/event"
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

func (t *Tab) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (t *Tab) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (t *Tab) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Preferences...",
		Icon: spsicon.ActionSettings,
	}
}

func (t *Tab) OnEventPre(app *spsgio.Application, evt event.Event) {

}

func (t *Tab) OnEventPost(app *spsgio.Application, evt event.Event) {

}

func (t *Tab) Layout(application *spsgio.Application, gtx layout.Context) layout.Dimensions {
	t.List.Axis = layout.Vertical
	return material.List(application.Theme, &t.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
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
						if t.decorated.Update(gtx) {
							application.Pref.Settings.Decorated = t.decorated.Value
							application.Window.Option(app.Decorated(t.decorated.Value))
						}
						return material.Switch(application.Theme, &t.decorated, "Use decorated").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(application.Theme, "  * Use non-modal drawer").Layout,
					func(gtx layout.Context) layout.Dimensions {
						if t.nonModalDrawer.Update(gtx) {
							application.Pref.Settings.NonModalDrawer = t.nonModalDrawer.Value
							if t.nonModalDrawer.Value {
								t.Tabs.NavAnim.Appear(gtx.Now)
							} else {
								t.Tabs.NavAnim.Disappear(gtx.Now)
							}
						}
						return material.Switch(application.Theme, &t.nonModalDrawer, "Use Non-Modal Navigation Drawer").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(application.Theme, "  * Bottom App Bar").Layout,
					func(gtx layout.Context) layout.Dimensions {
						if t.bottomBar.Update(gtx) {
							if t.bottomBar.Value {
								t.Tabs.ModalNavDrawer.Anchor = component.Bottom
								t.Tabs.AppBar.Anchor = component.Bottom
							} else {
								t.Tabs.ModalNavDrawer.Anchor = component.Top
								t.Tabs.AppBar.Anchor = component.Top
							}
							application.Pref.Settings.BottomBar = t.bottomBar.Value
						}
						return material.Switch(application.Theme, &t.bottomBar, "Use Bottom App Bar").Layout(gtx)
					})
			}),
		)
	})
}
