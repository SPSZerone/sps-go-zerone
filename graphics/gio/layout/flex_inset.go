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

func (f FlexInset) LayoutABWidget(gtx layout.Context, aWidget, bWidget layout.Widget) layout.Dimensions {
	if f.Ratio == 0 {
		f.Ratio = 0.333
	}
	if f.Inset == (layout.Inset{}) {
		f.Inset = DefaultInset
	}
	return f.Flex.Layout(
		gtx,
		layout.Flexed(f.Ratio, func(gtx layout.Context) layout.Dimensions {
			return f.Inset.Layout(gtx, aWidget)
		}),
		layout.Flexed(1-f.Ratio, func(gtx layout.Context) layout.Dimensions {
			return f.Inset.Layout(gtx, bWidget)
		}),
	)
}

func (f FlexInset) LayoutWidgets(gtx layout.Context, widgets ...func() (float32, layout.Widget)) layout.Dimensions {
	if f.Inset == (layout.Inset{}) {
		f.Inset = DefaultInset
	}
	flexChildren := make([]layout.FlexChild, len(widgets))
	for i, customWidget := range widgets {
		ratio, widget := customWidget()
		flexChildren[i] = layout.Flexed(
			ratio,
			func(gtx layout.Context) layout.Dimensions {
				return f.Inset.Layout(gtx, widget)
			},
		)
	}
	return f.Flex.Layout(gtx, flexChildren...)
}
