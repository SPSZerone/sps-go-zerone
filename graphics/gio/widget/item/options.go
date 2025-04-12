package item

import (
	"image/color"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

func NewOptions() Options {
	return Options{
		Dimensions: NewDimensions(),
		BgColor:    spscolor.DynamicColor(0),
	}
}

type Options struct {
	Dimensions Dimensions
	BgColor    color.NRGBA
}

type Option func(*Options)

func OptDimensions(value Dimensions) Option {
	return func(o *Options) {
		o.Dimensions = value
	}
}

func OptBgColor(value color.NRGBA) Option {
	return func(o *Options) {
		o.BgColor = value
	}
}
