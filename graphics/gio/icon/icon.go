package icon

import (
	"golang.org/x/exp/shiny/materialdesign/icons"

	"gioui.org/widget"
)

var NavMenu = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()

var ActSettings = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionSettings)
	return icon
}()

var ActHelp = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionHelp)
	return icon
}()
