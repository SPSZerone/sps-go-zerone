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
	content func(gtx layout.Context) layout.Dimensions,
) (dimensions layout.Dimensions, clicked bool) {
	contentWidth := gtx.Dp(unit.Dp(i.Opts.Dimensions.ContentWidth))
	contentHeight := gtx.Dp(unit.Dp(i.Opts.Dimensions.ContentHeight))
	padding := gtx.Dp(unit.Dp(i.Opts.Dimensions.Padding))
	highlightThickness := gtx.Dp(unit.Dp(i.Opts.Dimensions.HighlightThickness))
	highlightRoundness := gtx.Dp(unit.Dp(i.Opts.Dimensions.HighlightRoundness))

	doublePadding := padding << 1
	doubleHighlightThickness := highlightThickness << 1
	halfHighlightThickness := highlightThickness >> 1
	offset := padding + highlightThickness
	width := contentWidth + doublePadding + doubleHighlightThickness
	height := contentHeight + doublePadding + doubleHighlightThickness
	totalSize := image.Point{X: width, Y: height}

	dimensions = layout.Stack{Alignment: layout.S}.Layout(gtx,
		// content background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			contentBgRect := image.Rect(offset, offset, offset+contentWidth, offset+contentHeight)
			paint.FillShape(gtx.Ops, i.Opts.BgColor, clip.Rect(contentBgRect).Op())
			return layout.Dimensions{
				Size: totalSize,
			}
		}),
		// click area
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if i.Clickable.Clicked(gtx) {
				clicked = true
			}
			clickDimensions := material.Clickable(gtx, &i.Clickable, func(gtx layout.Context) layout.Dimensions {
				return layout.Dimensions{
					Size: totalSize,
				}
			})
			return clickDimensions
		}),
		// content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return content(gtx)
		}),
		// highlight
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if !highlight {
				return layout.Dimensions{}
			}
			size := image.Pt(
				contentWidth+doublePadding+highlightThickness,
				contentHeight+doublePadding+highlightThickness,
			)
			pos := image.Pt(halfHighlightThickness, halfHighlightThickness)
			spsdrawing.DrawStrokeRectR(
				&app.Ops,
				size,
				app.Theme.Palette.ContrastBg,
				pos,
				float32(highlightThickness),
				highlightRoundness)
			return layout.Dimensions{
				Size: totalSize,
			}
		}),
	)

	return
}
