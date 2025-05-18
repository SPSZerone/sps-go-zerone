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
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsbg "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bg"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
)

func New(name string, data any) Tab {
	t := Tab{
		FlexInset: spslayout.New(),
		Name:      name,
		Data:      data,
	}
	t.FlexInset.Flex.Axis = layout.Horizontal
	t.BGColor1, t.BGColor2 = spscolor.Rand2Color(0, 10)
	//t.BGColor2 = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	return t
}

type (
	Style     func(style *material.LabelStyle)
	CloseMode int
)

const (
	CloseModeNone CloseMode = iota
	CloseModeNormal
	CloseModeMenu
)

type Tab struct {
	FlexInset spslayout.FlexInset

	Name           string
	Clickable      widget.Clickable
	CloseClickable widget.Clickable
	Data           any

	BGColor1, BGColor2 color.NRGBA
	size               image.Point

	close bool
}

func (t *Tab) Close() {
	t.close = true
}

func (t *Tab) IsClose() bool {
	return t.close
}

func (t *Tab) LayoutDefault(
	theme *material.Theme, gtx layout.Context,
	highlight bool,
	axis layout.Axis,
	widthLimit int,
	colorfulBG bool,
	closeMode CloseMode,
) (dimensions layout.Dimensions, clicked bool) {
	return t.Layout(
		theme, gtx,
		highlight,
		axis,
		widthLimit,
		colorfulBG,
		closeMode,
		func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(
				gtx,
				func(gtx layout.Context) layout.Dimensions {
					labelStyle := material.H6(theme, t.Name)
					return labelStyle.Layout(gtx)
				},
			)
		},
	)
}

func (t *Tab) Layout(
	theme *material.Theme, gtx layout.Context,
	highlight bool,
	axis layout.Axis,
	widthLimit int,
	colorfulBG bool,
	closeMode CloseMode,
	nameWidget layout.Widget,
) (dimensions layout.Dimensions, clicked bool) {
	if t.Clickable.Clicked(gtx) {
		clicked = true
	}
	if t.CloseClickable.Clicked(gtx) {
		t.Close()
	}

	isVertical := axis == layout.Vertical
	highlightThickness := gtx.Dp(unit.Dp(4))

	var size image.Point
	dimensions = layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// click area
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			dims := t.LayoutContent(
				theme, gtx,
				isVertical,
				widthLimit,
				highlightThickness,
				colorfulBG,
				closeMode,
				nameWidget,
			)
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

func (t *Tab) LayoutContent(
	theme *material.Theme, gtx layout.Context,
	isVertical bool,
	widthLimit int,
	highlightThickness int,
	colorfulBG bool,
	closeMode CloseMode,
	nameWidget layout.Widget,
) layout.Dimensions {
	if widthLimit > 0 {
		if isVertical {
			gtx.Constraints.Max.X = widthLimit
		}
	}

	if !colorfulBG {
		return t.doLayoutContent(theme, gtx, isVertical, widthLimit, closeMode, nameWidget)
	}

	return layout.Stack{
		Alignment: layout.Center,
	}.Layout(
		gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max = t.size
			return spsbg.NewColorful(t.BGColor1, t.BGColor2).LayoutBG(theme, gtx, t.size)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			dim := t.doLayoutContent(theme, gtx, isVertical, widthLimit, closeMode, nameWidget)
			t.size = dim.Size
			return dim
		}),
	)
}

func (t *Tab) doLayoutContent(
	theme *material.Theme, gtx layout.Context,
	isVertical bool,
	widthLimit int,
	closeMode CloseMode,
	nameWidget layout.Widget,
) (dimensions layout.Dimensions) {
	if closeMode == CloseModeNormal {
		spacer := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spsspacer.NewWithWidth(8).Layout(gtx)
		})
		dimensions = layout.Flex{
			Axis:      layout.Horizontal,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return t.Clickable.Layout(gtx, nameWidget)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				//   󰅖    󰅗 󰅙 󰅜 󰛉 󱎘 󰖭  󰅘 󰅚 󰅝
				//return material.Button(theme, &t.Close, ``).Layout(gtx)
				return material.IconButton(theme, &t.CloseClickable, spsicon.NavigationClose, "Close").Layout(gtx)
			}),
			spacer,
		)
		return
	}

	if closeMode == CloseModeMenu {
		// TODO
		dimensions = t.Clickable.Layout(gtx, nameWidget)
		return
	}

	name := func(gtx layout.Context) layout.Dimensions {
		dims := nameWidget(gtx)
		size := dims.Size
		if widthLimit > 0 {
			if isVertical {
				size.X = widthLimit
			}
		}
		return layout.Dimensions{Size: size}
	}

	dimensions = t.Clickable.Layout(gtx, name)
	return
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
