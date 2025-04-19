package window

import (
	"gioui.org/io/event"
	"gioui.org/io/system"
)

type Option func(o *Options)

type OnInitPre func(app *Window)
type OnInitPost func(app *Window)
type OnStart func(app *Window)
type OnLoop func(app *Window) error
type OnStop func(app *Window)

type OnEventPre func(app *Window, evt event.Event, param any)
type OnEventPost func(app *Window, evt event.Event, param any)

type Options struct {
	ID    any
	Title string

	StartAction system.Action
	LoopMode    LoopMode

	OnInitPre   OnInitPre
	OnInitPost  OnInitPost
	OnStart     OnStart
	OnLoop      OnLoop
	OnStop      OnStop
	OnEventPre  OnEventPre
	OnEventPost OnEventPost
}

func OptID(value any) Option {
	return func(o *Options) {
		o.ID = value
	}
}

func OptTitle(value string) Option {
	return func(o *Options) {
		o.Title = value
	}
}

func OptStartAction(value system.Action) Option {
	return func(o *Options) {
		o.StartAction = value
	}
}

func OptLoopMode(value LoopMode) Option {
	return func(o *Options) {
		o.LoopMode = value
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
