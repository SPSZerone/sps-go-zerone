package bag

import (
	"gioui.org/layout"
)

type (
	ListBagItemLength func(bagIndex int, bag *Bag) int
	ListBagItem       func(gtx layout.Context, bagIndex, itemIndex, itemSelectedIndex int, onClick func()) layout.Dimensions
	ListGridItem      func(gtx layout.Context, index, selectedIndex int, onClick func()) layout.Dimensions
	ItemDetail        func(gtx layout.Context, bagIndex int, bag *Bag, itemSelectedIndex int) layout.Dimensions
)
