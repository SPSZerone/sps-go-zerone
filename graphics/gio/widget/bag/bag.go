package bag

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spslist "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/list"
)

func New(id any, name string) Bag {
	return Bag{
		Name: name,
		Split: spslayout.Split{
			Flex: layout.Flex{
				Axis: layout.Horizontal,
			},
			Ratio: 0,
		},

		Grid:       NewGrid(),
		DetailList: spslist.New(),
	}
}

type Bag struct {
	Id   any
	Name string

	Split spslayout.Split

	Grid       Grid
	DetailList spslist.List
}

func (b *Bag) Layout(
	theme *material.Theme, gtx layout.Context,
	bagIndex int,
	itemLength int, listGridItem ListGridItem,
	itemDetail ItemDetail,
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
					return b.Grid.Layout(theme, gtx, itemLength, listGridItem)
				},
				// item detail
				func(gtx layout.Context) layout.Dimensions {
					b.DetailList.Axis = layout.Vertical
					return b.DetailList.Layout(theme, gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
						return itemDetail(gtx, bagIndex, b, b.Grid.CurSelectedIndex)
					})
				},
			)
		}),
	)
}
