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

func (f *ValueInFront) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return f.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (f *ValueInFront) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return f.Bool.Layout(theme, gtx, valueInFront, ratioInFront, nil)
}

func (f *ValueInFront) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return f.Bool.LayoutSwitch(theme, gtx, nil)
}
