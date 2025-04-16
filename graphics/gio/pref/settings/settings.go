package settings

func NewSettings() Settings {
	return Settings{
		ModalNavDrawer: NewModalNavDrawer(),
		TabAxis:        NewTabAxis(),
		ValueInFront:   NewValueInFront(),
		BottomBar:      NewBottomBar(),
		Decorated:      NewDecorated(),
	}
}

type Settings struct {
	ModalNavDrawer ModalNavDrawer
	TabAxis        TabAxis
	ValueInFront   ValueInFront
	BottomBar      BottomBar
	Decorated      Decorated
}
