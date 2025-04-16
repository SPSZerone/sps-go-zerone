package sample

import (
	"time"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/page/sample/bag"
	spsbag "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
)

const (
	TabIdxGridAndItem = iota
)

const (
	TabNameGridAndItem = "Grid & Item"
)

func New(app *spsgio.Application) *Page {
	bags := spsbag.NewBags()
	bags.AddBag(spsbag.NewBag("Items A"), spsbag.NewBag("Items B"))
	p := &Page{
		Pages:   &app.Pages,
		Tabs:    spstab.NewTabsByNames([]string{TabNameGridAndItem}),
		Bags:    bags,
		BagData: bag.NewTestData(app, bags.GetCount(), 100),
	}
	return p
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	*spsgio.Pages
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

func (p *Page) OnEventPre(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) OnEventPost(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) Layout(app *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	return p.Tabs.Layout(app, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		switch selected {
		case TabIdxGridAndItem:
			return p.Bags.Layout(
				app, gtx, param,
				func(index int) (items []spsitem.Item, itemUpdateTime time.Time) {
					return p.BagData.GetItems(index)
				},
			)
		default:
			return layout.Dimensions{}
		}
	})
}
