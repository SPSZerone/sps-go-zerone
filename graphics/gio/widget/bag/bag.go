package bag

import (
	"time"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

func NewBag(name string) Bag {
	return Bag{
		Name: name,
		Split: spslayout.Split{
			Flex: layout.Flex{
				Axis: layout.Horizontal,
			},
			Ratio: 0.2,
		},

		Grid: NewGrid(),
		DetailList: widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

type Data func(index int) (items []Item, itemUpdateTime time.Time)

type Bag struct {
	Name  string
	Split spslayout.Split

	Grid       Grid
	DetailList widget.List
}

func (b *Bag) Layout(
	app *spsgio.Application, gtx layout.Context, param any,
	itemData ItemData,
) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return b.Split.Layout(
				gtx,
				// item grid
				func(gtx layout.Context) layout.Dimensions {
					return b.Grid.Layout(app, gtx, itemData)
				},
				// item detail
				func(gtx layout.Context) layout.Dimensions {
					b.DetailList.Axis = layout.Vertical
					return material.List(app.Theme, &b.DetailList).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
						if b.Grid.ItemSelected == nil {
							return layout.Dimensions{}
						}
						return b.Grid.ItemSelected.LayoutDetail(app, gtx, "Item Detail")
					})
				},
			)
		}),
	)
}
