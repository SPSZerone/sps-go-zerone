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
	spsbag "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func NewItem(data any, app *spsgio.Application) *Item {
	dimensions := spsitem.NewDimensions()
	dimensions.ContentWidth = 200
	dimensions.ContentHeight = 200
	highlightStyle := spsitem.HighlightStyleDefault
	//highlightStyle := spsitem.HighlightStyle(rand.RandomInt(int(spsitem.HighlightStyleDefault), int(spsitem.HighlightStyleCount-1)))
	stackAlignment := layout.Center

	i := &Item{}
	item := spsitem.NewItem(
		data,
		i.LayoutContent,
		spsitem.OptHighlightStyle(highlightStyle),
		spsitem.OptStackAlignment(stackAlignment),
		spsitem.OptDimensions(dimensions),
	)
	i.Item = item
	i.InitMenu(app)
	return i
}

var _ spsbag.Item = (*Item)(nil)

type Item struct {
	Item spsitem.Item
}

func (i *Item) InitMenu(app *spsgio.Application) {
	i.Item.UI.MenuItems = []widget.Clickable{
		{},
	}
	i.Item.UI.Menu = component.MenuState{
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
				}.Layout(gtx, material.H6(app.Theme, fmt.Sprintf("%v", i.Item.Data)).Layout)
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{Subheading: "Action"}.Layout(app, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return component.MenuItem(app.Theme, &i.Item.UI.MenuItems[0], "Use").Layout(gtx)
			},
		},
	}
}

func (i *Item) GetItem() *spsitem.Item {
	return &i.Item
}

func (i *Item) UpdateItem(item spsitem.Item) {
	i.Item = item
}

func (i *Item) Layout(
	app *spsgio.Application, gtx layout.Context, highlight bool,
) (dimensions layout.Dimensions, clicked bool) {
	return i.Item.Layout(app, gtx, highlight)
}

func (i *Item) LayoutDetail(app *spsgio.Application, gtx layout.Context, title string) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H5(app.Theme, title)
			baseInfo.Font.Style = font.Italic
			baseInfo.Font.Weight = font.Bold
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H6(app.Theme, fmt.Sprintf("%v", i.Item.Data))
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
	)
}

func (i *Item) LayoutContent(app *spsgio.Application, gtx layout.Context, item *spsitem.Item, layoutCtx spsitem.LayoutContext) layout.Dimensions {
	baseInfo := material.Body1(app.Theme, fmt.Sprintf("%v", item.Data))
	baseInfo.Font.Style = font.Italic
	baseInfo.Font.Weight = font.Bold
	if item.UI.MenuItems[0].Clicked(gtx) {
		app.Logger.Info().Msgf("Use %v", item.Data)
	}
	return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
}
