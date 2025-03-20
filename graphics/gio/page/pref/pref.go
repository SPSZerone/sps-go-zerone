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

func New(app *spsgio.Application) *Page {
	return &Page{
		Pages: &app.Pages,
	}
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	widget.List
	*spsgio.Pages

	decorated      widget.Bool
	nonModalDrawer widget.Bool
	bottomBar      widget.Bool
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
								p.Pages.NavAnim.Appear(gtx.Now)
							} else {
								p.Pages.NavAnim.Disappear(gtx.Now)
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
								p.Pages.ModalNavDrawer.Anchor = component.Bottom
								p.Pages.AppBar.Anchor = component.Bottom
							} else {
								p.Pages.ModalNavDrawer.Anchor = component.Top
								p.Pages.AppBar.Anchor = component.Top
							}
							application.Pref.Settings.BottomBar = p.bottomBar.Value
						}
						return material.Switch(application.Theme, &p.bottomBar, "Use Bottom App Bar").Layout(gtx)
					})
			}),
		)
	})
}
