package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/io/event"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsabout "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/about"
	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/pref"
	spssample "github.com/SPSZerone/sps-go-zerone/graphics/gio/sample/page/sample"
)

func main() {
	spsgio.Run(
		spsgio.OptTitle("SPS Tools"),
		//spsgio.OptStartAction(system.ActionMaximize),
		spsgio.OptLoopMode(spsgio.LoopModeSimple),
		//spsgio.OptLoopMode(spsgio.LoopModeParam),
		//spsgio.OptLoopMode(spsgio.LoopModeCustom),

		spsgio.OptOnInitPre(func(app *spsgio.Application) {
			app.Logger.Info().Msg("SPS Tools InitPre")
			//app.Pref.Settings.NonModalDrawer = true
		}),
		spsgio.OptOnInitPost(func(app *spsgio.Application) {
			app.Logger.Info().Msg("SPS Tools InitPost")

			pages := &app.Pages
			pageTag := 0
			pages.Register(pageTag, spsabout.New(pages))

			pageTag++
			pref := spspref.New(pages)
			for i := 0; i < 20; i++ {
				pref.Tabs.AddTabByNames(fmt.Sprintf("test-%d", i))
			}
			pref.Tabs.SetSelected(spspref.TabIdxSettings)
			pages.Register(pageTag, pref)

			pageTag++
			pages.Register(pageTag, spssample.New(app, pages))
		}),
		spsgio.OptOnStart(func(app *spsgio.Application) {
			app.Logger.Info().Msg("SPS Tools Start")
			app.Pages.Start(2)
		}),
		spsgio.OptOnLoop(onLoop),
		spsgio.OptOnStop(func(app *spsgio.Application) {
			app.Logger.Info().Msg("SPS Tools Stop")
		}),
	)
}

func onLoop(a *spsgio.Application) error {
	a.Logger.Info().Msg("SPS Tools Loop")

	chanEvent := make(chan event.Event)
	chanEventDone := make(chan struct{})

	a.GoRun(func() {
		for {
			evt := a.Window.Event()
			chanEvent <- evt
			<-chanEventDone
			if _, ok := evt.(app.DestroyEvent); ok {
				a.Logger.Info().Msg("loopCustom Window.Event app.DestroyEvent ...")
				return
			}
		}
	})

	for {
		select {
		case evt := <-chanEvent:
			switch e := evt.(type) {
			case app.DestroyEvent:
				a.Logger.Info().Msg("loopCustom app.DestroyEvent ...")
				chanEventDone <- struct{}{}
				return e.Err
			case app.FrameEvent:
				a.OnFrameEvent(e, nil)
			}

			chanEventDone <- struct{}{}
		}
	}
}
