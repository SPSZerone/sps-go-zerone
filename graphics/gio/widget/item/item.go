package item

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsdrawing "github.com/SPSZerone/sps-go-zerone/graphics/gio/architecture/drawing"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func New(opts ...Option) Item {
	i := Item{
		Opts: NewOptions(opts...),
		UI:   NewUI(),
	}
	i.UI.Surface.InsetOuter = i.Opts.Dimensions.InsetOuter
	i.UI.Surface.InsetInner = i.Opts.Dimensions.InsetInner
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
	theme *material.Theme, gtx layout.Context,
	highlight bool,
	layoutContent LayoutContent,
) (dimensions layout.Dimensions, clicked bool) {
	layoutCtx := NewLayoutContext(gtx, i.Opts.Dimensions)

	dimensions = layout.Stack{Alignment: i.Opts.StackAlignment}.Layout(gtx,
		// content background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return i.UI.Surface.Layout(
				theme, gtx,
				func(gtx layout.Context) layout.Dimensions {
					contentBgRect := image.Rect(
						layoutCtx.Offset, layoutCtx.Offset,
						layoutCtx.Offset+layoutCtx.ContentSize.X, layoutCtx.Offset+layoutCtx.ContentSize.Y)
					paint.FillShape(gtx.Ops, i.Opts.BgColor, clip.Rect(contentBgRect).Op())
					return layout.Dimensions{
						Size: layoutCtx.Size,
					}
				},
				func(surface *spssurface.Surface, style *component.SurfaceStyle) {
					style.Fill = color.NRGBA{R: 0xD3, G: 0xD3, B: 0xD3, A: 0xFF}
					//style.Fill = color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0x7F}
				},
			)
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
		i.UI.Menu.LayoutExpandedContextArea(theme, image.Point{}),
	)

	return
}

func (i *Item) LayoutHighlightStrokeRect(theme *material.Theme, gtx layout.Context, layoutCtx LayoutContext) {
	size := image.Pt(
		layoutCtx.ContentSize.X+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
		layoutCtx.ContentSize.Y+layoutCtx.PaddingDouble+layoutCtx.HighlightThickness,
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
		layoutCtx.Offset+layoutCtx.ContentSize.X, layoutCtx.Offset)
	paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(contentBgRect).Op())
}

func (i *Item) LayoutHighlightBottom(theme *material.Theme, gtx layout.Context, layoutCtx LayoutContext) {
	contentBgRect := image.Rect(
		layoutCtx.Offset, layoutCtx.Size.Y-layoutCtx.Offset,
		layoutCtx.Offset+layoutCtx.ContentSize.X, layoutCtx.Size.Y)
	paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(contentBgRect).Op())
}
