package gio

import (
	"context"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

type (
	NewWindow func(app App, fromWin Window) Window

	OnInitPre  func(win Window)
	OnInitPost func(win Window)
	OnStart    func(win Window)
	OnLoop     func(win Window) error
	OnStop     func(win Window)

	OnEventPre  func(win Window, evt event.Event, param any)
	OnEventPost func(win Window, evt event.Event, param any)
)

type Window interface {
	GetId() any
	GetTitle() string

	GetContext() context.Context
	GetPref() *spspref.Preferences

	GetPages() Pages

	GetWindow() *app.Window
	GetOps() *op.Ops
	GetTheme() *material.Theme
	GetDeco() *widget.Decorations

	GetLogger() *zerolog.Logger
	LogPrefix() string

	OnFrameEvent(e app.FrameEvent, param any)
	Layout(gtx layout.Context, param any)

	Run()
	GoRun(run func())
}
