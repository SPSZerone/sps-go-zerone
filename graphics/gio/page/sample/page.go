package sample

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
)

const (
	TabIdxGridAndItem = iota
)

const (
	TabNameGridAndItem = "Grid & Item"
)

func New(app *spsgio.Application) *Page {
	p := &Page{
		Pages: &app.Pages,
		Tabs:  spstab.NewTabsByNames(TabNameGridAndItem),
		Grid:  NewGrid(app),
	}
	return p
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	*spsgio.Pages
	Tabs spstab.Tabs

	Grid Grid
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Sample",
		Icon: spsicon.ActionHelp,
	}
}

func (p *Page) OnEventPre(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) OnEventPost(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) Layout(app *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	return p.Tabs.Layout(app, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		switch selected {
		case TabIdxGridAndItem:
			return p.Grid.Layout(app, gtx, param)
		default:
			return layout.Dimensions{}
		}
	})
}
