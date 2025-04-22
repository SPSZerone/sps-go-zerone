package icon

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var FileFolder = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.FileFolder)
	return icon
}()

var FileFolderOpen = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.FileFolderOpen)
	return icon
}()
