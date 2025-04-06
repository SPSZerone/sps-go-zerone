package gio

import "gioui.org/io/event"

type Option func(o *Options)

type OnInitPre func(app *Application)
type OnInitPost func(app *Application)
type OnStart func(app *Application)
type OnLoop func(app *Application) error
type OnStop func(app *Application)

type OnEventPre func(app *Application, evt event.Event, param any)
type OnEventPost func(app *Application, evt event.Event, param any)

type Options struct {
	LoopMode    LoopMode
	Title       string
	OnInitPre   OnInitPre
	OnInitPost  OnInitPost
	OnStart     OnStart
	OnLoop      OnLoop
	OnStop      OnStop
	OnEventPre  OnEventPre
	OnEventPost OnEventPost
}

func OptLoopMode(value LoopMode) Option {
	return func(o *Options) {
		o.LoopMode = value
	}
}

func OptTitle(value string) Option {
	return func(o *Options) {
		o.Title = value
	}
}

func OptOnInitPre(value OnInitPre) Option {
	return func(o *Options) {
		o.OnInitPre = value
	}
}

func OptOnInitPost(value OnInitPost) Option {
	return func(o *Options) {
		o.OnInitPost = value
	}
}

func OptOnStart(value OnStart) Option {
	return func(o *Options) {
		o.OnStart = value
	}
}

func OptOnLoop(value OnLoop) Option {
	return func(o *Options) {
		o.OnLoop = value
	}
}

func OptOnStop(value OnStop) Option {
	return func(o *Options) {
		o.OnStop = value
	}
}

func OptOnEventPre(value OnEventPre) Option {
	return func(o *Options) {
		o.OnEventPre = value
	}
}

func OptOnEventPost(value OnEventPost) Option {
	return func(o *Options) {
		o.OnEventPost = value
	}
}
