package bag

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func NewItem(data any, theme *material.Theme) spsitem.Item {
	dimensions := spsitem.NewDimensions()
	dimensions.ContentWidth = 100
	dimensions.ContentHeight = 100
	highlightStyle := spsitem.HighlightStyleDefault
	//highlightStyle := spsitem.HighlightStyle(rand.RandomInt(int(spsitem.HighlightStyleDefault), int(spsitem.HighlightStyleCount-1)))
	stackAlignment := layout.Center

	item := spsitem.NewItem(
		data,
		LayoutContent,
		spsitem.OptHighlightStyle(highlightStyle),
		spsitem.OptStackAlignment(stackAlignment),
		spsitem.OptDimensions(dimensions),
		spsitem.OptLayoutDetail(LayoutDetail),
	)
	InitMenu(&item, theme)
	return item
}

func InitMenu(item *spsitem.Item, theme *material.Theme) {
	item.UI.MenuItems = []widget.Clickable{
		{},
	}
	item.UI.Menu = component.MenuState{
		Options: []func(gtx layout.Context) layout.Dimensions{
			func(gtx layout.Context) layout.Dimensions {
				return spssurface.NewSurface().Layout(theme, gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H6(theme, "Item Info").Layout(gtx)
				})
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{}.Layout(theme, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Left:  unit.Dp(16),
					Right: unit.Dp(16),
				}.Layout(gtx, material.H6(theme, fmt.Sprintf("%v", item.Data)).Layout)
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{Subheading: "Action"}.Layout(theme, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return component.MenuItem(theme, &item.UI.MenuItems[0], "Use").Layout(gtx)
			},
		},
	}
}

func LayoutContent(
	theme *material.Theme, gtx layout.Context,
	item *spsitem.Item, layoutCtx spsitem.LayoutContext,
) layout.Dimensions {
	baseInfo := material.Body1(theme, fmt.Sprintf("%v", item.Data))
	baseInfo.Font.Style = font.Italic
	baseInfo.Font.Weight = font.Bold
	if item.UI.MenuItems[0].Clicked(gtx) {
	}
	return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
}

func LayoutDetail(
	theme *material.Theme, gtx layout.Context,
	item *spsitem.Item,
) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H5(theme, "Item Detail")
			baseInfo.Font.Style = font.Italic
			baseInfo.Font.Weight = font.Bold
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H6(theme, fmt.Sprintf("%v", item.Data))
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
	)
}
