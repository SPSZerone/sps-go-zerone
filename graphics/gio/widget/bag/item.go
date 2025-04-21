package bag

import (
	"time"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ItemDataSource func(bagIndex int) (items []Item, updateTime time.Time)

type Item interface {
	Layout(theme *material.Theme, gtx layout.Context, highlight bool) (dimensions layout.Dimensions, clicked bool)
	LayoutDetail(theme *material.Theme, gtx layout.Context) layout.Dimensions
}
