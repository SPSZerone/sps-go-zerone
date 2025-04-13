package item

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func NewUI() UI {
	return UI{}
}

type UI struct {
	Clickable   widget.Clickable
	Menu        component.MenuState
	MenuItems   []widget.Clickable
	ContextArea component.ContextArea
}

func Divider(app *spsgio.Application, gtx layout.Context, item *Item, layoutCtx LayoutContext) layout.Dimensions {
	return component.Divider(app.Theme).Layout(gtx)
}

func SubheadingDivider(
	app *spsgio.Application, gtx layout.Context,
	item *Item, layoutCtx LayoutContext,
	subheading string,
) layout.Dimensions {
	return component.SubheadingDivider(app.Theme, subheading).Layout(gtx)
}

func MenuItem(
	app *spsgio.Application, gtx layout.Context,
	item *Item, layoutCtx LayoutContext,
	clickable *widget.Clickable, label string,
) layout.Dimensions {
	return component.MenuItem(app.Theme, clickable, label).Layout(gtx)
}
