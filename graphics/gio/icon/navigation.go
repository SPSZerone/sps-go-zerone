package icon

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var NavigationMenu = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()
