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

func Run(newWindow NewWindow, opts ...Option) {
	if newWindow == nil {
		return
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		defer os.Exit(0)

		a := NewApp(ctx, newWindow, opts...)
		win := newWindow(a)
		if win == nil {
			return
		}
		a.Run(win)
	}()

	app.Main()
}

func NewApp(ctx context.Context, newWindow NewWindow, opts ...Option) *App {
	ctx, cancel := context.WithCancel(ctx)
	a := &App{
		Context:  ctx,
		Shutdown: cancel,

		Pref: spspref.NewPreferences(),

		Opts: NewOptions(newWindow, opts...),

		Logger: spslog.NewLogger(),
	}
	a.Init(opts...)
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

func (a *App) Init(opts ...Option) {
	a.UpdateOpts(opts...)

	if a.Opts.OnCreate != nil {
		a.Opts.OnCreate(a)
	}
}

func (a *App) UpdateOpts(opts ...Option) {
	for _, opt := range opts {
		opt(&a.Opts)
	}
}

func (a *App) Run(win *spswin.Window) {
	// OnStart
	a.Logger.Info().Msg("SPS Gio Hello!!")
	if a.Opts.OnStart != nil {
		a.Opts.OnStart(a)
	}

	a.NewWindow(win)
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

func (a *App) NewWindow(win *spswin.Window) {
	a.GoRun(func() {
		win.Run()
	})
}
