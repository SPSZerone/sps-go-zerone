package bag

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	spsgrid "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/grid"
)

func NewGrid() Grid {
	g := spsgrid.New()
	g.Num = 4
	return Grid{
		Grid: g,
	}
}

type Grid struct {
	Grid spsgrid.Grid

	CurSelectedIndex int
}

func (g *Grid) Layout(
	theme *material.Theme, gtx layout.Context,
	length int, listItem ListGridItem,
) layout.Dimensions {
	return g.Grid.Layout(theme, gtx, length, func(gtx layout.Context, index int) layout.Dimensions {
		return listItem(gtx, index, g.CurSelectedIndex, func() {
			g.CurSelectedIndex = index
		})
	})
}
