package gio

import (
	"context"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

type NewWindow func(app App, fromWin Window) Window

type Window interface {
	GetContext() context.Context
	GetPref() *spspref.Preferences

	Run()
}
