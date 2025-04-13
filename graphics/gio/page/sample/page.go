package sample

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spsgrid "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/grid"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
	spstab "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/tab"
)

const (
	TabIdxGridAndItem = iota
)

const (
	TabNameGridAndItem = "Grid & Item"
)

func New(app *spsgio.Application) *Page {
	p := &Page{
		Pages: &app.Pages,
		Tabs:  spstab.NewTabsByNames(TabNameGridAndItem),
		Grid:  spsgrid.NewGrid(),
	}

	// test item
	const count = 100
	p.Items = make([]spsitem.Item, count)
	highlightStyle := spsitem.HighlightStyleDefault
	//highlightStyle := spsitem.HighlightStyle(rand.RandomInt(int(spsitem.HighlightStyleDefault), int(spsitem.HighlightStyleCount-1)))
	stackAlignment := layout.Center
	content := func(app *spsgio.Application, gtx layout.Context, item *spsitem.Item, layoutCtx spsitem.LayoutContext) layout.Dimensions {
		baseInfo := material.Body1(app.Theme, fmt.Sprintf("%v", item.Data))
		baseInfo.Font.Style = font.Italic
		baseInfo.Font.Weight = font.Bold
		if item.UI.MenuItems[0].Clicked(gtx) {
			app.Logger.Info().Msgf("Use %v", item.Data)
		}
		return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
	}
	for i := 0; i < 100; i++ {
		item := spsitem.NewItem(
			fmt.Sprintf("item-%d", i),
			content,
			spsitem.OptHighlightStyle(highlightStyle),
			spsitem.OptStackAlignment(stackAlignment),
		)
		itemMenu(app, &item)
		p.Items[i] = item
	}
	return p
}

func itemMenu(app *spsgio.Application, item *spsitem.Item) {
	item.UI.MenuItems = []widget.Clickable{
		{},
	}
	item.UI.Menu = component.MenuState{
		Options: []func(gtx layout.Context) layout.Dimensions{
			func(gtx layout.Context) layout.Dimensions {
				return spssurface.NewSurface().Layout(app.Theme, gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H6(app.Theme, "Item Info").Layout(gtx)
				})
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{}.Layout(app.Theme, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Left:  unit.Dp(16),
					Right: unit.Dp(16),
				}.Layout(gtx, material.H6(app.Theme, fmt.Sprintf("%v", item.Data)).Layout)
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{Subheading: "Action"}.Layout(app.Theme, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return component.MenuItem(app.Theme, &item.UI.MenuItems[0], "Use").Layout(gtx)
			},
		},
	}
}

var _ spsgio.Page = (*Page)(nil)

type Page struct {
	widget.List
	*spsgio.Pages

	Tabs spstab.Tabs

	Grid       spsgrid.Grid
	SelectItem int
	Items      []spsitem.Item
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Sample",
		Icon: spsicon.ActionHelp,
	}
}

func (p *Page) OnEventPre(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) OnEventPost(app *spsgio.Application, evt event.Event, param any) {

}

func (p *Page) Layout(app *spsgio.Application, gtx layout.Context, param any) layout.Dimensions {
	p.List.Axis = layout.Vertical
	return material.List(app.Theme, &p.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return p.Tabs.Layout(app, gtx, param, func(gtx layout.Context, selected int) layout.Dimensions {
			switch selected {
			case TabIdxGridAndItem:
				return p.GridAndItem(app, gtx, param)
			default:
				return layout.Dimensions{}
			}
		})
	})
}

func (p *Page) GetItem(index int) *spsitem.Item {
	if index < 0 || index >= len(p.Items) {
		return nil
	}
	return &p.Items[index]
}
