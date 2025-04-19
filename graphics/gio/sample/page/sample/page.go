package sample

import (
	"time"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spsbag "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/sample/page/sample/bag"
)

const (
	TabIdxGridAndItem = iota
)

const (
	TabNameGridAndItem = "Grid & Item"
)

func New(theme *material.Theme, pages *spswin.Pages) *Page {
	bags := spsbag.NewBags()
	bags.AddBag(spsbag.NewBag("Items A"), spsbag.NewBag("Items B"))
	p := &Page{
		Pages:   pages,
		Tabs:    spstab.NewTabsByNames([]string{TabNameGridAndItem}),
		Bags:    bags,
		BagData: bag.NewTestData(theme, bags.GetCount(), 100),
	}
	return p
}

var _ spswin.Page = (*Page)(nil)

type Page struct {
	*spswin.Pages
	Tabs spstab.Tabs

	Bags    spsbag.Bags
	BagData bag.Data
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

func (p *Page) OnEventPre(win *spswin.Window, evt event.Event, param any) {

}

func (p *Page) OnEventPost(win *spswin.Window, evt event.Event, param any) {

}

func (p *Page) Layout(win *spswin.Window, gtx layout.Context, param any) layout.Dimensions {
	p.Tabs.Opts.Axis = win.Pref.Settings.TabAxis.GetAxis()
	return p.Tabs.Layout(win.Theme, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		switch selected {
		case TabIdxGridAndItem:
			p.Bags.Tabs.Opts.Axis = win.Pref.Settings.TabAxis.GetAxis()
			return p.Bags.Layout(
				win.Theme, gtx, param,
				func(index int) (items []spsitem.Item, itemUpdateTime time.Time) {
					return p.BagData.GetItems(index)
				},
			)
		default:
			return layout.Dimensions{}
		}
	})
}
