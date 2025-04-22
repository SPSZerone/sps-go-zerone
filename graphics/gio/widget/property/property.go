package property

import (
	"gioui.org/layout"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
)

func NewWithUniformSpacer(spacer int) Property {
	return NewWithSpacer(spacer, spacer, spacer)
}

func NewWithStartEndSpacer(spacerStart, spacerEnd int) Property {
	return NewWithSpacer(spacerStart, 0, spacerEnd)
}

func NewWithSpacer(spacerStart, spacerMiddle, spacerEnd int) Property {
	p := New()
	p.SpacerStart = spacerStart
	p.SpacerMiddle = spacerMiddle
	p.SpacerEnd = spacerEnd
	return p
}

func New() Property {
	return Property{
		FlexInset: spslayout.FlexInset{
			Flex: layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Baseline,
			},
		},
	}
}

type Property struct {
	spslayout.FlexInset

	SpacerStart  int
	SpacerMiddle int
	SpacerEnd    int
}

func (p *Property) LayoutRigidWidgets(gtx layout.Context, widgets ...layout.Widget) layout.Dimensions {
	finalWidgets := make([]layout.Widget, 0, (len(widgets)<<1)+1)
	finalWidgets = p.RigidAppendSpacer(finalWidgets, p.SpacerStart)
	for _, widget := range widgets {
		finalWidgets = p.RigidAppendSpacer(finalWidgets, p.SpacerMiddle)
		finalWidgets = append(finalWidgets, widget)
	}
	finalWidgets = p.RigidAppendSpacer(finalWidgets, p.SpacerEnd)
	return p.FlexInset.LayoutRigidWidgets(gtx, finalWidgets...)
}

func (p *Property) LayoutFlexWidgetAB(gtx layout.Context, ratio float32, aWidget, bWidget layout.Widget) layout.Dimensions {
	return p.LayoutFlexWidgets(
		gtx,
		func() (weight float32, widget layout.Widget) {
			weight = ratio
			widget = aWidget
			return
		},
		func() (weight float32, widget layout.Widget) {
			weight = 1 - ratio
			widget = bWidget
			return
		},
	)
}

func (p *Property) LayoutFlexWidgets(gtx layout.Context, widgets ...spslayout.FlexedWidget) layout.Dimensions {
	return p.FlexInset.LayoutFlexedWidgets(gtx, widgets...)
}

func (p *Property) NewSpacers() (start, middle, end spsspacer.Spacer) {
	if p.Flex.Axis == layout.Horizontal {
		start = spsspacer.NewWidthSpacer(p.SpacerStart)
		middle = spsspacer.NewWidthSpacer(p.SpacerMiddle)
		end = spsspacer.NewWidthSpacer(p.SpacerEnd)
	} else {
		start = spsspacer.NewHeightSpacer(p.SpacerStart)
		middle = spsspacer.NewHeightSpacer(p.SpacerMiddle)
		end = spsspacer.NewHeightSpacer(p.SpacerEnd)
	}
	return
}

func (p *Property) NewSpacer(spacer int) spsspacer.Spacer {
	if p.Flex.Axis == layout.Horizontal {
		return spsspacer.NewWidthSpacer(spacer)
	}
	return spsspacer.NewHeightSpacer(spacer)
}

func (p *Property) FlexAppendSpacer(widgets []spslayout.FlexedWidget, spacer int) []spslayout.FlexedWidget {
	if spacer > 0 {
		widgets = append(widgets, func() (weight float32, widget layout.Widget) {
			weight = 0
			widget = p.NewSpacer(spacer).Layout
			return
		})
	}
	return widgets
}

func (p *Property) RigidAppendSpacer(widgets []layout.Widget, spacer int) []layout.Widget {
	if spacer > 0 {
		widgets = append(widgets, p.NewSpacer(spacer).Layout)
	}
	return widgets
}
