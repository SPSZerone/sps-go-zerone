package bag

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spsrand "github.com/SPSZerone/sps-go-zerone/math/rand"
)

func NewItem(data Data, theme *material.Theme) *Item {
	item := newItem(data)
	spsItem := NewSPSItem(spsitem.OptData(item))
	item.Init(spsItem, theme)
	return item
}

func NewSPSItem(opts ...spsitem.Option) spsitem.Item {
	dimensions := spsitem.NewDimensions()
	dimensions.ContentSize.X = 100
	dimensions.ContentSize.Y = 100
	//highlightStyle := spsitem.HighlightStyleDefault
	highlightStyle := spsitem.HighlightStyle(spsrand.RandomInt(int(spsitem.HighlightStyleDefault), int(spsitem.HighlightStyleCount-1)))
	stackAlignment := layout.Center

	item := spsitem.New(
		spsitem.OptHighlightStyle(highlightStyle),
		spsitem.OptStackAlignment(stackAlignment),
		spsitem.OptDimensions(dimensions),
		spsitem.OptBgColor(spscolor.Rand2Color(0, 10)),
	)
	item.Update(opts...)

	//item.Opts.Surface.InsetOuter, item.Opts.Surface.InsetInner = testItemSurfaceInset()
	return item
}

func testItemSurfaceInset() (outer, inner layout.Inset) {
	outer = layout.Inset{
		Left:   unit.Dp(16),
		Top:    unit.Dp(12),
		Right:  unit.Dp(8),
		Bottom: unit.Dp(4),
	}
	inner = layout.Inset{
		Left:   unit.Dp(4),
		Top:    unit.Dp(8),
		Right:  unit.Dp(12),
		Bottom: unit.Dp(16),
	}
	return
}
