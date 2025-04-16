package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewTabAxis() TabAxis {
	return TabAxis{
		Bool: setting.Bool{
			Name: "TabAxis Horizontal",
			Desc: "TabAxis Horizontal",
			Bool: widget.Bool{
				Value: false,
			},
		},
	}
}

type TabAxis struct {
	setting.Bool
}

func (a *TabAxis) GetAxis() layout.Axis {
	if a.Value {
		return layout.Horizontal
	}
	return layout.Vertical
}

func (a *TabAxis) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return a.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (a *TabAxis) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return a.Bool.Layout(theme, gtx, valueInFront, ratioInFront, nil)
}

func (a *TabAxis) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return a.Bool.LayoutSwitch(theme, gtx, nil)
}
