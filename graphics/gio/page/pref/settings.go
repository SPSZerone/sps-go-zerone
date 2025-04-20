package pref

import (
	"gioui.org/layout"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func (p *Page) LayoutSettings(win spsgio.Window, gtx layout.Context, param any) []layout.FlexChild {
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

func (p *Page) LayoutSettingsNonModalDrawer(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	var ratio float32
	if pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return pref.Settings.ModalNavDrawer.Layout(
		window, theme, gtx,
		pref.Settings.ValueInFront.Value, ratio,
		p.Pages.GetNavAnim(),
	)
}

func (p *Page) LayoutSettingsTabAxis(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	var ratio float32
	if pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return pref.Settings.TabAxis.Layout(
		window, theme, gtx,
		pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsValueInFront(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	var ratio float32
	if pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return pref.Settings.ValueInFront.Layout(
		window, theme, gtx,
		pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsBottomBar(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	var ratio float32
	if pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return pref.Settings.BottomBar.Layout(
		window, theme, gtx,
		pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) LayoutSettingsDecorated(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	var ratio float32
	if pref.Settings.ValueInFront.Value {
		ratio = RatioValue
	} else {
		ratio = RatioKey
	}
	return pref.Settings.Decorated.Layout(
		window, theme, gtx,
		pref.Settings.ValueInFront.Value, ratio,
	)
}

func (p *Page) NonModalDrawer(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	return pref.Settings.ModalNavDrawer.LayoutSwitch(
		window, theme, gtx,
		p.Pages.GetNavAnim(),
	)
}

func (p *Page) TabAxis(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	return pref.Settings.TabAxis.LayoutSwitch(
		window, theme, gtx,
	)
}

func (p *Page) ValueInFront(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	return pref.Settings.ValueInFront.LayoutSwitch(
		window, theme, gtx,
	)
}

func (p *Page) BottomBar(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	return pref.Settings.BottomBar.LayoutSwitch(
		window, theme, gtx,
		func() {
			if pref.Settings.BottomBar.Value {
				p.Pages.GetModalNavDrawer().Anchor = component.Bottom
				p.Pages.GetAppBar().Anchor = component.Bottom
			} else {
				p.Pages.GetModalNavDrawer().Anchor = component.Top
				p.Pages.GetAppBar().Anchor = component.Top
			}
		},
	)
}

func (p *Page) Decorated(win spsgio.Window, gtx layout.Context) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	return pref.Settings.Decorated.LayoutSwitch(
		window, theme, gtx,
	)
}
