package sample

import (
	"fmt"
	"time"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spsbag "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"

	"github.com/SPSZerone/sps-go-zerone/sample/graphics/gio/page/sample/bag"
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

func New(theme *material.Theme, pages spsgio.Pages) *Page {
	pref := pages.GetWindow().GetPref()
	tabs := spstab.NewTabsByNames(
		[]string{GetTabName(TabIdxBags)},
		spstab.TabsOptAxisSetting(&pref.Settings.TabAxis),
		spstab.TabsOptColorfulBG(true),
		spstab.TabsOptCloseMode(spstab.CloseModeNone),
	)
	bags := spsbag.NewBags(
		spstab.TabsOptWidthWhenVertical(128),
		spstab.TabsOptCloseMode(spstab.CloseModeMenu),
	)
	bags.Tabs.Update(
		spstab.TabsOptAxisSetting(&pref.Settings.TabAxis),
	)
	for i := 0; i < 32; i++ {
		bags.AddBag(
			spsbag.New(i, fmt.Sprintf("Bag %v", i)),
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

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	spsgio.Pages
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

func (p *Page) OnEventPre(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) OnEventPost(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) Layout(win spsgio.Window, gtx layout.Context, param any) layout.Dimensions {
	theme := win.GetTheme()
	return p.Tabs.Layout(theme, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		switch selected {
		case TabIdxBags:
			if p.Bags.Tabs.Count() > 0 {
				p.Bags.Tabs.Tabs[0].CloseMode = spstab.CloseModeNormal
			}
			return p.Bags.Layout(
				theme, gtx, param,
				func(index int) (items []spsbag.Item, itemUpdateTime time.Time) {
					return p.BagData.GetItems(index)
				},
			)
		default:
			return layout.Dimensions{}
		}
	})
}
