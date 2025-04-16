package pref

import (
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/pref"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/settings"
)

func NewPreferences() Preferences {
	return Preferences{
		Settings: settings.NewSettings(),
		Pref:     pref.NewPref(),
	}
}

type Preferences struct {
	Settings settings.Settings
	Pref     pref.Pref
}
