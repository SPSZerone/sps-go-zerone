package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

var DefaultInset = layout.UniformInset(unit.Dp(4))

type FlexInset struct {
	Flex  layout.Flex
	Inset layout.Inset

	Ratio float32
}

func (i FlexInset) Layout(gtx layout.Context, aWidget, bWidget layout.Widget) layout.Dimensions {
	if i.Ratio == 0 {
		i.Ratio = 0.333
	}
	if i.Inset == (layout.Inset{}) {
		i.Inset = DefaultInset
	}
	return i.Flex.Layout(gtx,
		layout.Flexed(i.Ratio, func(gtx layout.Context) layout.Dimensions {
			return i.Inset.Layout(gtx, aWidget)
		}),
		layout.Flexed(1-i.Ratio, func(gtx layout.Context) layout.Dimensions {
			return i.Inset.Layout(gtx, bWidget)
		}),
	)
}
