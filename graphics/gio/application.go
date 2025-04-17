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
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
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

func NewApplication(ctx context.Context, opts ...Option) *Application {
	ctx, cancel := context.WithCancel(ctx)
	a := &Application{
		Context:  ctx,
		Shutdown: cancel,
		Logger:   spslog.NewLogger(),

		Pref: spspref.NewPreferences(),
		Tabs: spstab.NewTabs(),
	}
	a.Init(opts...)
	return a
}

type LoopMode int

const (
	LoopModeSimple LoopMode = iota
	LoopModeParam
	LoopModeCustom
)

type Application struct {
	Context   context.Context
	Shutdown  func()
	waitGroup sync.WaitGroup

	Pref spspref.Preferences
	Opts Options

	Window *app.Window
	Tabs   spstab.Tabs
	Pages  []Pages

	Ops   op.Ops
	Theme *material.Theme
	Deco  widget.Decorations

	ChanParam chan any

	Logger zerolog.Logger

	startAction system.Action
}

func (a *Application) Init(opts ...Option) {
	// default init
	a.Opts.StartAction = system.ActionMaximize

	for _, opt := range opts {
		opt(&a.Opts)
	}

	if a.Opts.OnInitPre != nil {
		a.Opts.OnInitPre(a)
	}

	a.startAction = a.Opts.StartAction

	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	a.Theme = theme

	a.AddPages(NewPages(a, "Default"))
	a.Window = new(app.Window)

	a.Window.Option(app.Title(a.Opts.Title), app.Decorated(a.Pref.Settings.Decorated.Value))

	if a.Opts.OnInitPost != nil {
		a.Opts.OnInitPost(a)
	}
}

func (a *Application) AddPages(pages ...Pages) {
	for _, page := range pages {
		a.Pages = append(a.Pages, page)
		tab := spstab.Tab{
			Name: page.Name,
			Data: a.GetPages(len(a.Pages) - 1),
		}
		a.Tabs.AddTab(tab)
	}
}

func (a *Application) GetPages(index int) *Pages {
	if index < 0 || index >= len(a.Pages) {
		return nil
	}
	return &a.Pages[index]
}

func (a *Application) CurPages() *Pages {
	return a.GetPages(a.Tabs.GetSelected())
}

func (a *Application) GetDefaultPages() *Pages {
	return a.GetPages(0)
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

	if a.Opts.LoopMode == LoopModeCustom && a.Opts.OnLoop != nil {
		a.Logger.Info().Msg("loopCustom...")
		return a.Opts.OnLoop(a)
	}

	if a.Opts.LoopMode == LoopModeParam {
		return a.loopParam()
	}

	return a.loopSimple()
}

func (a *Application) loopSimple() error {
	a.Logger.Info().Msg("loopSimple...")

	for {
		evt := a.Window.Event()

		a.CurPages().OnEventPre(a, evt, nil)

		switch e := evt.(type) {
		case app.DestroyEvent:
			a.Logger.Info().Msg("loopSimple app.DestroyEvent ...")
			a.CurPages().OnEventPost(a, evt, nil)
			return e.Err
		case app.FrameEvent:
			a.OnFrameEvent(e, nil)
		}

		a.CurPages().OnEventPost(a, evt, nil)
	}
}

func (a *Application) loopParam() error {
	a.Logger.Info().Msg("loopParam...")

	a.ChanParam = make(chan any)
	chanEvent := make(chan event.Event)
	chanEventDone := make(chan struct{})

	a.GoRun(func() {
		for {
			evt := a.Window.Event()
			chanEvent <- evt
			<-chanEventDone
			if _, ok := evt.(app.DestroyEvent); ok {
				a.Logger.Info().Msg("loopParam Window.Event app.DestroyEvent ...")
				return
			}
		}
	})

	var param any
	for {
		select {
		case param = <-a.ChanParam:
			a.Window.Invalidate()
		case evt := <-chanEvent:
			a.CurPages().OnEventPre(a, evt, param)

			switch e := evt.(type) {
			case app.DestroyEvent:
				a.Logger.Info().Msg("loopParam app.DestroyEvent ...")
				a.CurPages().OnEventPost(a, evt, param)
				chanEventDone <- struct{}{}
				return e.Err
			case app.FrameEvent:
				a.OnFrameEvent(e, param)
			}

			a.CurPages().OnEventPost(a, evt, param)
			chanEventDone <- struct{}{}
		}
	}
}

func (a *Application) OnFrameEvent(e app.FrameEvent, param any) {
	gtx := app.NewContext(&a.Ops, e)

	if a.Tabs.Count() == 1 {
		a.Layout(gtx, param)
	}

	e.Frame(gtx.Ops)
}

func (a *Application) Layout(gtx layout.Context, param any) {
	a.CurPages().Layout(a, gtx, param, func() layout.FlexChild {
		if a.startAction != 0 {
			clickable := a.Deco.Clickable(a.startAction)
			if clickable != nil {
				clickable.Click()
			}
			a.startAction = 0
		}

		a.Window.Perform(a.Deco.Update(gtx))
		return a.decorationsFlexChild()
	})
}

func (a *Application) decorationsFlexChild() layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return material.Decorations(a.Theme, &a.Deco, ^system.Action(0), a.Opts.Title).Layout(gtx)
	})
}

func (a *Application) SendParam(param any) {
	if a.ChanParam == nil || param == nil {
		return
	}
	a.ChanParam <- param
}

func (a *Application) GetNavigationRatioLimit() (min, max float32) {
	min = 0.05
	max = 0.75
	return
}
