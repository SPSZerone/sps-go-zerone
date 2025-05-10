package item

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spsdrawing "github.com/SPSZerone/sps-go-zerone/graphics/gio/architecture/drawing"
)

func NewItem(opts ...Option) Item {
	i := Item{
		Opts: NewOptions(opts...),
	}
	return i
}

type Item struct {
	Opts Options
	UI   UI
}

func (i *Item) Update(opts ...Option) {
	for _, opt := range opts {
		opt(&i.Opts)
	}
}

func (i *Item) Layout(
	theme *material.Theme, gtx layout.Context, highlight bool,
	layoutContent LayoutContent,
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
			return i.UI.Clickable.Layout(
				gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{
						Size: layoutCtx.Size,
					}
				},
			)
		}),
		// content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if layoutContent == nil {
				return layout.Dimensions{}
			}
			return layoutContent(theme, gtx, i, layoutCtx)
		}),
		// highlight
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if !highlight {
				return layout.Dimensions{}
			}
			switch i.Opts.HighlightStyle {
			case HighlightStyleTop:
				i.LayoutHighlightTop(theme, gtx, layoutCtx)
			case HighlightStyleBottom:
				i.LayoutHighlightBottom(theme, gtx, layoutCtx)
			case HighlightStyleStrokeRect:
				i.LayoutHighlightStrokeRect(theme, gtx, layoutCtx)
			default:
				i.LayoutHighlightStrokeRect(theme, gtx, layoutCtx)
			}
			return layout.Dimensions{
				Size: layoutCtx.Size,
			}
		}),
		// menu
		i.UI.Menu.LayoutExpandedContextArea(theme, gtx),
	)

	return
}

func (i *Item) LayoutHighlightStrokeRect(theme *material.Theme, gtx layout.Context, layoutCtx LayoutContext) {
	size := image.Pt(
		layoutCtx.ContentWidth+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
		layoutCtx.ContentHeight+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
	)
	pos := image.Pt(layoutCtx.HighlightThicknessHalf, layoutCtx.HighlightThicknessHalf)
	spsdrawing.DrawStrokeRectR(
		gtx.Ops,
		size,
		theme.Palette.ContrastBg,
		pos,
		float32(layoutCtx.HighlightThickness),
		layoutCtx.HighlightRoundness)
}

func (i *Item) LayoutHighlightTop(theme *material.Theme, gtx layout.Context, layoutCtx LayoutContext) {
	contentBgRect := image.Rect(
		layoutCtx.Offset, 0,
		layoutCtx.Offset+layoutCtx.ContentWidth, layoutCtx.Offset)
	paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(contentBgRect).Op())
}

func (i *Item) LayoutHighlightBottom(theme *material.Theme, gtx layout.Context, layoutCtx LayoutContext) {
	contentBgRect := image.Rect(
		layoutCtx.Offset, layoutCtx.Size.Y-layoutCtx.Offset,
		layoutCtx.Offset+layoutCtx.ContentWidth, layoutCtx.Size.Y)
	paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(contentBgRect).Op())
}
