package item

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

type HighlightStyle int

const (
	HighlightStyleDefault HighlightStyle = iota
	HighlightStyleStrokeRect
	HighlightStyleTop
	HighlightStyleBottom
	HighlightStyleCount
)

type LayoutContent func(
	theme *material.Theme, gtx layout.Context,
	item *Item, layoutCtx LayoutContext,
) layout.Dimensions

type LayoutDetail func(
	theme *material.Theme, gtx layout.Context, item *Item,
) layout.Dimensions

func NewOptions(opts ...Option) Options {
	o := Options{
		Surface:        spssurface.NewUniformInset(4, 2),
		Dimensions:     NewDimensions(),
		StackAlignment: layout.Center,
		BgColor:        spscolor.DynamicColor(2),
		HighlightStyle: HighlightStyleStrokeRect,
	}
	o.Update(opts...)
	return o
}

type Option func(*Options)

type Options struct {
	Surface        spssurface.Surface
	Dimensions     Dimensions
	StackAlignment layout.Direction
	BgColor        color.NRGBA
	HighlightStyle HighlightStyle

	Data any
}

func (p *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(p)
	}
}

func OptSurface(value spssurface.Surface) Option {
	return func(o *Options) {
		o.Surface = value
	}
}

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

func OptData(value any) Option {
	return func(o *Options) {
		o.Data = value
	}
}
