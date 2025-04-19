package pref

import (
	"gioui.org/layout"
	"gioui.org/x/component"

	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

func (p *Page) LayoutSettings(win *spswin.Window, gtx layout.Context, param any) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsNonModalDrawer(win, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsTabAxis(win, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsValueInFront(win, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsBottomBar(win, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.LayoutSettingsDecorated(win, gtx)
		}),
	}
}

func (p *Page) LayoutSettingsNonModalDrawer(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if win.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return win.Pref.Settings.ModalNavDrawer.Layout(
		win.Window, win.Theme, gtx,
		win.Pref.Settings.ValueInFront.Value, ratio,
		&p.Pages.NavAnim,
	)
}

func (p *Page) LayoutSettingsTabAxis(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if win.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return win.Pref.Settings.TabAxis.Layout(
		win.Window, win.Theme, gtx,
		win.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsValueInFront(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if win.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return win.Pref.Settings.ValueInFront.Layout(
		win.Window, win.Theme, gtx,
		win.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsBottomBar(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if win.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return win.Pref.Settings.BottomBar.Layout(
		win.Window, win.Theme, gtx,
		win.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsDecorated(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	var ratio float32
	if win.Pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return win.Pref.Settings.Decorated.Layout(
		win.Window, win.Theme, gtx,
		win.Pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) NonModalDrawer(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	return win.Pref.Settings.ModalNavDrawer.LayoutSwitch(
		win.Window, win.Theme, gtx,
		&p.Pages.NavAnim,
	)
}

func (p *Page) TabAxis(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	return win.Pref.Settings.TabAxis.LayoutSwitch(
		win.Window, win.Theme, gtx,
	)
}

func (p *Page) ValueInFront(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	return win.Pref.Settings.ValueInFront.LayoutSwitch(
		win.Window, win.Theme, gtx,
	)
}

func (p *Page) BottomBar(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	return win.Pref.Settings.BottomBar.LayoutSwitch(
		win.Window, win.Theme, gtx,
		func() {
			if win.Pref.Settings.BottomBar.Value {
				p.Pages.ModalNavDrawer.Anchor = component.Bottom
				p.Pages.AppBar.Anchor = component.Bottom
			} else {
				p.Pages.ModalNavDrawer.Anchor = component.Top
				p.Pages.AppBar.Anchor = component.Top
			}
		},
	)
}

func (p *Page) Decorated(win *spswin.Window, gtx layout.Context) layout.Dimensions {
	return win.Pref.Settings.Decorated.LayoutSwitch(
		win.Window, win.Theme, gtx,
	)
}
