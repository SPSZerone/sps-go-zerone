package window

import (
	"context"
	"fmt"
	"sync"

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

func NewWindow(ctx context.Context, opts ...Option) *Window {
	ctx, cancel := context.WithCancel(ctx)
	a := &Window{
		Context:  ctx,
		Shutdown: cancel,
		Logger:   spslog.NewLogger(),

		Pref: spspref.NewPreferences(),
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
	Context   context.Context
	Shutdown  func()
	waitGroup sync.WaitGroup

	Pref spspref.Preferences
	Opts Options

	Window *app.Window
	Pages  Pages

	Ops   op.Ops
	Theme *material.Theme
	Deco  widget.Decorations

	ChanParam chan any

	Logger zerolog.Logger

	startAction system.Action
}

func (w *Window) Init(opts ...Option) {
	// default init
	w.Opts.StartAction = system.ActionMaximize

	for _, opt := range opts {
		opt(&w.Opts)
	}

	if w.Opts.OnInitPre != nil {
		w.Opts.OnInitPre(w)
	}

	w.startAction = w.Opts.StartAction

	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	w.Theme = theme

	w.Pages = NewPages(w)
	w.Window = new(app.Window)

	w.Window.Option(app.Title(w.Opts.Title), app.Decorated(w.Pref.Settings.Decorated.Value))

	if w.Opts.OnInitPost != nil {
		w.Opts.OnInitPost(w)
	}
}

func (w *Window) LogPrefix() string {
	return fmt.Sprintf("Window {ID: %v Title: %s}", w.Opts.ID, w.Opts.Title)
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
	w.Pages.Layout(w, gtx, param, func() layout.FlexChild {
		if w.startAction != 0 {
			clickable := w.Deco.Clickable(w.startAction)
			if clickable != nil {
				clickable.Click()
			}
			w.startAction = 0
		}

		w.Window.Perform(w.Deco.Update(gtx))
		return w.decorationsFlexChild()
	})
}

func (w *Window) decorationsFlexChild() layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return material.Decorations(w.Theme, &w.Deco, ^system.Action(0), w.Opts.Title).Layout(gtx)
	})
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
