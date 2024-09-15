package gio

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gioui.org/app"
	"github.com/rs/zerolog"

	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/page"
	pageabout "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/about"
	pagepref "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/pref"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

func Run(opts ...Option) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		a := NewApplication(ctx, opts...)
		a.Run()
		os.Exit(0)
	}()

	app.Main()
}

type Application struct {
	Context  context.Context
	Shutdown func()
	active   sync.WaitGroup

	Pref pref.Preferences
	Opts Options

	Logger zerolog.Logger
}

func NewApplication(ctx context.Context, opts ...Option) *Application {
	ctx, cancel := context.WithCancel(ctx)
	a := &Application{
		Context:  ctx,
		Shutdown: cancel,
		Logger:   spslog.NewLogger(),
	}
	a.Init(opts...)
	return a
}

func (a *Application) Init(opts ...Option) {
	for _, opt := range opts {
		opt(&a.Opts)
	}
}

func (a *Application) Run() {
	a.Logger.Info().Msg("Hello!!")
	if a.Opts.OnStart != nil {
		a.Opts.OnStart(a)
	}

	pages := page.NewPages(&a.Pref)

	if a.Opts.RegisterPage == nil {
		pages.Register(0, pageabout.New(&pages))
		pages.Register(1, pagepref.New(&pages))
	} else {
		a.Opts.RegisterPage(&pages)
	}

	a.NewWindow(a.Opts.Title, &pages, app.Decorated(a.Pref.Settings.Decorated))
	a.Wait()

	if a.Opts.OnEnd != nil {
		a.Opts.OnEnd(a)
	}
	a.Logger.Info().Msg("Bye!!")
}

func (a *Application) Wait() {
	a.active.Wait()
}

func (a *Application) NewWindow(title string, pages *page.Pages, opts ...app.Option) {
	opts = append(opts, app.Title(title))
	a.active.Add(1)

	go func() {
		defer a.active.Done()

		w := NewWindow(a, title, pages, opts...)

		if err := w.Run(); err != nil {
			a.Logger.Info().Msgf("window %s err: %+v", title, err)
		}
	}()
}
