package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/rs/zerolog"

	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

func NewWindow(opts ...Option) *Window {
	win := &Window{
		Opts:   NewOptions(opts...),
		logger: spslog.NewLogger(),
	}
	win.CreateWindow()
	return win
}

type Window struct {
	Opts   Options
	window *glfw.Window
	logger zerolog.Logger
}

func (w *Window) GetWindow() *glfw.Window {
	return w.window
}

func (w *Window) CreateWindow() *glfw.Window {
	if w.window != nil {
		return w.window
	}

	window, err := DoCreateWindow(w.Opts)
	if err != nil {
		w.logger.Fatal().Msgf("Failed to create window %v", err)
	}

	w.window = window
	return window
}

func (w *Window) GetLogger() *zerolog.Logger {
	return &w.logger
}

func CreateWindow(opts ...Option) (*glfw.Window, error) {
	o := NewOptions(opts...)
	return DoCreateWindow(o)
}

func DoCreateWindow(o Options) (*glfw.Window, error) {
	window, err := glfw.CreateWindow(o.Width, o.Height, o.Title, o.Monitor, o.Share)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(o.FramebufferSizeCallback)

	return window, nil
}
