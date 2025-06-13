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

type (
	TabsOption func(*TabsOptions)
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
