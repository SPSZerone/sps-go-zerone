package sample

import (
	"fmt"

	"gioui.org/layout"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsgrid "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/grid"
)

func NewGrid(app *spsgio.Application) Grid {
	g := Grid{
		split: spslayout.Split{
			Flex: layout.Flex{
				Axis: layout.Horizontal,
			},
			Ratio: 0.2,
		},
		Grid: spsgrid.NewGrid(),
	}
	testInit(&g, app)
	return g
}

func testInit(g *Grid, app *spsgio.Application) {
	// test Item
	const count = 100
	g.Items = make([]Item, count)
	for i := 0; i < 100; i++ {
		g.Items[i] = NewItem(fmt.Sprintf("Item %d", i), app)
	}
}

type Grid struct {
	split spslayout.Split
	Grid  spsgrid.Grid

	Items      []Item
	SelectItem int
}

func (g *Grid) Layout(app *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return g.split.Layout(
				gtx,
				func(gtx layout.Context) layout.Dimensions {
					return g.Grid.Layout(app, gtx, len(g.Items), func(gtx layout.Context, index int) layout.Dimensions {
						item := g.GetItem(index)
						highlight := index == g.SelectItem
						dimensions, clicked := item.Item.Layout(app, gtx, highlight)
						if clicked {
							g.SelectItem = index
						}
						return dimensions
					})
				},
				func(gtx layout.Context) layout.Dimensions {
					item := g.GetSelectItem()
					if item == nil {
						return layout.Dimensions{}
					}
					return item.layoutDetail(app, gtx, "Item Detail")
				},
			)
		}),
	)
}

func (g *Grid) GetSelectItem() *Item {
	return g.GetItem(g.SelectItem)
}

func (g *Grid) GetItem(index int) *Item {
	if index < 0 || index >= len(g.Items) {
		return nil
	}
	return &g.Items[index]
}
