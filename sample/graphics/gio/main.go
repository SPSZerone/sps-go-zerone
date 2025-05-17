package main

import (
	"gioui.org/app"
	"gioui.org/io/event"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsapp "github.com/SPSZerone/sps-go-zerone/graphics/gio/app"
	spsabout "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/about"
	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/pref"
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
	spssample "github.com/SPSZerone/sps-go-zerone/sample/graphics/gio/page/sample"
)

const (
	Name = "SPS Gio Sample"
)

func main() {
	spsapp.Run(
		NewWindow,
		spsapp.OptOnCreate(func(app spsgio.App) {
			app.GetLogger().Info().Msgf("Window %s Create", Name)
		}),
		spsapp.OptOnStart(func(app spsgio.App) {
			app.GetLogger().Info().Msgf("Window %s Start", Name)
		}),
		spsapp.OptOnStop(func(app spsgio.App) {
			app.GetLogger().Info().Msgf("Window %s Stop", Name)
		}),
	)
}

func NewWindow(app spsgio.App, fromWin spsgio.Window) spsgio.Window {
	pref := app.GetPref()
	if fromWin != nil {
		pref = fromWin.GetPref()
	}
	return spswin.NewWindow(
		app,
		spswin.OptID("Main"),
		spswin.OptTitle(Name),
		spswin.OptPref(*pref),
		//spswin.OptStartAction(system.ActionMaximize),
		spswin.OptLoopMode(spswin.LoopModeSimple),
		//spswin.OptLoopMode(spswin.LoopModeParam),
		//spswin.OptLoopMode(spswin.LoopModeCustom),

		spswin.OptOnInitPre(func(win spsgio.Window) {
			win.GetLogger().Info().Msgf("%s InitPre", win.LogPrefix())
		}),
		spswin.OptOnInitPost(func(win spsgio.Window) {
			win.GetLogger().Info().Msgf("%s InitPost", win.LogPrefix())

			pages := win.GetPages()
			pageTag := 0
			pages.Register(pageTag, spsabout.New(pages))

			pageTag++
			pagePref := spspref.New(pages, app, NewWindow)
			pagePref.Tabs.SetSelected(spspref.TabIdxSettings)
			pages.Register(pageTag, pagePref)

			pageTag++
			pages.Register(pageTag, spssample.New(win.GetTheme(), pages))
		}),
		spswin.OptOnStart(func(win spsgio.Window) {
			win.GetLogger().Info().Msgf("%s Start", win.LogPrefix())
			win.GetPages().Start(2)
		}),
		spswin.OptOnLoop(onLoop),
		spswin.OptOnStop(func(win spsgio.Window) {
			win.GetLogger().Info().Msgf("%s Stop", win.LogPrefix())
		}),
	)
}

func onLoop(win spsgio.Window) error {
	win.GetLogger().Info().Msgf("%s Loop", win.LogPrefix())

	chanEvent := make(chan event.Event)
	chanEventDone := make(chan struct{})

	win.GoRun(func() {
		for {
			evt := win.GetWindow().Event()
			chanEvent <- evt
			<-chanEventDone
			if _, ok := evt.(app.DestroyEvent); ok {
				win.GetLogger().Info().Msgf("%s Loop loopCustom Window.Event app.DestroyEvent ...", win.LogPrefix())
				return
			}
		}
	})

	for {
		select {
		case evt := <-chanEvent:
			switch e := evt.(type) {
			case app.DestroyEvent:
				win.GetLogger().Info().Msgf("%s loopCustom app.DestroyEvent ...", win.LogPrefix())
				chanEventDone <- struct{}{}
				return e.Err
			case app.FrameEvent:
				win.OnFrameEvent(e, nil)
			}

			chanEventDone <- struct{}{}
		}
	}
}
