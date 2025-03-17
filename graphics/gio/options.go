package gio

type Option func(o *Options)

type OnInit func(app *Application)
type OnStart func(app *Application)
type OnLoop func(app *Application) error
type OnStop func(app *Application)

type Options struct {
	LoopMode LoopMode
	Title    string
	OnInit   OnInit
	OnStart  OnStart
	OnLoop   OnLoop
	OnStop   OnStop
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

func OptOnInit(value OnInit) Option {
	return func(o *Options) {
		o.OnInit = value
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
