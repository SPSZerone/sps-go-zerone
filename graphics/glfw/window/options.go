package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	spsglfw "github.com/SPSZerone/sps-go-zerone/graphics/glfw"
)

func NewOptions(opts ...Option) Options {
	o := Options{
		Width:  1920,
		Height: 1080,
		Title:  "GLFW",

		FramebufferSizeCallback: spsglfw.FramebufferSizeCallback,
	}
	o.UpdateOpts(opts...)
	return o
}

type Option func(*Options)

type Options struct {
	Width  int
	Height int
	Title  string

	Monitor *glfw.Monitor
	Share   *glfw.Window

	FramebufferSizeCallback glfw.FramebufferSizeCallback
	CursorPosCallback       glfw.CursorPosCallback
	ScrollCallback          glfw.ScrollCallback
}

func (o *Options) UpdateOpts(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func OptWidth(value int) Option {
	return func(o *Options) {
		o.Width = value
	}
}

func OptHeight(value int) Option {
	return func(o *Options) {
		o.Height = value
	}
}

func OptTitle(value string) Option {
	return func(o *Options) {
		o.Title = value
	}
}

func OptMonitor(value *glfw.Monitor) Option {
	return func(o *Options) {
		o.Monitor = value
	}
}

func OptShare(value *glfw.Window) Option {
	return func(o *Options) {
		o.Share = value
	}
}

func OptFramebufferSizeCallback(value glfw.FramebufferSizeCallback) Option {
	return func(o *Options) {
		o.FramebufferSizeCallback = value
	}
}

func OptCursorPosCallback(value glfw.CursorPosCallback) Option {
	return func(o *Options) {
		o.CursorPosCallback = value
	}
}

func OptScrollCallback(value glfw.ScrollCallback) Option {
	return func(o *Options) {
		o.ScrollCallback = value
	}
}
