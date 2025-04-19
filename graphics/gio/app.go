package gio

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gioui.org/app"
	"github.com/rs/zerolog"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

func Run(opts ...Option) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		a := NewApp(ctx, opts...)
		a.Run(a.Opts.WinOpts...)
		os.Exit(0)
	}()

	app.Main()
}

func NewApp(ctx context.Context, opts ...Option) *App {
	ctx, cancel := context.WithCancel(ctx)
	a := &App{
		Context:  ctx,
		Shutdown: cancel,

		Pref: spspref.NewPreferences(),

		Logger: spslog.NewLogger(),
	}
	a.UpdateOpts(opts...)
	return a
}

type App struct {
	Context   context.Context
	Shutdown  func()
	waitGroup sync.WaitGroup

	Pref spspref.Preferences
	Opts Options

	Logger zerolog.Logger
}

func (a *App) UpdateOpts(opts ...Option) {
	for _, opt := range opts {
		opt(&a.Opts)
	}
}

func (a *App) Run(opts ...spswin.Option) {
	// OnStart
	a.Logger.Info().Msg("SPS Gio Hello!!")
	if a.Opts.OnStart != nil {
		a.Opts.OnStart(a)
	}

	a.NewWindow(opts...)
	a.waitGroup.Wait()

	// OnStop
	if a.Opts.OnStop != nil {
		a.Opts.OnStop(a)
	}
	a.Logger.Info().Msg("SPS Gio Bye!!")
}

func (a *App) GoRun(run func()) {
	if run == nil {
		return
	}

	a.waitGroup.Add(1)

	go func() {
		defer a.waitGroup.Done()

		run()
	}()
}

func (a *App) NewWindow(opts ...spswin.Option) {
	a.GoRun(func() {
		a := spswin.NewWindow(a.Context, opts...)
		a.Run()
	})
}
