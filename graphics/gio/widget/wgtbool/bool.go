package wgtbool

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spsnerdfont "github.com/SPSZerone/sps-go-zerone/graphics/gio/font/nerdfont"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

func NewBool(opts ...Option) Bool {
	b := newBool()
	b.Update(opts...)
	return b
}

func newBool() Bool {
	return Bool{}
}

type Bool struct {
	widget.Bool
	Name string
	Desc string
}

func (b *Bool) Update(opts ...Option) {
	for _, opt := range opts {
		opt(b)
	}
}

func (b *Bool) Layout(
	theme *material.Theme, gtx layout.Context,
	keyWidget layout.Widget,
	valueInFront bool, ratioInFront float32,
	onValueChanged func(),
) layout.Dimensions {
	if keyWidget == nil {
		keyWidget = func(gtx layout.Context) layout.Dimensions {
			labelStyle := material.Body1(theme, b.Name)
			labelStyle.Font.Typeface = spsnerdfont.MesloLGSNerdFontMono
			return labelStyle.Layout(gtx)
		}
	}

	valueWidget := func(gtx layout.Context) layout.Dimensions {
		return b.LayoutSwitch(theme, gtx, onValueChanged)
	}

	var aWidget, bWidget layout.Widget
	if valueInFront {
		aWidget, bWidget = valueWidget, keyWidget
	} else {
		aWidget, bWidget = keyWidget, valueWidget
	}
	return spslayout.FlexInset{}.LayoutFlexedWidgetAB(gtx, ratioInFront, aWidget, bWidget)
}

func (b *Bool) LayoutDefaultKeyWidget(
	theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
	onValueChanged func(),
) layout.Dimensions {
	return b.Layout(
		theme, gtx,
		nil,
		valueInFront, ratioInFront,
		onValueChanged)
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
