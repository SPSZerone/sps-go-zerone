package tab

import (
	"gioui.org/layout"

	spssettings "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/settings"
)

func NewTabsOptions(opts ...TabsOption) TabsOptions {
	o := TabsOptions{
		WidthLimitWhenVertical: 128,
		Axis:                   layout.Vertical,
		ColorfulBG:             true,
		CloseMode:              CloseModeNone,
	}
	o.Update(opts...)
	return o
}

func NewOptions(opts ...Option) Options {
	o := Options{
		CloseMode:            CloseModeNone,
		CloseModeInheritTabs: true,
	}
	o.Update(opts...)
	return o
}

type (
	TabsOption func(*TabsOptions)
	Option     func(*Options)

	OnClose func(tab *Tab) (close bool)
)

type TabsOptions struct {
	WidthLimitWhenVertical int

	Axis        layout.Axis
	AxisSetting *spssettings.TabAxis

	ColorfulBG bool
	CloseMode  CloseMode
}

func (o *TabsOptions) Update(opts ...TabsOption) {
	for _, opt := range opts {
		opt(o)
	}
}

type Options struct {
	CloseMode            CloseMode
	CloseModeInheritTabs bool

	OnClose OnClose
}

func (o *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func TabsOptWidthWhenVertical(value int) TabsOption {
	return func(o *TabsOptions) {
		o.WidthLimitWhenVertical = value
	}
}

func TabsOptAxis(value layout.Axis) TabsOption {
	return func(o *TabsOptions) {
		o.Axis = value
	}
}

func TabsOptAxisSetting(value *spssettings.TabAxis) TabsOption {
	return func(o *TabsOptions) {
		o.AxisSetting = value
	}
}

func TabsOptColorfulBG(value bool) TabsOption {
	return func(o *TabsOptions) {
		o.ColorfulBG = value
	}
}

func TabsOptCloseMode(value CloseMode) TabsOption {
	return func(o *TabsOptions) {
		o.CloseMode = value
	}
}

func OptCloseMode(value CloseMode) Option {
	return func(o *Options) {
		o.CloseMode = value
	}
}

func OptCloseModeInheritTabs(value bool) Option {
	return func(o *Options) {
		o.CloseModeInheritTabs = value
	}
}

func OptOnClose(value OnClose) Option {
	return func(o *Options) {
		o.OnClose = value
	}
}
