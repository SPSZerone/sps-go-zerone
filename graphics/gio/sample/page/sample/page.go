package sample

import (
	"fmt"
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
	TabIdxBags = iota
)

func GetTabName(index int) string {
	switch index {
	case TabIdxBags:
		return "Bags"
	}
	return "Unknown"
}

func New(theme *material.Theme, pages *spswin.Pages) *Page {
	tabs := spstab.NewTabsByNames(
		[]string{GetTabName(TabIdxBags)},
		spstab.OptAxisSetting(&pages.Window.Pref.Settings.TabAxis),
	)
	bags := spsbag.NewBags()
	bags.Tabs.Update(
		spstab.OptAxisSetting(&pages.Window.Pref.Settings.TabAxis),
	)
	for i := 0; i < 32; i++ {
		bags.AddBag(
			spsbag.NewBag(fmt.Sprintf("Bag %v", i)),
		)
	}
	p := &Page{
		Pages:   pages,
		Tabs:    tabs,
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
	BagData bag.TestData
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
	return p.Tabs.Layout(win.Theme, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		switch selected {
		case TabIdxBags:
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
