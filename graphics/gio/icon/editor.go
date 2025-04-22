package icon

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var EditorMergeType = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.EditorMergeType)
	return icon
}()

var EditorModeEdit = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.EditorModeEdit)
	return icon
}()

var EditorPublish = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.EditorPublish)
	return icon
}()
