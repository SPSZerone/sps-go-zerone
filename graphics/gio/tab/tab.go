package tab

import (
	"gioui.org/widget"
)

type Tab struct {
	Name string
	Btn  widget.Clickable
	Data any
}
