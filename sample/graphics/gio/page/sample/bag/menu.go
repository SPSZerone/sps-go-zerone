package bag

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func (i *Item) InitMenu(theme *material.Theme) {
	i.InitMenuItem(theme)
	i.InitMenuDetail(theme)
}

func (i *Item) InitMenuItem(theme *material.Theme) {
	item := &i.Item
	item.UI.MenuItems = []widget.Clickable{
		{},
	}
	item.UI.Menu.AddWidgets(
		func(gtx layout.Context) layout.Dimensions {
			return spssurface.NewSurface().Layout(theme, gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H6(theme, "Item Info").Layout(gtx)
			})
		},
		func(gtx layout.Context) layout.Dimensions {
			return spsdivider.Divider{}.Layout(theme, gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Left: unit.Dp(16), Right: unit.Dp(16)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					property := i.Data.GetProperty(PropertyKeyId)
					return i.LayoutContentProperty(theme, gtx,
						fmt.Sprintf("%v", property.Key),
						fmt.Sprintf(" %v", property.Value),
					)
				},
			)
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Left: unit.Dp(16), Right: unit.Dp(16)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					property := i.Data.GetProperty(PropertyKeyName)
					return i.LayoutContentProperty(theme, gtx,
						fmt.Sprintf("%v", property.Key),
						fmt.Sprintf(" %v", property.Value),
					)
				},
			)
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Left: unit.Dp(16), Right: unit.Dp(16)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					property := i.Data.GetProperty(PropertyKeyIcon)
					return i.LayoutContentProperty(theme, gtx,
						fmt.Sprintf("%v", property.Key),
						fmt.Sprintf(" %v", property.Value),
					)
				},
			)
		},
		func(gtx layout.Context) layout.Dimensions {
			return spsdivider.Divider{Subheading: "Action"}.Layout(theme, gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return component.MenuItem(theme, &item.UI.MenuItems[0], "Use").Layout(gtx)
		},
	)
}

func (i *Item) InitMenuDetail(theme *material.Theme) {
	i.UI.MenuDetail.AddWidgets(
		func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyId))
		},
		func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyName))
		},
		func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyIcon))
		},
	)
}
