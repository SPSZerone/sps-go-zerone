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

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	widget.List
	*spsgio.Pages

	decorated      widget.Bool
	nonModalDrawer widget.Bool
	bottomBar      widget.Bool
}

func New(pages *spsgio.Pages) *Page {
	return &Page{
		Pages: pages,
	}
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

func (p *Page) Layout(application *spsgio.Application, gtx layout.Context, w *app.Window, th *material.Theme) layout.Dimensions {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.DefaultInset.Layout(gtx, material.Body1(th, `Settings...`).Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(th, "  * Decorated").Layout,
					func(gtx layout.Context) layout.Dimensions {
						if p.decorated.Update(gtx) {
							application.Pref.Settings.Decorated = p.decorated.Value
							w.Option(app.Decorated(p.decorated.Value))
						}
						return material.Switch(th, &p.decorated, "Use decorated").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(th, "  * Use non-modal drawer").Layout,
					func(gtx layout.Context) layout.Dimensions {
						if p.nonModalDrawer.Update(gtx) {
							application.Pref.Settings.NonModalDrawer = p.nonModalDrawer.Value
							if p.nonModalDrawer.Value {
								p.Pages.NavAnim.Appear(gtx.Now)
							} else {
								p.Pages.NavAnim.Disappear(gtx.Now)
							}
						}
						return material.Switch(th, &p.nonModalDrawer, "Use Non-Modal Navigation Drawer").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.FlexInset{}.LayoutABWidget(gtx,
					material.Body1(th, "  * Bottom App Bar").Layout,
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
						return material.Switch(th, &p.bottomBar, "Use Bottom App Bar").Layout(gtx)
					})
			}),
		)
	})
}
