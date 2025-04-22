package copy

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"

	spsclipboard "github.com/SPSZerone/sps-go-zerone/clipboard"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
)

type Widget func() (widget layout.Widget)
type OnCopy func() string

func New() Copy {
	return Copy{
		FlexInset: spslayout.New(),
		Spacer:    4,
		Border: widget.Border{
			Color: color.NRGBA{A: 255},
			Width: unit.Dp(2),
		},
		WithBorder: true,
	}
}

type Copy struct {
	FlexInset spslayout.FlexInset

	Border     widget.Border
	WithBorder bool
	Spacer     int

	Clickable widget.Clickable
}

func (c *Copy) GetClickable() *widget.Clickable {
	return &c.Clickable
}

func (c *Copy) LayoutCopy(
	gtx layout.Context, clickWidget layout.Widget, onCopy OnCopy,
) layout.Dimensions {
	if c.Clickable.Clicked(gtx) {
		if onCopy != nil {
			spsclipboard.WriteTextString(onCopy())
		}
	}
	return c.Clickable.Layout(gtx, clickWidget)
}

func (c *Copy) LayoutCopyRigidContent(
	gtx layout.Context,
	copyWidget layout.Widget,
	onCopy OnCopy,
	contentWidgets ...layout.Widget,
) layout.Dimensions {
	if c.WithBorder {
		return c.Border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return c.doLayoutCopyRigidContent(gtx, copyWidget, onCopy, contentWidgets...)
		})
	}
	return c.doLayoutCopyRigidContent(gtx, copyWidget, onCopy, contentWidgets...)
}

func (c *Copy) doLayoutCopyRigidContent(
	gtx layout.Context,
	copyWidget layout.Widget,
	onCopy OnCopy,
	contentWidgets ...layout.Widget,
) layout.Dimensions {
	spacer := c.NewSpacer()
	widgets := make([]layout.Widget, 0, len(contentWidgets)+1)
	for _, contentWidget := range contentWidgets {
		widgets = append(widgets, contentWidget)
		widgets = append(widgets, spacer.Layout)
	}
	widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
		return c.LayoutCopy(gtx, copyWidget, onCopy)
	})
	return c.FlexInset.LayoutRigidWidgets(gtx, widgets...)
}

func (c *Copy) LayoutCopyFlexedContent(
	gtx layout.Context,
	copyWeight float32,
	copyWidget layout.Widget,
	onCopy OnCopy,
	contentWidgets ...spslayout.FlexedWidget,
) layout.Dimensions {
	if c.WithBorder {
		return c.Border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return c.doLayoutCopyFlexedContent(gtx, copyWeight, copyWidget, onCopy, contentWidgets...)
		})
	}
	return c.doLayoutCopyFlexedContent(gtx, copyWeight, copyWidget, onCopy, contentWidgets...)
}

func (c *Copy) doLayoutCopyFlexedContent(
	gtx layout.Context,
	copyWeight float32,
	copyWidget layout.Widget,
	onCopy OnCopy,
	contentWidgets ...spslayout.FlexedWidget,
) layout.Dimensions {
	widgets := make([]spslayout.FlexedWidget, 0, len(contentWidgets)+1)
	for _, contentWidget := range contentWidgets {
		widgets = append(widgets, contentWidget)
	}
	widgets = append(widgets, func() (weight float32, widget layout.Widget) {
		weight = copyWeight
		widget = func(gtx layout.Context) layout.Dimensions {
			return c.LayoutCopy(gtx, copyWidget, onCopy)
		}
		return
	})
	return c.FlexInset.LayoutFlexedWidgets(gtx, widgets...)
}

func (c *Copy) NewSpacer() (spacer spsspacer.Spacer) {
	if c.FlexInset.Flex.Axis == layout.Horizontal {
		spacer = spsspacer.NewWidthSpacer(c.Spacer)
	} else {
		spacer = spsspacer.NewHeightSpacer(c.Spacer)
	}
	return
}
