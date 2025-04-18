package bag

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func NewItem(data any, app *spsgio.Application) spsitem.Item {
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
	InitMenu(&item, app)
	return item
}

func InitMenu(item *spsitem.Item, app *spsgio.Application) {
	item.UI.MenuItems = []widget.Clickable{
		{},
	}
	item.UI.Menu = component.MenuState{
		Options: []func(gtx layout.Context) layout.Dimensions{
			func(gtx layout.Context) layout.Dimensions {
				return spssurface.NewSurface().Layout(app, gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H6(app.Theme, "Item Info").Layout(gtx)
				})
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{}.Layout(app, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Left:  unit.Dp(16),
					Right: unit.Dp(16),
				}.Layout(gtx, material.H6(app.Theme, fmt.Sprintf("%v", item.Data)).Layout)
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{Subheading: "Action"}.Layout(app, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return component.MenuItem(app.Theme, &item.UI.MenuItems[0], "Use").Layout(gtx)
			},
		},
	}
}

func LayoutContent(
	app *spsgio.Application, gtx layout.Context,
	item *spsitem.Item, layoutCtx spsitem.LayoutContext,
) layout.Dimensions {
	baseInfo := material.Body1(app.Theme, fmt.Sprintf("%v", item.Data))
	baseInfo.Font.Style = font.Italic
	baseInfo.Font.Weight = font.Bold
	if item.UI.MenuItems[0].Clicked(gtx) {
		app.Logger.Info().Msgf("Use %v", item.Data)
	}
	return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
}

func LayoutDetail(
	app *spsgio.Application, gtx layout.Context,
	item *spsitem.Item,
) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H5(app.Theme, "Item Detail")
			baseInfo.Font.Style = font.Italic
			baseInfo.Font.Weight = font.Bold
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H6(app.Theme, fmt.Sprintf("%v", item.Data))
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
	)
}
