package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsbool "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/wgtbool"
)

func NewBottomBar() BottomBar {
	return BottomBar{
		Bool: spsbool.NewBool(
			spsbool.OptName("Bottom Bar"),
			spsbool.OptDesc("Use Bottom Bar"),
		),
	}
}

type BottomBar struct {
	spsbool.Bool
}

func (b *BottomBar) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return b.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (b *BottomBar) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return b.Bool.Layout(
		theme, gtx,
		valueInFront, ratioInFront,
		func() {

		},
	)
}

func (b *BottomBar) LayoutSwitch(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	onValueChanged func(),
) layout.Dimensions {
	return b.Bool.LayoutSwitch(theme, gtx, onValueChanged)
}
