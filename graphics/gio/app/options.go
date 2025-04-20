package app

import (
	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func NewOptions(newWindow spsgio.NewWindow, opts ...Option) Options {
	o := Options{
		NewWindow: newWindow,
	}
	o.Update(opts...)
	return o
}

type Option func(o *Options)

type Options struct {
	NewWindow spsgio.NewWindow

	OnCreate spsgio.OnAppCreate
	OnStart  spsgio.OnAppStart
	OnStop   spsgio.OnAppStop
}

func (o *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func OptNewWindow(value spsgio.NewWindow) Option {
	return func(o *Options) {
		o.NewWindow = value
	}
}

func OptOnCreate(value spsgio.OnAppCreate) Option {
	return func(o *Options) {
		o.OnCreate = value
	}
}

func OptOnStart(value spsgio.OnAppStart) Option {
	return func(o *Options) {
		o.OnStart = value
	}
}

func OptOnStop(value spsgio.OnAppStop) Option {
	return func(o *Options) {
		o.OnStop = value
	}
}
