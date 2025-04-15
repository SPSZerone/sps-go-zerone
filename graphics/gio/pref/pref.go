package pref

import (
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/pref/settings"
)

func NewPreferences() Preferences {
	return Preferences{
		Settings: settings.NewSettings(),
	}
}

type Preferences struct {
	Settings settings.Settings
	Pref     Settings
}

type Settings struct {
	NonModalDrawer bool
	BottomBar      bool
	PrefTableStyle bool
	ValueInFront   bool
}
