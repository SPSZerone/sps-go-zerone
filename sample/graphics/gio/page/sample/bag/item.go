package bag

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsbag "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
	spseditor "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/editor"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func newItem(data Data) *Item {
	return &Item{
		UI:   NewUI(),
		Data: data,
	}
}

var _ spsbag.Item = (*Item)(nil)

type Item struct {
	spsitem.Item
	UI
	Data

	Suggests []*spseditor.Suggest
}

func (i *Item) Init(item spsitem.Item, theme *material.Theme) {
	i.Item = item
	i.InitMenu(theme)
}

func (i *Item) Layout(
	theme *material.Theme, gtx layout.Context,
	highlight bool,
	onClick func(),
) layout.Dimensions {
	return i.Item.Layout(theme, gtx, highlight, i.LayoutContent, onClick)
}
