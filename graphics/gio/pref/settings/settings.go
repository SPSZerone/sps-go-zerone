package settings

func NewSettings() Settings {
	return Settings{
		Decorated:      NewDecorated(),
		ModalNavDrawer: NewModalNavDrawer(),
		BottomBar:      NewBottomBar(),
	}
}

type Settings struct {
	Decorated      Decorated
	ModalNavDrawer ModalNavDrawer
	BottomBar      BottomBar
}
