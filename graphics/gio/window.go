package gio

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/page"
	pageabout "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/about"
	pagepref "github.com/SPSZerone/sps-go-zerone/graphics/gio/page/pref"
)

type Window struct {
	App    *Application
	Window *app.Window

	Title string

	Pages *page.Pages

	ops   op.Ops
	theme *material.Theme
	deco  widget.Decorations
}

func NewWindow(a *Application, title string, pages *page.Pages, opts ...app.Option) *Window {
	w := &Window{
		App:    a,
		Window: new(app.Window),
	}
	w.Init(title, pages, opts...)

	if a.Opts.OnWindowInit == nil {
		pages.Register(0, pageabout.New(pages))
		pages.Register(1, pagepref.New(pages))
	} else {
		a.Opts.OnWindowInit(w)
	}

	return w
}

func (w *Window) Init(title string, pages *page.Pages, opts ...app.Option) {
	w.Title = title

	w.Pages = pages

	opts = append(opts, app.Title(title))
	w.Window.Option(opts...)

	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	w.theme = th
}

func (w *Window) Run(opts ...app.Option) error {
	w.Window.Option(opts...)

	go func() {
		<-w.App.Context.Done()
		w.App.Logger.Info().Msg("close by signal ...")
		w.Window.Perform(system.ActionClose)
	}()

	for {
		switch e := w.Window.Event().(type) {
		case app.DestroyEvent:
			w.App.Logger.Info().Msg("app.DestroyEvent ...")
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&w.ops, e)

			w.Pages.Layout(gtx, w.Window, w.theme, func() layout.FlexChild {
				w.Window.Perform(w.deco.Update(gtx))
				return w.decorationsFlexChild()
			})

			e.Frame(gtx.Ops)
		}
	}
}

func (w *Window) decorationsFlexChild() layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return material.Decorations(w.theme, &w.deco, ^system.Action(0), w.Title).Layout(gtx)
	})
}
