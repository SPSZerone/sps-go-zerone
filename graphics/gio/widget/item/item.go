package item

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsdrawing "github.com/SPSZerone/sps-go-zerone/graphics/gio/architecture/drawing"
)

func NewItem(data any, content LayoutContent, opts ...Option) Item {
	i := Item{
		Data: data,
		Opts: NewOptions(content, opts...),
	}
	return i
}

type Item struct {
	Data any
	Opts Options
	UI   UI
}

func (i *Item) Update(opts ...Option) {
	for _, opt := range opts {
		opt(&i.Opts)
	}
}

func (i *Item) Layout(
	app *spsgio.Application, gtx layout.Context, highlight bool,
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

	dimensions = layout.Stack{Alignment: i.Opts.StackAlignment}.Layout(gtx,
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
			if i.UI.Clickable.Clicked(gtx) {
				clicked = true
			}
			clickDimensions := material.Clickable(gtx, &i.UI.Clickable, func(gtx layout.Context) layout.Dimensions {
				return layout.Dimensions{
					Size: layoutCtx.Size,
				}
			})
			return clickDimensions
		}),
		// content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if i.Opts.LayoutContent == nil {
				return layout.Dimensions{}
			}
			return i.Opts.LayoutContent(app, gtx, i, layoutCtx)
		}),
		// highlight
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if !highlight {
				return layout.Dimensions{}
			}
			switch i.Opts.HighlightStyle {
			case HighlightStyleTop:
				i.LayoutHighlightTop(app, gtx, layoutCtx)
			case HighlightStyleBottom:
				i.LayoutHighlightBottom(app, gtx, layoutCtx)
			case HighlightStyleStrokeRect:
				i.LayoutHighlightStrokeRect(app, gtx, layoutCtx)
			default:
				i.LayoutHighlightStrokeRect(app, gtx, layoutCtx)
			}
			return layout.Dimensions{
				Size: layoutCtx.Size,
			}
		}),
		// menu
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return i.UI.ContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min = image.Point{}
				return component.Menu(app.Theme, &i.UI.Menu).Layout(gtx)
			})
		}),
	)

	return
}

func (i *Item) LayoutHighlightStrokeRect(app *spsgio.Application, gtx layout.Context, layoutCtx LayoutContext) {
	size := image.Pt(
		layoutCtx.ContentWidth+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
		layoutCtx.ContentHeight+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
	)
	pos := image.Pt(layoutCtx.HighlightThicknessHalf, layoutCtx.HighlightThicknessHalf)
	spsdrawing.DrawStrokeRectR(
		gtx.Ops,
		size,
		app.Theme.Palette.ContrastBg,
		pos,
		float32(layoutCtx.HighlightThickness),
		layoutCtx.HighlightRoundness)
}

func (i *Item) LayoutHighlightTop(app *spsgio.Application, gtx layout.Context, layoutCtx LayoutContext) {
	contentBgRect := image.Rect(
		layoutCtx.Offset, 0,
		layoutCtx.Offset+layoutCtx.ContentWidth, layoutCtx.Offset)
	paint.FillShape(gtx.Ops, app.Theme.Palette.ContrastBg, clip.Rect(contentBgRect).Op())
}

func (i *Item) LayoutHighlightBottom(app *spsgio.Application, gtx layout.Context, layoutCtx LayoutContext) {
	contentBgRect := image.Rect(
		layoutCtx.Offset, layoutCtx.Size.Y-layoutCtx.Offset,
		layoutCtx.Offset+layoutCtx.ContentWidth, layoutCtx.Size.Y)
	paint.FillShape(gtx.Ops, app.Theme.Palette.ContrastBg, clip.Rect(contentBgRect).Op())
}
