package property

import (
	"gioui.org/layout"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
)

func NewWithUniformSpacer(spacer int) Property {
	p := Property{
		FlexInset: spslayout.FlexInset{
			Flex: layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Baseline,
			},
		},
		SpacerStart:  spacer,
		SpacerMiddle: spacer,
		SpacerEnd:    spacer,
	}
	return p
}

type Property struct {
	spslayout.FlexInset

	SpacerStart  int
	SpacerMiddle int
	SpacerEnd    int
}

func (p *Property) appendSpacer(widgets []layout.Widget, spacer int) []layout.Widget {
	if spacer > 0 {
		widgets = append(widgets, p.NewSpacer(spacer).Layout)
	}
	return widgets
}

func (p *Property) LayoutRigidWidgets(gtx layout.Context, widgets ...layout.Widget) layout.Dimensions {
	finalWidgets := make([]layout.Widget, 0, (len(widgets)<<1)+1)
	finalWidgets = p.appendSpacer(finalWidgets, p.SpacerStart)
	for _, widget := range widgets {
		finalWidgets = p.appendSpacer(finalWidgets, p.SpacerMiddle)
		finalWidgets = append(finalWidgets, widget)
	}
	finalWidgets = p.appendSpacer(finalWidgets, p.SpacerEnd)
	return p.FlexInset.LayoutRigidWidgets(gtx, finalWidgets...)
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
