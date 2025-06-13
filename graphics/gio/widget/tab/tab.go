package tab

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func New(name string, data any, opts ...Option) Tab {
	t := Tab{
		Name: name,
		Data: data,
		Opts: NewOptions(opts...),
		UI:   NewUI(),
	}
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
	Name string
	Data any

	Opts Options
	UI   UI

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
) (dimensions layout.Dimensions, clicked bool) {
	return t.Layout(
		theme, gtx,
		highlight,
		axis,
		widthLimit,
		colorfulBG,
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
	nameWidget layout.Widget,
) (dimensions layout.Dimensions, clicked bool) {
	if t.UI.Clickable.Clicked(gtx) {
		clicked = true
	}

	if t.UI.CloseClickable.Clicked(gtx) {
		isClose := true
		if t.Opts.OnClose != nil {
			isClose = t.Opts.OnClose(t)
		}

		if isClose {
			t.Close()
		}
	}

	isVertical := axis == layout.Vertical
	highlightThickness := gtx.Dp(unit.Dp(4))

	var size image.Point
	dimensions = layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// click area
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			dims := t.UI.LayoutContent(
				theme, gtx,
				isVertical,
				widthLimit,
				highlightThickness,
				colorfulBG,
				t.Opts.CloseMode,
				t.Name,
				nameWidget,
			)
			size = dims.Size
			return dims
		}),
		// highlight
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return t.UI.LayoutHighlight(theme, gtx, highlightThickness, highlight, isVertical, size)
		}),
	)

	return
}
