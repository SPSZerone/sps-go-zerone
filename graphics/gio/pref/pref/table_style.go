package pref

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/wgtbool"
)

func NewTableStyle() TableStyle {
	return TableStyle{
		Bool: wgtbool.Bool{
			Name: "TableStyle",
			Desc: "TableStyle",
			Bool: widget.Bool{
				Value: false,
			},
		},
	}
}

type TableStyle struct {
	wgtbool.Bool
}

func (s *TableStyle) FlexChild(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) []layout.FlexChild {
	return []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.Layout(window, theme, gtx, valueInFront, ratioInFront)
		}),
	}
}

func (s *TableStyle) Layout(
	window *app.Window, theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
) layout.Dimensions {
	return s.Bool.Layout(theme, gtx, valueInFront, ratioInFront, nil)
}

func (s *TableStyle) LayoutSwitch(window *app.Window, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return s.Bool.LayoutSwitch(theme, gtx, nil)
}
