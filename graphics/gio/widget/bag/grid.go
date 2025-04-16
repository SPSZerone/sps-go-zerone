package bag

import (
	"time"

	"gioui.org/layout"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsgrid "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/grid"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func NewGrid() Grid {
	g := spsgrid.NewGrid()
	g.Num = 4
	return Grid{
		Grid: g,
	}
}

type Grid struct {
	Grid spsgrid.Grid

	Items          []spsitem.Item
	ItemUpdateTime time.Time
	ItemSelected   *spsitem.Item
}

func (g *Grid) Layout(app *spsgio.Application, gtx layout.Context, itemData ItemData) layout.Dimensions {
	g.UpdateItems(itemData())

	return g.Grid.Layout(app, gtx, g.GetItemCount(), func(gtx layout.Context, index int) layout.Dimensions {
		curItem := g.GetItem(index)
		highlight := curItem == g.ItemSelected
		dimensions, clicked := curItem.Layout(app, gtx, highlight)
		if clicked {
			g.ItemSelected = curItem
		}
		return dimensions
	})
}

func (g *Grid) UpdateItems(items []spsitem.Item, itemUpdateTime time.Time) {
	if itemUpdateTime == g.ItemUpdateTime {
		return
	}

	g.Items = items
	g.ItemUpdateTime = itemUpdateTime

	if len(g.Items) > 0 {
		g.ItemSelected = &g.Items[0]
	}
}

func (g *Grid) GetItem(index int) *spsitem.Item {
	if index < 0 || index >= len(g.Items) {
		return nil
	}
	return &g.Items[index]
}

func (g *Grid) GetItemCount() int {
	return len(g.Items)
}
