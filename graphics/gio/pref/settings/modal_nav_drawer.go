package settings

import (
	gioapp "gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewModalNavDrawer() ModalNavDrawer {
	return ModalNavDrawer{
		Bool: setting.Bool{
			Name: "Use non-modal drawer",
			Desc: "Use Non-Modal Navigation Drawer",
		},
	}
}

type ModalNavDrawer struct {
	setting.Bool
}

func (d *ModalNavDrawer) FlexChild(
	window *gioapp.Window, theme *material.Theme, gtx layout.Context,
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
	window *gioapp.Window, theme *material.Theme, gtx layout.Context,
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
	window *gioapp.Window, theme *material.Theme, gtx layout.Context,
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
