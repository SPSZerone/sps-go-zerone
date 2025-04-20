package gio

import (
	"context"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

type NewWindow func(app App, fromWin Window) Window

type Window interface {
	GetId() any
	GetTitle() string

	GetContext() context.Context
	GetPref() *spspref.Preferences

	GetWindow() *app.Window
	GetOps() *op.Ops
	GetTheme() *material.Theme
	GetDeco() *widget.Decorations

	GetLogger() *zerolog.Logger

	Run()
}
