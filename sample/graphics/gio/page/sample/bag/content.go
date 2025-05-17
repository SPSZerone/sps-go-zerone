package bag

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget/material"

	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func (i *Item) LayoutContent(
	theme *material.Theme, gtx layout.Context,
	item *spsitem.Item, layoutCtx spsitem.LayoutContext,
) layout.Dimensions {
	gtx.Constraints.Max = layoutCtx.ContentSize

	if item.UI.MenuItems[0].Clicked(gtx) {
	}
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			property := i.Data.GetProperty(PropertyKeyId)
			return material.H6(theme, fmt.Sprintf("%v", property.Value)).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			property := i.Data.GetProperty(PropertyKeyName)
			return material.H6(theme, fmt.Sprintf("%v", property.Value)).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			property := i.Data.GetProperty(PropertyKeyIcon)
			return material.H6(theme, fmt.Sprintf("%v", property.Value)).Layout(gtx)
		}),
	)
}

func (i *Item) LayoutContentProperty(theme *material.Theme, gtx layout.Context, name, value string) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, name).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, value).Layout(gtx)
		}),
	)
}
