package item

import (
	"image/color"

	"gioui.org/layout"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

type HighlightStyle int

const (
	HighlightStyleDefault HighlightStyle = iota
	HighlightStyleStrokeRect
	HighlightStyleTop
	HighlightStyleBottom
	HighlightStyleCount
)

type LayoutContent func(app *spsgio.Application, gtx layout.Context, item *Item, layoutCtx LayoutContext) layout.Dimensions

type LayoutDetail func(app *spsgio.Application, gtx layout.Context, item *Item) layout.Dimensions

func NewOptions(content LayoutContent, opts ...Option) Options {
	o := Options{
		Dimensions:     NewDimensions(),
		StackAlignment: layout.Center,
		BgColor:        spscolor.DynamicColor(2),
		HighlightStyle: HighlightStyleStrokeRect,
		LayoutContent:  content,
	}
	o.Update(opts...)
	return o
}

type Options struct {
	Dimensions     Dimensions
	StackAlignment layout.Direction
	BgColor        color.NRGBA
	HighlightStyle HighlightStyle

	LayoutContent LayoutContent
	LayoutDetail  LayoutDetail
}

func (p *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type Option func(*Options)

func OptDimensions(value Dimensions) Option {
	return func(o *Options) {
		o.Dimensions = value
	}
}

func OptStackAlignment(value layout.Direction) Option {
	return func(o *Options) {
		o.StackAlignment = value
	}
}

func OptBgColor(value color.NRGBA) Option {
	return func(o *Options) {
		o.BgColor = value
	}
}

func OptHighlightStyle(value HighlightStyle) Option {
	return func(o *Options) {
		o.HighlightStyle = value
	}
}

func OptLayoutContent(value LayoutContent) Option {
	return func(o *Options) {
		o.LayoutContent = value
	}
}

func OptLayoutDetail(value LayoutDetail) Option {
	return func(o *Options) {
		o.LayoutDetail = value
	}
}
