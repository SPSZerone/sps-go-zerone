package sample

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsgrid "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/grid"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
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

		Tabs: spstab.NewTabsByNames(TabNameGridAndItem),

		Grid: spsgrid.NewGrid(),
	}
	const count = 100
	p.Items = make([]spsitem.Item, count)
	highlightStyle := spsitem.HighlightStyleDefault
	//highlightStyle := spsitem.HighlightStyle(rand.RandomInt(int(spsitem.HighlightStyleDefault), int(spsitem.HighlightStyleCount-1)))
	stackAlignment := layout.Center
	layoutContent := func(i *spsitem.Item, gtx layout.Context, layoutCtx spsitem.LayoutContext) layout.Dimensions {
		baseInfo := material.Body1(app.Theme, fmt.Sprintf("%v", i.Data))
		baseInfo.Font.Style = font.Italic
		baseInfo.Font.Weight = font.Bold
		return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
	}
	for i := 0; i < 100; i++ {
		p.Items[i] = spsitem.NewItem(
			fmt.Sprintf("item-%d", i),
			layoutContent,
			spsitem.OptHighlightStyle(highlightStyle),
			spsitem.OptStackAlignment(stackAlignment),
		)
	}
	return p
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
