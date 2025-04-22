package tab

import (
	"gioui.org/widget"
)

type Tab struct {
	Name      string
	Clickable widget.Clickable
	Data      any
}
