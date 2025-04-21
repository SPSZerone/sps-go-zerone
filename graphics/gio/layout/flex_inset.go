package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

var DefaultInset = layout.UniformInset(unit.Dp(4))

type FlexedWidget func() (weight float32, widget layout.Widget)

type FlexInset struct {
	Flex  layout.Flex
	Inset layout.Inset
}

func (f FlexInset) LayoutRigidWidgets(gtx layout.Context, widgets ...layout.Widget) layout.Dimensions {
	inset := f.GetInset()
	children := make([]layout.FlexChild, len(widgets))
	for i, widget := range widgets {
		children[i] = layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return inset.Layout(gtx, widget)
			},
		)
	}
	return f.Flex.Layout(gtx, children...)
}

func (f FlexInset) LayoutFlexedWidgetAB(
	gtx layout.Context,
	ratio float32,
	aWidget, bWidget layout.Widget,
) layout.Dimensions {
	if ratio == 0 {
		ratio = 0.333
	}
	inset := f.GetInset()
	return f.Flex.Layout(
		gtx,
		layout.Flexed(ratio, func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, aWidget)
		}),
		layout.Flexed(1-ratio, func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, bWidget)
		}),
	)
}

func (f FlexInset) LayoutFlexedWidgets(gtx layout.Context, widgets ...FlexedWidget) layout.Dimensions {
	inset := f.GetInset()
	children := make([]layout.FlexChild, len(widgets))
	for i, customWidget := range widgets {
		ratio, widget := customWidget()
		children[i] = layout.Flexed(
			ratio,
			func(gtx layout.Context) layout.Dimensions {
				return inset.Layout(gtx, widget)
			},
		)
	}
	return f.Flex.Layout(gtx, children...)
}

func (f FlexInset) GetInset() layout.Inset {
	if f.Inset == (layout.Inset{}) {
		return DefaultInset
	}
	return f.Inset
}
