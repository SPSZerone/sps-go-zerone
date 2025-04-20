package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsbool "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/wgtbool"
)

func NewValueInFront() ValueInFront {
	return ValueInFront{
		Bool: spsbool.NewBool(
			spsbool.OptName("Value In Front"),
			spsbool.OptDesc("Use Value In Front"),
			spsbool.OptValue(true),
		),
	}
}

type ValueInFront struct {
	spsbool.Bool
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
