package gio

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
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
	Context   context.Context
	Shutdown  func()
	waitGroup sync.WaitGroup

	Pref spspref.Preferences
	Opts Options

	Window *app.Window
	Tabs   Tabs

	Ops   op.Ops
	Theme *material.Theme
	Deco  widget.Decorations

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

	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	a.Theme = th

	a.Tabs = NewTabs()
	a.Window = new(app.Window)

	a.Window.Option(app.Title(a.Opts.Title), app.Decorated(a.Pref.Settings.Decorated))

	if a.Opts.OnInit != nil {
		a.Opts.OnInit(a)
	}
}

func (a *Application) Run() {
	// OnStart
	a.Logger.Info().Msg("Hello!!")
	if a.Opts.OnStart != nil {
		a.Opts.OnStart(a)
	}

	a.run()

	// OnStop
	if a.Opts.OnStop != nil {
		a.Opts.OnStop(a)
	}
	a.Logger.Info().Msg("Bye!!")
}

func (a *Application) GoRun(run func()) {
	if run == nil {
		return
	}

	a.waitGroup.Add(1)

	go func() {
		defer a.waitGroup.Done()

		run()
	}()
}

func (a *Application) run() {
	a.GoRun(func() {
		if err := a.loop(); err != nil {
			a.Logger.Info().Msgf("App %s err: %+v", a.Opts.Title, err)
		}
	})

	a.waitGroup.Wait()
}

func (a *Application) loop() error {
	go func() {
		<-a.Context.Done()
		a.Logger.Info().Msg("close by signal ...")
		a.Window.Perform(system.ActionClose)
	}()

	if a.Opts.OnLoop != nil {
		return a.Opts.OnLoop(a)
	}

	for {
		destroy, err := a.OnEvent(a.Window.Event())
		if destroy {
			return err
		}
	}
}

func (a *Application) OnEvent(evt event.Event) (destroy bool, err error) {
	switch e := evt.(type) {
	case app.DestroyEvent:
		a.Logger.Info().Msg("app.DestroyEvent ...")
		return true, e.Err
	case app.FrameEvent:
		a.OnFrameEvent(e)
	}
	return
}

func (a *Application) OnFrameEvent(e app.FrameEvent) {
	gtx := app.NewContext(&a.Ops, e)

	a.Tabs.Layout(a, gtx, func() layout.FlexChild {
		a.Window.Perform(a.Deco.Update(gtx))
		return a.decorationsFlexChild()
	})

	e.Frame(gtx.Ops)
}

func (a *Application) decorationsFlexChild() layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return material.Decorations(a.Theme, &a.Deco, ^system.Action(0), a.Opts.Title).Layout(gtx)
	})
}
