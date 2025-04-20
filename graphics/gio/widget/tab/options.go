package tab

import (
	"gioui.org/layout"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/settings"
)

func NewOptions(opts ...Option) Options {
	o := Options{
		Axis:              layout.Vertical,
		WidthWhenVertical: 150,
	}
	o.Update(opts...)
	return o
}

type Options struct {
	WidthWhenVertical int

	Axis        layout.Axis
	AxisSetting *settings.TabAxis
}

func (o *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(*Options)

func OptWidthWhenVertical(value int) Option {
	return func(o *Options) {
		o.WidthWhenVertical = value
	}
}

func OptAxis(value layout.Axis) Option {
	return func(o *Options) {
		o.Axis = value
	}
}

func OptAxisSetting(value *settings.TabAxis) Option {
	return func(o *Options) {
		o.AxisSetting = value
	}
}
