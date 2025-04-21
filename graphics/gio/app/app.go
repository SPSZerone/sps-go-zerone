package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gioui.org/app"
	"github.com/rs/zerolog"

	spsslice "github.com/SPSZerone/sps-go-zerone/generic/slice"
	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

func Run(newWindow spsgio.NewWindow, opts ...Option) {
	if newWindow == nil {
		return
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		defer os.Exit(0)

		a := NewApp(ctx, newWindow, opts...)
		win := newWindow(a, nil)
		if win == nil {
			return
		}
		a.Run(win)
	}()

	app.Main()
}

func NewApp(ctx context.Context, newWindow spsgio.NewWindow, opts ...Option) *App {
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

	Windows []spsgio.Window
	WinLock sync.RWMutex

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

func (a *App) GetContext() context.Context {
	return a.Context
}

func (a *App) GetPref() *spspref.Preferences {
	return &a.Pref
}

func (a *App) Run(win spsgio.Window) {
	// OnAppStart
	a.Logger.Info().Msg("SPS Gio Hello!!")
	if a.Opts.OnStart != nil {
		a.Opts.OnStart(a)
	}

	a.RunWindow(win)
	a.waitGroup.Wait()

	// OnAppStop
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

func (a *App) RunWindow(win spsgio.Window) {
	a.AddWindows(win)

	a.GoRun(func() {
		defer a.DelWindows(win)

		win.Run()
	})
}

func (a *App) GetWindows() []spsgio.Window {
	return a.Windows
}

func (a *App) AddWindows(win spsgio.Window) []spsgio.Window {
	a.WinLock.Lock()
	defer a.WinLock.Unlock()

	a.Windows = append(a.Windows, win)
	return a.Windows
}

func (a *App) DelWindows(win spsgio.Window) []spsgio.Window {
	a.WinLock.Lock()
	defer a.WinLock.Unlock()

	rmIdx := -1
	for i, w := range a.Windows {
		if w != win {
			continue
		}
		rmIdx = i
		break
	}

	a.Windows = spsslice.RemoveFast(a.Windows, rmIdx)
	return a.Windows
}

func (a *App) GetLogger() *zerolog.Logger {
	return &a.Logger
}
