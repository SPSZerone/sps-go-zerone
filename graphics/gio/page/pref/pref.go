package pref

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/page"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Page struct {
	widget.List
	*page.Pages

	decorated      widget.Bool
	nonModalDrawer widget.Bool
	bottomBar      widget.Bool
}

func New(pages *page.Pages) *Page {
	return &Page{
		Pages: pages,
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Preferences...",
		Icon: icon.ActionSettings,
	}
}

func (p *Page) Layout(gtx C, w *app.Window, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.DefaultInset.Layout(gtx, material.Body1(th, `Settings...`).Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.Setting{}.Layout(gtx,
					material.Body1(th, "  * Decorated").Layout,
					func(gtx C) D {
						if p.decorated.Update(gtx) {
							p.Pages.Pref.Settings.Decorated = p.decorated.Value
							w.Option(app.Decorated(p.decorated.Value))
						}
						return material.Switch(th, &p.decorated, "Use decorated").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return spslayout.Setting{}.Layout(gtx,
					material.Body1(th, "  * Use non-modal drawer").Layout,
					func(gtx C) D {
						if p.nonModalDrawer.Update(gtx) {
							p.Pages.Pref.Settings.NonModalDrawer = p.nonModalDrawer.Value
							if p.nonModalDrawer.Value {
								p.Pages.NavAnim.Appear(gtx.Now)
							} else {
								p.Pages.NavAnim.Disappear(gtx.Now)
							}
						}
						return material.Switch(th, &p.nonModalDrawer, "Use Non-Modal Navigation Drawer").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx C) D {
				return spslayout.Setting{}.Layout(gtx,
					material.Body1(th, "  * Bottom App Bar").Layout,
					func(gtx C) D {
						if p.bottomBar.Update(gtx) {
							if p.bottomBar.Value {
								p.Pages.ModalNavDrawer.Anchor = component.Bottom
								p.Pages.AppBar.Anchor = component.Bottom
							} else {
								p.Pages.ModalNavDrawer.Anchor = component.Top
								p.Pages.AppBar.Anchor = component.Top
							}
							p.Pages.Pref.Settings.BottomBar = p.bottomBar.Value
						}
						return material.Switch(th, &p.bottomBar, "Use Bottom App Bar").Layout(gtx)
					})
			}),
		)
	})
}
