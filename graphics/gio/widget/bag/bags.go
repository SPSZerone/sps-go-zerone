package bag

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

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

func (b *Bags) Layout(
	theme *material.Theme, gtx layout.Context, param any,
	itemDataSource ItemDataSource,
) layout.Dimensions {
	return b.Tabs.Layout(theme, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
		bag := b.GetBag(selected)
		if bag == nil {
			return layout.Dimensions{}
		}
		bag.UpdateItems(itemDataSource(selected))
		return bag.Layout(theme, gtx, param)
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

func (b *Bags) GetCount() int {
	return len(b.Bags)
}
