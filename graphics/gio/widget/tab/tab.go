package tab

import (
	spsclickable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/clickable"
)

type Tab struct {
	Name      string
	Clickable spsclickable.Clickable
	Data      any
}
