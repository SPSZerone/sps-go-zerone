package pref

import "gioui.org/widget"

type SettingBool struct {
	Name   string
	Desc   string
	Widget widget.Bool
}

func NewDecorated() SettingBool {
	return SettingBool{
		Name: "Decorated",
		Desc: "Use decorated",
	}
}

func NewNonModalDrawer() SettingBool {
	return SettingBool{
		Name: "Use non-modal drawer",
		Desc: "Use Non-Modal Navigation Drawer",
	}
}

func NewBottomBar() SettingBool {
	return SettingBool{
		Name: "Bottom Bar",
		Desc: "Use Bottom Bar",
	}
}

func NewPrefTableStyle() SettingBool {
	return SettingBool{
		Name: "Pref Table Style",
		Desc: "Use Pref Table Style",
	}
}

func NewSwitchInFront() SettingBool {
	return SettingBool{
		Name: "Switch In Front",
		Desc: "Use Switch In Front",
	}
}
