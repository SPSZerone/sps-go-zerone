package settings

func NewSettings() Settings {
	return Settings{
		ModalNavDrawer: NewModalNavDrawer(),
		TabAxis:        NewTabAxis(),
		BottomBar:      NewBottomBar(),
		Decorated:      NewDecorated(),
	}
}

type Settings struct {
	ModalNavDrawer ModalNavDrawer
	TabAxis        TabAxis
	BottomBar      BottomBar
	Decorated      Decorated
}
