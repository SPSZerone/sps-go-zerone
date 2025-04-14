package tab

import (
	"gioui.org/layout"
)

func NewOptions() Options {
	return Options{
		Axis: layout.Vertical,
	}
}

type Options struct {
	Axis layout.Axis
}

type Option func(*Options)

func OptAxis(axis layout.Axis) Option {
	return func(o *Options) {
		o.Axis = axis
	}
}
