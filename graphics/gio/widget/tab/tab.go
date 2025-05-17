package tab

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func New(name string) Tab {
	return Tab{
		Name: name,
	}
}

type Style func(style *material.LabelStyle)

type Tab struct {
	Name      string
	Clickable widget.Clickable
	Data      any
}

func (t *Tab) LayoutDefault(
	theme *material.Theme, gtx layout.Context,
	highlight bool,
	axis layout.Axis,
	widthLimit int,
) (dimensions layout.Dimensions, clicked bool) {
	return t.Layout(theme, gtx, highlight, axis, widthLimit, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(12)).Layout(
			gtx,
			func(gtx layout.Context) layout.Dimensions {
				labelStyle := material.H6(theme, t.Name)
				return labelStyle.Layout(gtx)
			},
		)
	})
}

func (t *Tab) Layout(
	theme *material.Theme, gtx layout.Context,
	highlight bool,
	axis layout.Axis,
	widthLimit int,
	nameWidget layout.Widget,
) (dimensions layout.Dimensions, clicked bool) {
	if t.Clickable.Clicked(gtx) {
		clicked = true
	}

	isVertical := axis == layout.Vertical
	highlightThickness := gtx.Dp(unit.Dp(4))

	var size image.Point
	dimensions = layout.Stack{Alignment: layout.S}.Layout(gtx,
		// click area
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			dims := t.LayoutName(theme, gtx, isVertical, widthLimit, highlightThickness, nameWidget)
			size = dims.Size
			return dims
		}),
		// highlight
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return t.LayoutHighlight(theme, gtx, highlightThickness, highlight, isVertical, size)
		}),
	)

	return
}

func (t *Tab) LayoutName(
	theme *material.Theme, gtx layout.Context,
	isVertical bool,
	widthLimit int,
	highlightThickness int,
	nameWidget layout.Widget,
) layout.Dimensions {
	if widthLimit > 0 {
		if isVertical {
			gtx.Constraints.Max.X = widthLimit
		}
	}

	return t.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		dims := nameWidget(gtx)
		size := dims.Size
		if widthLimit > 0 {
			if isVertical {
				size.X = widthLimit
			}
		}
		return layout.Dimensions{Size: size}
	})
}

func (t *Tab) LayoutHighlight(
	theme *material.Theme, gtx layout.Context,
	highlightThickness int, highlight bool,
	isVertical bool,
	tabSize image.Point,
) layout.Dimensions {
	if !highlight {
		return layout.Dimensions{}
	}

	var highlightRect image.Rectangle
	var size image.Point
	if isVertical {
		size = image.Pt(highlightThickness, tabSize.Y)
		// right
		startX := tabSize.X>>1 - highlightThickness>>1
		// left
		//startX = -startX
		highlightRect = image.Rect(startX, 0, startX+highlightThickness, tabSize.Y)
	} else {
		size = image.Pt(tabSize.X, highlightThickness)
		highlightRect = image.Rect(0, 0, tabSize.X, highlightThickness)
	}

	paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(highlightRect).Op())
	return layout.Dimensions{Size: size}
}
