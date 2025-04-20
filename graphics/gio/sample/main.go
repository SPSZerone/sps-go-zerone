package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/io/event"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsabout "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/about"
	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/pref"
	spssample "github.com/SPSZerone/sps-go-zerone/graphics/gio/sample/page/sample"
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

func main() {
	spsgio.Run(
		spsgio.OptOnCreate(func(app *spsgio.App) {
			app.Logger.Info().Msg("Window SPS Sample Create")
		}),
		spsgio.OptOnStart(func(app *spsgio.App) {
			app.Logger.Info().Msg("Window SPS Sample Start")
		}),
		spsgio.OptOnStop(func(app *spsgio.App) {
			app.Logger.Info().Msg("Window SPS Sample Stop")
		}),
		spsgio.OptWinOpts(
			spswin.OptID("Main"),
			spswin.OptTitle("SPS Sample"),
			//spswin.OptStartAction(system.ActionMaximize),
			spswin.OptLoopMode(spswin.LoopModeSimple),
			//spswin.OptLoopMode(spswin.LoopModeParam),
			//spswin.OptLoopMode(spswin.LoopModeCustom),

			spswin.OptOnInitPre(func(win *spswin.Window) {
				win.Logger.Info().Msgf("%s InitPre", win.LogPrefix())
			}),
			spswin.OptOnInitPost(func(win *spswin.Window) {
				win.Logger.Info().Msgf("%s InitPost", win.LogPrefix())

				pages := &win.Pages
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
				pages.Register(pageTag, spssample.New(win.Theme, pages))
			}),
			spswin.OptOnStart(func(win *spswin.Window) {
				win.Logger.Info().Msgf("%s Start", win.LogPrefix())
				win.Pages.Start(2)
			}),
			spswin.OptOnLoop(onLoop),
			spswin.OptOnStop(func(win *spswin.Window) {
				win.Logger.Info().Msgf("%s Stop", win.LogPrefix())
			}),
		),
	)
}

func onLoop(win *spswin.Window) error {
	win.Logger.Info().Msgf("%s Loop", win.LogPrefix())

	chanEvent := make(chan event.Event)
	chanEventDone := make(chan struct{})

	win.GoRun(func() {
		for {
			evt := win.Window.Event()
			chanEvent <- evt
			<-chanEventDone
			if _, ok := evt.(app.DestroyEvent); ok {
				win.Logger.Info().Msgf("%s Loop loopCustom Window.Event app.DestroyEvent ...", win.LogPrefix())
				return
			}
		}
	})

	for {
		select {
		case evt := <-chanEvent:
			switch e := evt.(type) {
			case app.DestroyEvent:
				win.Logger.Info().Msgf("%s loopCustom app.DestroyEvent ...", win.LogPrefix())
				chanEventDone <- struct{}{}
				return e.Err
			case app.FrameEvent:
				win.OnFrameEvent(e, nil)
			}

			chanEventDone <- struct{}{}
		}
	}
}
