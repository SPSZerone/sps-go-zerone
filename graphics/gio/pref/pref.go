package pref

import (
	gioapp "gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/settings"
)

func NewPreferences() Preferences {
	return Preferences{
		Settings: NewSettings(),
	}
}

func NewSettings() Settings {
	return Settings{
		Decorated: settings.NewDecorated(),
	}
}

type Preferences struct {
	Settings Settings
}

type Settings struct {
	Decorated      settings.Decorated
	NonModalDrawer bool
	BottomBar      bool
	PrefTableStyle bool
	ValueInFront   bool
}

func (p *Settings) DecoratedFlexChild(window *gioapp.Window, theme *material.Theme, gtx layout.Context) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.DecoratedLayout(window, theme, gtx)
		}),
	}
}

func (p *Settings) DecoratedLayout(window *gioapp.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	valueInFront := p.ValueInFront
	return p.Decorated.Layout(
		theme, gtx,
		valueInFront, 0.3,
		func() {
			window.Option(gioapp.Decorated(p.Decorated.Value))
		},
	)
}

func (p *Settings) DecoratedLayoutSwitch(window *gioapp.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return p.Decorated.LayoutSwitch(
		theme, gtx,
		func() {
			window.Option(gioapp.Decorated(p.Decorated.Value))
		},
	)
}
