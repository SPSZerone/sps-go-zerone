package setting

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

type Bool struct {
	widget.Bool
	Name string
	Desc string
}

func (b *Bool) Layout(
	theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
	onValueChanged func(),
) layout.Dimensions {
	key := material.Body1(theme, b.Name).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return b.LayoutSwitch(theme, gtx, onValueChanged)
	}
	var aWidget, bWidget layout.Widget
	if valueInFront {
		aWidget, bWidget = value, key
	} else {
		aWidget, bWidget = key, value
	}
	return spslayout.FlexInset{Ratio: ratioInFront}.LayoutABWidget(gtx, aWidget, bWidget)
}

func (b *Bool) LayoutSwitch(
	theme *material.Theme, gtx layout.Context,
	onValueChanged func(),
) layout.Dimensions {
	if b.Bool.Update(gtx) {
		if onValueChanged != nil {
			onValueChanged()
		}
	}
	return material.Switch(theme, &b.Bool, b.Desc).Layout(gtx)
}
