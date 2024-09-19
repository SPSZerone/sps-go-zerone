package gio

type Option func(o *Options)

type OnStart func(app *Application)
type OnEnd func(app *Application)
type OnWindowInit func(win *Window)

type Options struct {
	Title        string
	OnStart      OnStart
	OnEnd        OnEnd
	OnWindowInit OnWindowInit
}

func OptTitle(value string) Option {
	return func(o *Options) {
		o.Title = value
	}
}

func OptOnStart(value OnStart) Option {
	return func(o *Options) {
		o.OnStart = value
	}
}

func OptOnEnd(value OnEnd) Option {
	return func(o *Options) {
		o.OnEnd = value
	}
}

func OptOnWindowInit(value OnWindowInit) Option {
	return func(o *Options) {
		o.OnWindowInit = value
	}
}
