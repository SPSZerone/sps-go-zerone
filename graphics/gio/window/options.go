package window

import (
	"gioui.org/io/event"
	"gioui.org/io/system"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

type Option func(win *Window)

type OnInitPre func(win *Window)
type OnInitPost func(win *Window)
type OnStart func(win *Window)
type OnLoop func(win *Window) error
type OnStop func(win *Window)

type OnEventPre func(win *Window, evt event.Event, param any)
type OnEventPost func(win *Window, evt event.Event, param any)

type Options struct {
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
	return func(win *Window) {
		win.ID = value
	}
}

func OptTitle(value string) Option {
	return func(win *Window) {
		win.Title = value
	}
}

func OptPref(value spspref.Preferences) Option {
	return func(win *Window) {
		win.Pref = value
	}
}

func OptStartAction(value system.Action) Option {
	return func(win *Window) {
		win.Opts.StartAction = value
	}
}

func OptLoopMode(value LoopMode) Option {
	return func(win *Window) {
		win.Opts.LoopMode = value
	}
}

func OptOnInitPre(value OnInitPre) Option {
	return func(win *Window) {
		win.Opts.OnInitPre = value
	}
}

func OptOnInitPost(value OnInitPost) Option {
	return func(win *Window) {
		win.Opts.OnInitPost = value
	}
}

func OptOnStart(value OnStart) Option {
	return func(win *Window) {
		win.Opts.OnStart = value
	}
}

func OptOnLoop(value OnLoop) Option {
	return func(win *Window) {
		win.Opts.OnLoop = value
	}
}

func OptOnStop(value OnStop) Option {
	return func(win *Window) {
		win.Opts.OnStop = value
	}
}

func OptOnEventPre(value OnEventPre) Option {
	return func(win *Window) {
		win.Opts.OnEventPre = value
	}
}

func OptOnEventPost(value OnEventPost) Option {
	return func(win *Window) {
		win.Opts.OnEventPost = value
	}
}
