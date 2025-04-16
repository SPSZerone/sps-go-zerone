package settings

func NewSettings() Settings {
	return Settings{
		Navigation:     NewNavigation(),
		ModalNavDrawer: NewModalNavDrawer(),
		TabAxis:        NewTabAxis(),
		ValueInFront:   NewValueInFront(),
		BottomBar:      NewBottomBar(),
		Decorated:      NewDecorated(),
	}
}

type Settings struct {
	Navigation     Navigation
	ModalNavDrawer ModalNavDrawer
	TabAxis        TabAxis
	ValueInFront   ValueInFront
	BottomBar      BottomBar
	Decorated      Decorated
}
