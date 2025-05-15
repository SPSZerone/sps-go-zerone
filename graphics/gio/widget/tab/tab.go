package tab

import (
	"gioui.org/widget"
)

func New(name string) Tab {
	return Tab{
		Name: name,
	}
}

type Tab struct {
	Name      string
	Clickable widget.Clickable
	Data      any
}
