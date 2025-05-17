package item

import (
	"gioui.org/widget"

	spsclickable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/clickable"
	spsmenu "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/menu"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func NewUI() UI {
	return UI{
		Surface: spssurface.NewUniformInset(4, 4),
	}
}

type UI struct {
	Surface   spssurface.Surface
	Clickable spsclickable.Clickable
	Menu      spsmenu.Menu
	MenuItems []widget.Clickable
}
