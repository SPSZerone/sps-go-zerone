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

type OnCreate func(app *App)
type OnStart func(app *App)
type OnStop func(app *App)

type Options struct {
	NewWindow spsgio.NewWindow

	OnCreate OnCreate
	OnStart  OnStart
	OnStop   OnStop
}

func (o *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func OptOnCreate(value OnCreate) Option {
	return func(o *Options) {
		o.OnCreate = value
	}
}

func OptOnStart(value OnStart) Option {
	return func(o *Options) {
		o.OnStart = value
	}
}

func OptOnStop(value OnStop) Option {
	return func(o *Options) {
		o.OnStop = value
	}
}

func OptNewWindow(value spsgio.NewWindow) Option {
	return func(o *Options) {
		o.NewWindow = value
	}
}
