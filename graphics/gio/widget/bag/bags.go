package bag

import (
	"time"

	"gioui.org/layout"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
)

func NewBags() Bags {
	return Bags{
		Tabs: spstab.NewTabs(),
	}
}

type Bags struct {
	Tabs spstab.Tabs
	Bags []Bag
}

func (b *Bags) Layout(app *spsgio.Application, gtx layout.Context, param any, bagData Data) layout.Dimensions {
	return b.Tabs.Layout(app, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		bag := b.GetBag(selected)
		if bag == nil {
			return layout.Dimensions{}
		}
		return bag.Layout(app, gtx, param, func() (items []Item, itemUpdateTime time.Time) {
			return bagData(selected)
		})
	})
}

func (b *Bags) AddBag(bags ...Bag) {
	for _, bag := range bags {
		b.Bags = append(b.Bags, bag)
		tab := spstab.Tab{
			Name: bag.Name,
			Data: b.GetBag(len(b.Bags) - 1),
		}
		b.Tabs.AddTab(tab)
	}
}

func (b *Bags) GetBag(index int) *Bag {
	if index < 0 || index >= len(b.Bags) {
		return nil
	}
	return &b.Bags[index]
}
