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
		spswin.OptTitle("SPS Tools"),
		//spswin.OptStartAction(system.ActionMaximize),
		spswin.OptLoopMode(spswin.LoopModeSimple),
		//spswin.OptLoopMode(spswin.LoopModeParam),
		//spswin.OptLoopMode(spswin.LoopModeCustom),

		spswin.OptOnInitPre(func(win *spswin.Window) {
			win.Logger.Info().Msg("SPS Tools InitPre")
		}),
		spswin.OptOnInitPost(func(win *spswin.Window) {
			win.Logger.Info().Msg("SPS Tools InitPost")

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
			win.Logger.Info().Msg("SPS Tools Start")
			win.Pages.Start(2)
		}),
		spswin.OptOnLoop(onLoop),
		spswin.OptOnStop(func(win *spswin.Window) {
			win.Logger.Info().Msg("SPS Tools Stop")
		}),
	)
}

func onLoop(win *spswin.Window) error {
	win.Logger.Info().Msg("SPS Tools Loop")

	chanEvent := make(chan event.Event)
	chanEventDone := make(chan struct{})

	win.GoRun(func() {
		for {
			evt := win.Window.Event()
			chanEvent <- evt
			<-chanEventDone
			if _, ok := evt.(app.DestroyEvent); ok {
				win.Logger.Info().Msg("loopCustom Window.Event app.DestroyEvent ...")
				return
			}
		}
	})

	for {
		select {
		case evt := <-chanEvent:
			switch e := evt.(type) {
			case app.DestroyEvent:
				win.Logger.Info().Msg("loopCustom app.DestroyEvent ...")
				chanEventDone <- struct{}{}
				return e.Err
			case app.FrameEvent:
				win.OnFrameEvent(e, nil)
			}

			chanEventDone <- struct{}{}
		}
	}
}
