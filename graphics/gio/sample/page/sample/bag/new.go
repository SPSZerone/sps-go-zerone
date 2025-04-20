package bag

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func NewItem(data Data, theme *material.Theme) Item {
	item := newItem(data)
	item.Init(
		NewSPSItem(
			spsitem.OptData(item),
			spsitem.OptLayoutContent(item.LayoutContent),
			spsitem.OptLayoutDetail(item.LayoutDetail),
		),
		theme,
	)
	return item
}

func NewSPSItem(opts ...spsitem.Option) spsitem.Item {
	dimensions := spsitem.NewDimensions()
	dimensions.ContentWidth = 100
	dimensions.ContentHeight = 100
	highlightStyle := spsitem.HighlightStyleDefault
	//highlightStyle := spsitem.HighlightStyle(rand.RandomInt(int(spsitem.HighlightStyleDefault), int(spsitem.HighlightStyleCount-1)))
	stackAlignment := layout.Center

	item := spsitem.NewItem(
		spsitem.OptHighlightStyle(highlightStyle),
		spsitem.OptStackAlignment(stackAlignment),
		spsitem.OptDimensions(dimensions),
	)
	item.Update(opts...)
	return item
}
