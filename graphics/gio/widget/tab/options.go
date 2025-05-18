package tab

import (
	"gioui.org/layout"

	spssettings "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/settings"
)

func NewOptions(opts ...Option) Options {
	o := Options{
		WidthLimitWhenVertical: 128,
		Axis:                   layout.Vertical,
		ColorfulBG:             true,
		CloseMode:              CloseModeNone,
	}
	o.Update(opts...)
	return o
}

type Options struct {
	WidthLimitWhenVertical int

	Axis        layout.Axis
	AxisSetting *spssettings.TabAxis

	ColorfulBG bool
	CloseMode  CloseMode
}

func (o *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(*Options)

func OptWidthWhenVertical(value int) Option {
	return func(o *Options) {
		o.WidthLimitWhenVertical = value
	}
}

func OptAxis(value layout.Axis) Option {
	return func(o *Options) {
		o.Axis = value
	}
}

func OptAxisSetting(value *spssettings.TabAxis) Option {
	return func(o *Options) {
		o.AxisSetting = value
	}
}

func OptColorfulBG(value bool) Option {
	return func(o *Options) {
		o.ColorfulBG = value
	}
}

func OptCloseMode(value CloseMode) Option {
	return func(o *Options) {
		o.CloseMode = value
	}
}
