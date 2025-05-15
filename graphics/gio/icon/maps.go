package icon

import (
	"golang.org/x/exp/shiny/materialdesign/icons"

	"gioui.org/widget"
)

var MapsRestaurantMenu *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.MapsRestaurantMenu)
	return icon
}()
