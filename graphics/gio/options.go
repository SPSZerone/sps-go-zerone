package gio

import (
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

type Option func(o *Options)

type OnCreate func(app *App)
type OnStart func(app *App)
type OnStop func(app *App)

type Options struct {
	OnCreate OnCreate

	OnStart OnStart
	OnStop  OnStop

	WinOpts []spswin.Option
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

func OptWinOpts(value ...spswin.Option) Option {
	return func(o *Options) {
		o.WinOpts = value
	}
}
