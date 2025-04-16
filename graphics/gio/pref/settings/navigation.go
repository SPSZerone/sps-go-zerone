package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewNavigation() Navigation {
	return Navigation{
		Float: setting.Float{
			Name: "Navigation Ratio",
			Desc: "Navigation Ratio",
			Float: widget.Float{
				Value: 0.2,
			},
		},
	}
}

type Navigation struct {
	setting.Float
}

func (n *Navigation) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return n.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (n *Navigation) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return n.Float.Layout(theme, gtx, valueInFront, ratioInFront, nil)
}

func (n *Navigation) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return n.Float.LayoutSwitch(theme, gtx, nil)
}
