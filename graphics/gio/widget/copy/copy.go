package copy

import (
	"gioui.org/layout"

	spsclipboard "github.com/SPSZerone/sps-go-zerone/clipboard"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsclickable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/clickable"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
)

type Widget func() (widget layout.Widget)
type OnCopy func() string

type Copy struct {
	spslayout.FlexInset

	Spacer int

	Clickable spsclickable.Clickable
}

func (c *Copy) Layout(
	gtx layout.Context,
	contentWidget, clickWidget layout.Widget,
	onCopy OnCopy,
) layout.Dimensions {
	spacer := c.NewSpacer()
	return c.LayoutRigidWidgets(
		gtx,
		contentWidget,
		spacer.Layout,
		func(gtx layout.Context) layout.Dimensions {
			if c.Clickable.Clicked(gtx) {
				if onCopy != nil {
					spsclipboard.WriteTextString(onCopy())
				}
			}
			return c.Clickable.Layout(gtx, clickWidget)
		},
	)
}

func (c *Copy) NewSpacer() (spacer spsspacer.Spacer) {
	if c.Flex.Axis == layout.Horizontal {
		spacer = spsspacer.NewWidthSpacer(c.Spacer)
	} else {
		spacer = spsspacer.NewHeightSpacer(c.Spacer)
	}
	return
}
