package gio

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"gioui.org/app"
	"github.com/rs/zerolog"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

func Run(opts ...Option) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		a := NewWindow(ctx, opts...)
		a.Run()
		os.Exit(0)
	}()

	app.Main()
}

func NewApp(ctx context.Context) *App {
	ctx, cancel := context.WithCancel(ctx)
	a := &App{
		Context:  ctx,
		Shutdown: cancel,

		Pref: spspref.NewPreferences(),

		Logger: spslog.NewLogger(),
	}
	return a
}

type App struct {
	Context  context.Context
	Shutdown func()

	Pref spspref.Preferences

	Logger zerolog.Logger
}
