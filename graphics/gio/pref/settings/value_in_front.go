package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewValueInFront() ValueInFront {
	return ValueInFront{
		Bool: setting.Bool{
			Name: "Value In Front",
			Desc: "Use Value In Front",
			Bool: widget.Bool{
				Value: true,
			},
		},
	}
}

type ValueInFront struct {
	setting.Bool
}

func (d *ValueInFront) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return d.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (d *ValueInFront) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return d.Bool.Layout(theme, gtx, valueInFront, ratioInFront, nil)
}

func (d *ValueInFront) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return d.Bool.LayoutSwitch(theme, gtx, nil)
}
