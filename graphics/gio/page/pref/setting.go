package pref

import "gioui.org/widget"

type SettingBool struct {
	Name   string
	Desc   string
	Widget widget.Bool
}

func NewPrefTableStyle() SettingBool {
	return SettingBool{
		Name: "Pref Table Style",
		Desc: "Use Pref Table Style",
	}
}

func NewValueInFront() SettingBool {
	return SettingBool{
		Name: "Value In Front",
		Desc: "Use Value In Front",
	}
}
