package setting

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

type Float struct {
	widget.Float
	Name string
	Desc string
}

func (f *Float) Layout(
	theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
	onValueChanged func(),
) layout.Dimensions {
	key := material.Body1(theme, fmt.Sprintf("%v %v", f.Name, f.Float.Value)).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return f.LayoutSwitch(theme, gtx, onValueChanged)
	}
	var aWidget, bWidget layout.Widget
	if valueInFront {
		aWidget, bWidget = value, key
	} else {
		aWidget, bWidget = key, value
	}
	return spslayout.FlexInset{Ratio: ratioInFront}.LayoutABWidget(gtx, aWidget, bWidget)
}

func (f *Float) LayoutSwitch(
	theme *material.Theme, gtx layout.Context,
	onValueChanged func(),
) layout.Dimensions {
	if f.Float.Update(gtx) {
		if onValueChanged != nil {
			onValueChanged()
		}
	}
	return material.Slider(theme, &f.Float).Layout(gtx)
}
