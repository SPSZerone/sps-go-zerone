package pref

import (
	"gioui.org/layout"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func (p *Page) LayoutSettings(app *spsgio.Application, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsNavigation(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsNonModalDrawer(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsTabAxis(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsValueInFront(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsBottomBar(app, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsDecorated(app, gtx)
		}),
	}
}

func (p *Page) LayoutSettingsNavigation(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if app.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return app.Pref.Settings.Navigation.Layout(
		app.Window, app.Theme, gtx,
		app.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsNonModalDrawer(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if app.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return app.Pref.Settings.ModalNavDrawer.Layout(
		app.Window, app.Theme, gtx,
		app.Pref.Settings.ValueInFront.Value, ratio,
		&p.Pages.NavAnim,
	)
}

func (p *Page) LayoutSettingsTabAxis(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if app.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return app.Pref.Settings.TabAxis.Layout(
		app.Window, app.Theme, gtx,
		app.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsValueInFront(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if app.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return app.Pref.Settings.ValueInFront.Layout(
		app.Window, app.Theme, gtx,
		app.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsBottomBar(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if app.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return app.Pref.Settings.BottomBar.Layout(
		app.Window, app.Theme, gtx,
		app.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsDecorated(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if app.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return app.Pref.Settings.Decorated.Layout(
		app.Window, app.Theme, gtx,
		app.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) Navigation(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	return app.Pref.Settings.Navigation.LayoutSwitch(
		app.Window, app.Theme, gtx,
	)
}

func (p *Page) NonModalDrawer(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	return app.Pref.Settings.ModalNavDrawer.LayoutSwitch(
		app.Window, app.Theme, gtx,
		&p.Pages.NavAnim,
	)
}

func (p *Page) TabAxis(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	return app.Pref.Settings.TabAxis.LayoutSwitch(
		app.Window, app.Theme, gtx,
	)
}

func (p *Page) ValueInFront(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	return app.Pref.Settings.ValueInFront.LayoutSwitch(
		app.Window, app.Theme, gtx,
	)
}

func (p *Page) BottomBar(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	return app.Pref.Settings.BottomBar.LayoutSwitch(
		app.Window, app.Theme, gtx,
		func() {
			if app.Pref.Settings.BottomBar.Value {
				p.Pages.ModalNavDrawer.Anchor = component.Bottom
				p.Pages.AppBar.Anchor = component.Bottom
			} else {
				p.Pages.ModalNavDrawer.Anchor = component.Top
				p.Pages.AppBar.Anchor = component.Top
			}
		},
	)
}

func (p *Page) Decorated(app *spsgio.Application, gtx layout.Context) layout.Dimensions {
	return app.Pref.Settings.Decorated.LayoutSwitch(
		app.Window, app.Theme, gtx,
	)
}
