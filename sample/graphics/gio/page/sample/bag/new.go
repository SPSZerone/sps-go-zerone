package bag

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func NewItem(data Data, theme *material.Theme) *Item {
	item := newItem(data)
	spsItem := NewSPSItem(spsitem.OptData(item))
	item.Init(spsItem, theme)
	return item
}

func NewSPSItem(opts ...spsitem.Option) spsitem.Item {
	dimensions := spsitem.NewDimensions()
	dimensions.InsetInner = layout.UniformInset(unit.Dp(4))
	//dimensions.InsetInner = layout.Inset{Left: unit.Dp(0), Right: unit.Dp(16), Top: unit.Dp(8), Bottom: unit.Dp(32)}
	dimensions.ContentSize.X = 100
	dimensions.ContentSize.Y = 100
	highlightStyle := spsitem.HighlightStyleDefault
	//highlightStyle := spsitem.HighlightStyle(rand.RandomInt(int(spsitem.HighlightStyleDefault), int(spsitem.HighlightStyleCount-1)))
	stackAlignment := layout.Center

	item := spsitem.New(
		spsitem.OptHighlightStyle(highlightStyle),
		spsitem.OptStackAlignment(stackAlignment),
		spsitem.OptDimensions(dimensions),
	)
	item.Update(opts...)
	return item
}
