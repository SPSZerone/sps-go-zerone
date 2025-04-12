package item

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsdrawing "github.com/SPSZerone/sps-go-zerone/graphics/gio/architecture/drawing"
)

func NewItem(data any, opts ...Option) Item {
	i := Item{
		Data: data,
		Opts: NewOptions(),
	}
	i.Update(opts...)
	return i
}

type Item struct {
	Data any
	Opts Options

	Clickable widget.Clickable
}

func (i *Item) Update(opts ...Option) {
	for _, opt := range opts {
		opt(&i.Opts)
	}
}

func (i *Item) Layout(
	app *spsgio.Application, gtx layout.Context, highlight bool,
	content func(gtx layout.Context, layoutCtx LayoutContext) layout.Dimensions,
) (dimensions layout.Dimensions, clicked bool) {
	layoutCtx := LayoutContext{
		ContentWidth:       gtx.Dp(unit.Dp(i.Opts.Dimensions.ContentWidth)),
		ContentHeight:      gtx.Dp(unit.Dp(i.Opts.Dimensions.ContentHeight)),
		Padding:            gtx.Dp(unit.Dp(i.Opts.Dimensions.Padding)),
		HighlightThickness: gtx.Dp(unit.Dp(i.Opts.Dimensions.HighlightThickness)),
		HighlightRoundness: gtx.Dp(unit.Dp(i.Opts.Dimensions.HighlightRoundness)),
	}
	layoutCtx.PaddingDouble = layoutCtx.Padding << 1
	layoutCtx.HighlightThicknessDouble = layoutCtx.HighlightThickness << 1
	layoutCtx.HighlightThicknessHalf = layoutCtx.HighlightThickness >> 1
	layoutCtx.Offset = layoutCtx.Padding + layoutCtx.HighlightThickness
	layoutCtx.Size = image.Pt(
		layoutCtx.ContentWidth+layoutCtx.PaddingDouble+layoutCtx.HighlightThicknessDouble,
		layoutCtx.ContentHeight+layoutCtx.PaddingDouble+layoutCtx.HighlightThicknessDouble,
	)

	dimensions = layout.Stack{Alignment: layout.S}.Layout(gtx,
		// content background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			contentBgRect := image.Rect(
				layoutCtx.Offset, layoutCtx.Offset,
				layoutCtx.Offset+layoutCtx.ContentWidth, layoutCtx.Offset+layoutCtx.ContentHeight)
			paint.FillShape(gtx.Ops, i.Opts.BgColor, clip.Rect(contentBgRect).Op())
			return layout.Dimensions{
				Size: layoutCtx.Size,
			}
		}),
		// click area
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if i.Clickable.Clicked(gtx) {
				clicked = true
			}
			clickDimensions := material.Clickable(gtx, &i.Clickable, func(gtx layout.Context) layout.Dimensions {
				return layout.Dimensions{
					Size: layoutCtx.Size,
				}
			})
			return clickDimensions
		}),
		// content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return content(gtx, layoutCtx)
		}),
		// highlight
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if !highlight {
				return layout.Dimensions{}
			}
			size := image.Pt(
				layoutCtx.ContentWidth+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
				layoutCtx.ContentHeight+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
			)
			pos := image.Pt(layoutCtx.HighlightThicknessHalf, layoutCtx.HighlightThicknessHalf)
			spsdrawing.DrawStrokeRectR(
				&app.Ops,
				size,
				app.Theme.Palette.ContrastBg,
				pos,
				float32(layoutCtx.HighlightThickness),
				layoutCtx.HighlightRoundness)
			return layout.Dimensions{
				Size: layoutCtx.Size,
			}
		}),
	)

	return
}
