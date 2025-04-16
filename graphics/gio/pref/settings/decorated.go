package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewDecorated() Decorated {
	return Decorated{
		Bool: setting.Bool{
			Name: "Decorated",
			Desc: "Use decorated",
		},
	}
}

type Decorated struct {
	setting.Bool
}

func (d *Decorated) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return d.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (d *Decorated) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return d.Bool.Layout(
		theme, gtx,
		valueInFront, ratioInFront,
		func() {
			d.UpdateWindow(window)
		},
	)
}

func (d *Decorated) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return d.Bool.LayoutSwitch(
		theme, gtx,
		func() {
			d.UpdateWindow(window)
		},
	)
}

func (d *Decorated) UpdateWindow(window *app.Window) {
	window.Option(app.Decorated(d.Bool.Value))
}
