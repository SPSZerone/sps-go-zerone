package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsbool "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/wgtbool"
)

func NewTabAxis() TabAxis {
	return TabAxis{
		Bool: spsbool.New(
			spsbool.OptName("TabAxis Horizontal"),
			spsbool.OptDesc("TabAxis Horizontal"),
		),
	}
}

type TabAxis struct {
	spsbool.Bool
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
	return a.Bool.LayoutDefaultKeyWidget(theme, gtx, valueInFront, ratioInFront, nil)
}

func (a *TabAxis) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return a.Bool.LayoutSwitch(theme, gtx, nil)
}
