package bag

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	ListBagItemLength func(bagIndex int, bag *Bag) int
	ListBagItem       func(gtx layout.Context, bagIndex, itemIndex, itemSelectedIndex int, onClick func()) layout.Dimensions
	ListGridItem      func(gtx layout.Context, index, selectedIndex int, onClick func()) layout.Dimensions
	ItemDetail        func(gtx layout.Context, bagIndex int, bag *Bag, itemSelectedIndex int) layout.Dimensions
)

type Item interface {
	Layout(theme *material.Theme, gtx layout.Context, highlight bool, onClick func()) layout.Dimensions
	LayoutDetail(theme *material.Theme, gtx layout.Context) layout.Dimensions
}
