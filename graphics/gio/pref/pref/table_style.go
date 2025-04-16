package pref

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewTableStyle() TableStyle {
	return TableStyle{
		Bool: setting.Bool{
			Name: "TableStyle",
			Desc: "TableStyle",
			Bool: widget.Bool{
				Value: false,
			},
		},
	}
}

type TableStyle struct {
	setting.Bool
}

func (d *TableStyle) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return d.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (d *TableStyle) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return d.Bool.Layout(theme, gtx, valueInFront, ratioInFront, nil)
}

func (d *TableStyle) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return d.Bool.LayoutSwitch(theme, gtx, nil)
}
