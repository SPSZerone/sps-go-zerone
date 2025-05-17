package item

import (
	"gioui.org/widget"

	spsclickable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/clickable"
	spsmenu "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/menu"
)

func NewUI() UI {
	return UI{}
}

type UI struct {
	Clickable spsclickable.Clickable
	Menu      spsmenu.Menu
	MenuItems []widget.Clickable
}
