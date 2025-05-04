package tab

import (
	"gioui.org/font"
	"gioui.org/layout"

	spsnerdfont "github.com/SPSZerone/sps-go-zerone/graphics/gio/font/nerdfont"
	spssettings "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/settings"
)

func NewOptions(opts ...Option) Options {
	o := Options{
		Axis:              layout.Vertical,
		WidthWhenVertical: 128,
		Font:              spsnerdfont.MesloLGSNerdFontMono,
	}
	o.Update(opts...)
	return o
}

type Options struct {
	WidthWhenVertical int

	Axis        layout.Axis
	AxisSetting *spssettings.TabAxis

	Font font.Typeface
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

func OptAxisSetting(value *spssettings.TabAxis) Option {
	return func(o *Options) {
		o.AxisSetting = value
	}
}

func OptFont(value font.Typeface) Option {
	return func(o *Options) {
		o.Font = value
	}
}
