package gio

import (
	"context"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

type App interface {
	GetContext() context.Context
	GetPref() *spspref.Preferences

	Run(win Window)
	RunWindow(win Window)
}
