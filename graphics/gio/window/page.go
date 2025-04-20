package window

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/x/component"
)

type Page interface {
	Actions() []component.AppBarAction
	Overflow() []component.OverflowAction
	NavItem() component.NavItem

	OnEventPre(app *Window, evt event.Event, param any)
	OnEventPost(app *Window, evt event.Event, param any)
	Layout(app *Window, gtx layout.Context, param any) layout.Dimensions
}
