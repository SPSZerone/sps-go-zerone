package window

import (
	"gioui.org/io/system"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

const (
	DecoActionNone = 0
)

type Option func(win *Window)

type Options struct {
	StartAction system.Action
	LoopMode    LoopMode

	OnInitPre   spsgio.OnWinInitPre
	OnInitPost  spsgio.OnWinInitPost
	OnStart     spsgio.OnStart
	OnLoop      spsgio.OnWinLoop
	OnStop      spsgio.OnStop
	OnEventPre  spsgio.OnWinEventPre
	OnEventPost spsgio.OnWinEventPost
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

func OptOnInitPre(value spsgio.OnWinInitPre) Option {
	return func(win *Window) {
		win.Opts.OnInitPre = value
	}
}

func OptOnInitPost(value spsgio.OnWinInitPost) Option {
	return func(win *Window) {
		win.Opts.OnInitPost = value
	}
}

func OptOnStart(value spsgio.OnStart) Option {
	return func(win *Window) {
		win.Opts.OnStart = value
	}
}

func OptOnLoop(value spsgio.OnWinLoop) Option {
	return func(win *Window) {
		win.Opts.OnLoop = value
	}
}

func OptOnStop(value spsgio.OnStop) Option {
	return func(win *Window) {
		win.Opts.OnStop = value
	}
}

func OptOnEventPre(value spsgio.OnWinEventPre) Option {
	return func(win *Window) {
		win.Opts.OnEventPre = value
	}
}

func OptOnEventPost(value spsgio.OnWinEventPost) Option {
	return func(win *Window) {
		win.Opts.OnEventPost = value
	}
}
