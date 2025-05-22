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

	window, err := glfw.CreateWindow(w.Opts.Width, w.Opts.Height, w.Opts.Title, w.Opts.Monitor, w.Opts.Share)
	if err != nil {
		w.logger.Fatal().Msgf("Failed to create window %v", err)
	}

	window.MakeContextCurrent()

	w.window = window
	return window
}

func (w *Window) GetLogger() *zerolog.Logger {
	return &w.logger
}
