package gio

import (
	"context"

	"github.com/rs/zerolog"

	spspref "github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

type (
	OnAppCreate func(app App)
	OnAppStart  func(app App)
	OnAppStop   func(app App)
)

type App interface {
	GetContext() context.Context
	GetPref() *spspref.Preferences

	Run(win Window)
	RunWindow(win Window)

	GetLogger() *zerolog.Logger
}
