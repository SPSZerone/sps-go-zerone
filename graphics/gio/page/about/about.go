package about

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spstable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/table"
	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

func New(pages *spswin.Pages) *Page {
	return &Page{
		Pages: pages,
		Table: spstable.NewTable(spstable.OptHeaders([]spstable.Header{{Text: "Key"}, {Text: "Value"}}...)),
	}
}

var _ spswin.Page = (*Page)(nil)

type Page struct {
	widget.List
	*spswin.Pages

	Table spstable.Table
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "About",
		Icon: spsicon.ActionHelp,
	}
}

func (p *Page) OnEventPre(win *spswin.Window, evt event.Event, param any) {

}

func (p *Page) OnEventPost(win *spswin.Window, evt event.Event, param any) {

}

func (p *Page) Layout(win *spswin.Window, gtx layout.Context, param any) layout.Dimensions {
	p.List.Axis = layout.Vertical

	keys := []string{
		"Name", "Author", "License",
	}
	values := []string{
		"SPS Gio Framework base on Gio",
		"SPSZerone",
		"GPLv3",
	}

	dimensioner := func(axis layout.Axis, index, constraint, minSize, height int) int {
		switch axis {
		case layout.Horizontal:
			var widthUnit int
			switch index {
			case 0:
				widthUnit = gtx.Dp(unit.Dp(100))
			case 1:
				widthUnit = gtx.Dp(unit.Dp(300))
			}
			return widthUnit
		default:
			return height
		}
	}
	cell := func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
		switch col {
		case 1:
			labelStyle.Text = values[row]
		default:
			labelStyle.Text = keys[row]
		}
		return labelStyle.Layout(gtx)
	}

	return material.List(win.Theme, &p.List).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.Table.Layout(win.Theme, gtx, len(keys), dimensioner, cell)
			}),
		)
	})
}
