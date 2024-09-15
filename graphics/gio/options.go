package gio

import "github.com/SPSZerone/sps-go-zerone/graphics/gio/page"

type Option func(o *Options)

type RegisterPage func(pages *page.Pages)
type OnStart func(app *Application)
type OnEnd func(app *Application)

type Options struct {
	Title        string
	OnStart      OnStart
	OnEnd        OnEnd
	RegisterPage RegisterPage
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

func OptRegisterPage(value RegisterPage) Option {
	return func(o *Options) {
		o.RegisterPage = value
	}
}
