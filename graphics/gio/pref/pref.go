package pref

type Preferences struct {
	Settings Settings
}

type Settings struct {
	Decorated      bool
	NonModalDrawer bool
	BottomBar      bool
}
