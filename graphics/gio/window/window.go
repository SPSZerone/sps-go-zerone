package window

import (
	"context"
	"fmt"
	"sync"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

func NewWindow(app spsgio.App, opts ...Option) *Window {
	ctx, cancel := context.WithCancel(app.GetContext())
	a := &Window{
		Context:  ctx,
		Shutdown: cancel,
		Logger:   spslog.NewLogger(),

		Pref: spspref.NewPreferences(),

		App: app,
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

type Window struct {
	ID    any
	Title string

	Context   context.Context
	Shutdown  func()
	waitGroup sync.WaitGroup

	Pref spspref.Preferences
	Opts Options

	App   spsgio.App
	Pages Pages

	Window *app.Window
	Ops    op.Ops
	Theme  *material.Theme
	Deco   widget.Decorations

	ChanParam chan any

	Logger zerolog.Logger

	startAction system.Action
}

func (w *Window) Init(opts ...Option) {
	// default init
	if spsos.IsDarwin() && w.Pref.Settings.Decorated.Value {
		w.Opts.StartAction = system.ActionFullscreen
	} else {
		w.Opts.StartAction = system.ActionMaximize
	}
	w.Update(opts...)

	if w.Opts.OnInitPre != nil {
		w.Opts.OnInitPre(w)
	}

	w.startAction = w.Opts.StartAction

	w.Theme = spsgio.NewTheme()

	w.Pages = NewPages(w)
	w.Window = new(app.Window)

	w.Window.Option(app.Title(w.Title), app.Decorated(w.Pref.Settings.Decorated.Value))

	if w.Opts.OnInitPost != nil {
		w.Opts.OnInitPost(w)
	}
}

func (w *Window) GetId() any {
	return w.ID
}

func (w *Window) GetTitle() string {
	return w.Title
}

func (w *Window) GetContext() context.Context {
	return w.Context
}

func (w *Window) GetPref() *spspref.Preferences {
	return &w.Pref
}

func (w *Window) GetApp() spsgio.App {
	return w.App
}

func (w *Window) GetPages() spsgio.Pages {
	return &w.Pages
}

func (w *Window) GetWindow() *app.Window {
	return w.Window
}

func (w *Window) GetOps() *op.Ops {
	return &w.Ops
}

func (w *Window) GetTheme() *material.Theme {
	return w.Theme
}

func (w *Window) GetDeco() *widget.Decorations {
	return &w.Deco
}

func (w *Window) GetLogger() *zerolog.Logger {
	return &w.Logger
}

func (w *Window) Update(opts ...Option) {
	for _, opt := range opts {
		opt(w)
	}
}

func (w *Window) LogPrefix() string {
	return fmt.Sprintf("Window {ID: %v Title: %s}", w.ID, w.Title)
}

func (w *Window) Run() {
	// OnStart
	w.Logger.Info().Msgf("%s Hello!!", w.LogPrefix())
	if w.Opts.OnStart != nil {
		w.Opts.OnStart(w)
	}

	w.run()

	// OnStop
	if w.Opts.OnStop != nil {
		w.Opts.OnStop(w)
	}
	w.Logger.Info().Msgf("%s Bye!!", w.LogPrefix())
}

func (w *Window) GoRun(run func()) {
	if run == nil {
		return
	}

	w.waitGroup.Add(1)

	go func() {
		defer w.waitGroup.Done()

		run()
	}()
}

func (w *Window) run() {
	w.GoRun(func() {
		if err := w.loop(); err != nil {
			w.Logger.Info().Msgf("%s err: %+v", w.LogPrefix(), err)
		}
	})

	w.waitGroup.Wait()
}

func (w *Window) loop() error {
	go func() {
		<-w.Context.Done()
		w.Logger.Info().Msgf("%s close by signal ...", w.LogPrefix())
		w.Window.Perform(system.ActionClose)
	}()

	if w.Opts.LoopMode == LoopModeCustom && w.Opts.OnLoop != nil {
		w.Logger.Info().Msgf("%s loopCustom...", w.LogPrefix())
		return w.Opts.OnLoop(w)
	}

	if w.Opts.LoopMode == LoopModeParam {
		return w.loopParam()
	}

	return w.loopSimple()
}

func (w *Window) loopSimple() error {
	w.Logger.Info().Msgf("%s loopSimple...", w.LogPrefix())

	for {
		evt := w.Window.Event()

		w.Pages.OnEventPre(w, evt, nil)

		switch e := evt.(type) {
		case app.DestroyEvent:
			w.Logger.Info().Msgf("%s loopSimple app.DestroyEvent ...", w.LogPrefix())
			w.Pages.OnEventPost(w, evt, nil)
			return e.Err
		case app.FrameEvent:
			w.OnFrameEvent(e, nil)
		}

		w.Pages.OnEventPost(w, evt, nil)
	}
}

func (w *Window) loopParam() error {
	w.Logger.Info().Msgf("%s loopParam...", w.LogPrefix())

	w.ChanParam = make(chan any)
	chanEvent := make(chan event.Event)
	chanEventDone := make(chan struct{})

	w.GoRun(func() {
		for {
			evt := w.Window.Event()
			chanEvent <- evt
			<-chanEventDone
			if _, ok := evt.(app.DestroyEvent); ok {
				w.Logger.Info().Msgf("%s loopParam Window.Event app.DestroyEvent ...", w.LogPrefix())
				return
			}
		}
	})

	var param any
	for {
		select {
		case param = <-w.ChanParam:
			w.Window.Invalidate()
		case evt := <-chanEvent:
			w.Pages.OnEventPre(w, evt, param)

			switch e := evt.(type) {
			case app.DestroyEvent:
				w.Logger.Info().Msgf("%s loopParam app.DestroyEvent ...", w.LogPrefix())
				w.Pages.OnEventPost(w, evt, param)
				chanEventDone <- struct{}{}
				return e.Err
			case app.FrameEvent:
				w.OnFrameEvent(e, param)
			}

			w.Pages.OnEventPost(w, evt, param)
			chanEventDone <- struct{}{}
		}
	}
}

func (w *Window) OnFrameEvent(e app.FrameEvent, param any) {
	gtx := app.NewContext(&w.Ops, e)

	w.Layout(gtx, param)

	e.Frame(gtx.Ops)
}

func (w *Window) Layout(gtx layout.Context, param any) {
	w.doStartAction()
	w.Pages.Layout(w, gtx, param, func() layout.FlexChild {
		return layout.Rigid(w.decorationsWidget)
	})
}

func (w *Window) doStartAction() {
	if !w.Pref.Settings.Decorated.Value {
		if w.startAction != DecoActionNone {
			clickable := w.Deco.Clickable(w.startAction)
			if clickable != nil {
				clickable.Click()
			}
			w.startAction = DecoActionNone
		}
	} else {
		if w.startAction != DecoActionNone {
			w.Window.Perform(w.startAction)
			w.startAction = DecoActionNone
		}
	}
}

func (w *Window) decorationsWidget(gtx layout.Context) layout.Dimensions {
	w.Window.Perform(w.Deco.Update(gtx))
	return material.Decorations(w.Theme, &w.Deco, ^system.Action(0), w.Title).Layout(gtx)
}

func (w *Window) SendParam(param any) {
	if w.ChanParam == nil || param == nil {
		return
	}
	w.ChanParam <- param
}

func (w *Window) GetNavigationRatioLimit() (min, max float32) {
	min = 0.05
	max = 0.75
	return
}
