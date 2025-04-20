package gio

import (
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

func NewOptions(newWindow NewWindow, opts ...Option) Options {
	o := Options{
		NewWindow: newWindow,
	}
	o.Update(opts...)
	return o
}

type NewWindow func(app *App) *spswin.Window

type Option func(o *Options)

type OnCreate func(app *App)
type OnStart func(app *App)
type OnStop func(app *App)

type Options struct {
	NewWindow NewWindow

	OnCreate OnCreate
	OnStart  OnStart
	OnStop   OnStop

	//WinOpts   []spswin.Option
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

func OptNewWindow(value NewWindow) Option {
	return func(o *Options) {
		o.NewWindow = value
	}
}

//func OptWinOpts(value ...spswin.Option) Option {
//	return func(o *Options) {
//		o.WinOpts = value
//	}
//}
