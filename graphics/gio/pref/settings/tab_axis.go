package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewTabAxis() TabAxis {
	return TabAxis{
		Bool: setting.Bool{
			Name: "TabAxis Horizontal",
			Desc: "TabAxis Horizontal",
		},
	}
}

type TabAxis struct {
	setting.Bool
}

func (d *TabAxis) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return d.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (d *TabAxis) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return d.Bool.Layout(theme, gtx, valueInFront, ratioInFront, nil)
}

func (d *TabAxis) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return d.Bool.LayoutSwitch(theme, gtx, nil)
}
