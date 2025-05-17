package tab

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

func New(name string, data any) Tab {
	t := Tab{
		Name: name,
		Data: data,
	}
	t.BGColor1, t.BGColor2 = spscolor.Rand2Color(0, 10)
	return t
}

type Style func(style *material.LabelStyle)

type Tab struct {
	Name      string
	Clickable widget.Clickable
	Data      any

	BGColor1, BGColor2 color.NRGBA

	Size       image.Point
	sizeByName string
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
	dimensions = layout.Stack{Alignment: layout.Center}.Layout(gtx,
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

	name := func(gtx layout.Context) layout.Dimensions {
		dims := nameWidget(gtx)
		size := dims.Size
		if widthLimit > 0 {
			if isVertical {
				size.X = widthLimit
			}
		}
		t.Size = size
		t.sizeByName = t.Name
		return layout.Dimensions{Size: size}
	}
	if t.Size.X == 0 || t.Size.Y == 0 || t.sizeByName != t.Name {
		t.Clickable.Layout(gtx, name)
	}
	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rect(
				0, 0,
				t.Size.X, t.Size.Y,
			)
			spscolor.FillRect(gtx, rect, t.BGColor1, t.BGColor2)
			return layout.Dimensions{Size: t.Size}
		}),
		// name
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return t.Clickable.Layout(gtx, name)
		}),
	)
}

func (t *Tab) LayoutHighlight(
	theme *material.Theme, gtx layout.Context,
	highlightThickness int, highlight bool,
	isVertical bool,
	size image.Point,
) layout.Dimensions {
	if !highlight {
		return layout.Dimensions{}
	}

	var rect image.Rectangle
	if isVertical {
		// on right
		startX := size.X - highlightThickness
		rect = image.Rect(startX, 0, startX+highlightThickness, size.Y)
	} else {
		// on bottom
		rect = image.Rect(0, size.Y-highlightThickness, size.X, size.Y)
	}

	paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(rect).Op())
	return layout.Dimensions{Size: size}
}
