package gio

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gioui.org/app"
	"gioui.org/font/gofont"
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
	Context  context.Context
	Shutdown func()
	active   sync.WaitGroup

	Pref spspref.Preferences
	Opts Options

	Window *app.Window
	Pages  Pages

	ops   op.Ops
	theme *material.Theme
	deco  widget.Decorations

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
	a.theme = th

	a.Pages = NewPages()
	a.Window = new(app.Window)

	a.Window.Option(app.Title(a.Opts.Title), app.Decorated(a.Pref.Settings.Decorated))
}

func (a *Application) Run() {
	// OnStart
	a.Logger.Info().Msg("Hello!!")
	if a.Opts.OnStart != nil {
		a.Opts.OnStart(a)
	}

	a.run()

	a.wait()

	// OnEnd
	if a.Opts.OnEnd != nil {
		a.Opts.OnEnd(a)
	}
	a.Logger.Info().Msg("Bye!!")
}

func (a *Application) wait() {
	a.active.Wait()
}

func (a *Application) run() {
	a.active.Add(1)

	go func() {
		defer a.active.Done()

		if a.Opts.OnWindowInit != nil {
			a.Opts.OnWindowInit(a)
		}

		if err := a.runLogic(); err != nil {
			a.Logger.Info().Msgf("window %s err: %+v", a.Opts.Title, err)
		}
	}()
}

func (a *Application) runLogic() error {
	go func() {
		<-a.Context.Done()
		a.Logger.Info().Msg("close by signal ...")
		a.Window.Perform(system.ActionClose)
	}()

	for {
		switch e := a.Window.Event().(type) {
		case app.DestroyEvent:
			a.Logger.Info().Msg("app.DestroyEvent ...")
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&a.ops, e)

			a.Pages.Layout(a, gtx, a.Window, a.theme, func() layout.FlexChild {
				a.Window.Perform(a.deco.Update(gtx))
				return a.decorationsFlexChild()
			})

			e.Frame(gtx.Ops)
		}
	}
}

func (a *Application) decorationsFlexChild() layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return material.Decorations(a.theme, &a.deco, ^system.Action(0), a.Opts.Title).Layout(gtx)
	})
}
