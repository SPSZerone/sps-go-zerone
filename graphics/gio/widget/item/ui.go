package item

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsclickable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/clickable"
	spsmenu "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/menu"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func NewUI() UI {
	return UI{
		Surface: spssurface.NewUniformInset(4, 4),
	}
}

type UI struct {
	Surface   spssurface.Surface
	Clickable spsclickable.Clickable
	Menu      spsmenu.Menu
	MenuItems []widget.Clickable
}

func Divider(theme *material.Theme, gtx layout.Context, item *Item, layoutCtx LayoutContext) layout.Dimensions {
	return component.Divider(theme).Layout(gtx)
}

func SubheadingDivider(
	theme *material.Theme, gtx layout.Context,
	item *Item, layoutCtx LayoutContext,
	subheading string,
) layout.Dimensions {
	return component.SubheadingDivider(theme, subheading).Layout(gtx)
}

func MenuItem(
	theme *material.Theme, gtx layout.Context,
	item *Item, layoutCtx LayoutContext,
	clickable *widget.Clickable, label string,
) layout.Dimensions {
	return component.MenuItem(theme, clickable, label).Layout(gtx)
}
