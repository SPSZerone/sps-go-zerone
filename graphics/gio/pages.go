package gio

import (
	"gioui.org/io/event"
	"gioui.org/layout"
)

type Pages interface {
	Register(tag any, page Page)
	SwitchTo(tag any) Page
	Start(tag any) Page

	Current() Page
	Count() int

	OnEventPre(win Window, evt event.Event, param any)
	OnEventPost(win Window, evt event.Event, param any)
	Layout(win Window, gtx layout.Context, param any, deco func() layout.FlexChild) layout.Dimensions
}
