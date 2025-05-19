package bag

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsslice "github.com/SPSZerone/sps-go-zerone/generic/slice"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
)

func NewBags(opts ...spstab.Option) Bags {
	tabs := spstab.NewTabs(
		spstab.OptColorfulBG(true),
	)
	tabs.Update(opts...)
	return Bags{
		Tabs: tabs,
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
		bag := b.GetBagByIndex(selected)
		if bag == nil {
			return layout.Dimensions{}
		}
		bag.UpdateItems(itemDataSource(selected))
		return bag.Layout(theme, gtx, param)
	})
}

func (b *Bags) AddBag(bags ...Bag) {
	for _, bag := range bags {
		b.addBag(bag)
	}
}

func (b *Bags) addBag(bag Bag) {
	b.Bags = append(b.Bags, bag)

	b.Tabs.AddTab(spstab.New(bag.Name, b.GetLastBag()))
}

func (b *Bags) GetFirstBag() *Bag {
	return b.GetBagByIndex(0)
}

func (b *Bags) GetLastBag() *Bag {
	return b.GetBagByIndex(b.GetCount() - 1)
}

func (b *Bags) GetBagByIndex(index int) *Bag {
	if index < 0 || index >= len(b.Bags) {
		return nil
	}
	return &b.Bags[index]
}

func (b *Bags) GetBag(cb func(idx int, bag *Bag) (ok bool)) *Bag {
	for i := 0; i < len(b.Bags); i++ {
		if cb(i, &b.Bags[i]) {
			return &b.Bags[i]
		}
	}
	return nil
}

func (b *Bags) DelBag(cb func(idx int, bag *Bag) (del bool)) {
	if len(b.Bags) == 0 {
		return
	}

	delIndexes := make([]int, 0, len(b.Bags))
	for i, bag := range b.Bags {
		if cb(i, &bag) {
			delIndexes = append(delIndexes, i)
		}
	}

	for i := len(delIndexes) - 1; i >= 0; i-- {
		idx := delIndexes[i]
		b.Bags = append(b.Bags[:idx], b.Bags[idx+1:]...)
	}
}

func (b *Bags) DelBagByIndex(index int) {
	b.Bags = spsslice.RemoveByKeepOrder(b.Bags, index)
	b.Tabs.DelTabByIndex(index)
}

func (b *Bags) GetCount() int {
	return len(b.Bags)
}
