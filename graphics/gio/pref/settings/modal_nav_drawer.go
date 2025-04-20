package settings

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsbool "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/wgtbool"
)

func NewModalNavDrawer() ModalNavDrawer {
	return ModalNavDrawer{
		Bool: spsbool.NewBool(
			spsbool.OptName("Use non-modal drawer"),
			spsbool.OptDesc("Use Non-Modal Navigation Drawer"),
			spsbool.OptValue(true),
		),
	}
}

type ModalNavDrawer struct {
	spsbool.Bool
}

func (d *ModalNavDrawer) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
	navAnim *component.VisibilityAnimation,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return d.Layout(window, theme, gtx, valueInFront, ratioInFront, navAnim)
		}),
	}
}

func (d *ModalNavDrawer) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
	navAnim *component.VisibilityAnimation,
) layout.Dimensions {
	return d.Bool.Layout(
		theme, gtx,
		valueInFront, ratioInFront,
		func() {
			d.UpdateNavAnim(gtx, navAnim)
		},
	)
}

func (d *ModalNavDrawer) LayoutSwitch(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	navAnim *component.VisibilityAnimation,
) layout.Dimensions {
	return d.Bool.LayoutSwitch(
		theme, gtx,
		func() {
			d.UpdateNavAnim(gtx, navAnim)
		},
	)
}

func (d *ModalNavDrawer) UpdateNavAnim(gtx layout.Context, navAnim *component.VisibilityAnimation) {
	if d.Value {
		navAnim.Appear(gtx.Now)
	} else {
		navAnim.Disappear(gtx.Now)
	}
}
