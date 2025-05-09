package about

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spscopy "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/copy"
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

	CopyName     spscopy.Copy
	CopyAuthor   spscopy.Copy
	CopyLicense  spscopy.Copy
	CopyHomePage spscopy.Copy
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
				return p.LayoutInfo(theme, gtx, "Name", Name, &p.CopyName)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "Author", Author, &p.CopyAuthor)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "License", License, &p.CopyLicense)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "HomePage", HomePage, &p.CopyHomePage)
			}),
			//children...,
		)
	})
}

func (p *Page) LayoutInfo(theme *material.Theme, gtx layout.Context, name, value string, copy *spscopy.Copy) layout.Dimensions {
	return copy.LayoutCopyFlexChild(
		theme,
		gtx,
		func() string {
			return value
		},
		layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:   unit.Dp(8),
				Right:  unit.Dp(8),
				Top:    unit.Dp(4),
				Bottom: unit.Dp(4),
			}.Layout(gtx, material.H6(theme, name).Layout)
		}),
		layout.Flexed(0.8, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:   unit.Dp(8),
				Right:  unit.Dp(8),
				Top:    unit.Dp(4),
				Bottom: unit.Dp(4),
			}.Layout(gtx, material.Body1(theme, value).Layout)
		}),
	)
}
