package gio

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/x/component"
)

type Pages interface {
	GetWindow() Window

	GetAppBar() *component.AppBar
	GetModalNavDrawer() *component.ModalNavDrawer
	GetModalLayer() *component.ModalLayer
	GetNavAnim() *component.VisibilityAnimation

	Register(tag any, page Page)
	SwitchTo(tag any) Page
	Start(tag any) Page

	Current() Page
	Count() int

	OnEventPre(win Window, evt event.Event, param any)
	OnEventPost(win Window, evt event.Event, param any)
	Layout(win Window, gtx layout.Context, param any, deco func() layout.FlexChild) layout.Dimensions
}
