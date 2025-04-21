package about

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spslist "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/list"
)

const (
	Name     = "SPS Gio Framework base on Gio"
	Author   = "SPSZerone"
	License  = "GPLv3"
	HomePage = "https://github.com/SPSZerone"
)

func New(pages spsgio.Pages) *Page {
	return &Page{
		Pages: pages,
	}
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	spslist.List
	spsgio.Pages
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "About",
		Icon: spsicon.ActionHelp,
	}
}

func (p *Page) OnEventPre(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) OnEventPost(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) Layout(win spsgio.Window, gtx layout.Context, param any) layout.Dimensions {
	p.List.Axis = layout.Vertical
	theme := win.GetTheme()
	return p.List.Layout(theme, gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "Name", Name)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "Author", Author)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "License", License)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "HomePage", HomePage)
			}),
		)
	})
}

func (p *Page) LayoutInfo(theme *material.Theme, gtx layout.Context, name, value string) layout.Dimensions {
	return spslayout.FlexInset{
		Ratio: 0.2,
	}.LayoutABWidget(
		gtx,
		material.H6(theme, name).Layout,
		material.Body1(theme, value).Layout,
	)
}
