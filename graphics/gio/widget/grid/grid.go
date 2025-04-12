package grid

import (
	"gioui.org/layout"
	"gioui.org/x/outlay"
)

func NewGrid() Grid {
	return Grid{
		Axis:      layout.Horizontal,
		Alignment: layout.End,
		Num:       10,
	}
}

type Grid struct {
	Axis      layout.Axis
	Alignment layout.Alignment
	Wrap      bool
	Num       int
}

func (g *Grid) Layout(gtx layout.Context, num int, element outlay.FlowElement) layout.Dimensions {
	if g.Wrap {
		return outlay.FlowWrap{
			Axis:      g.Axis,
			Alignment: g.Alignment,
		}.Layout(gtx, num, element)
	}

	flow := outlay.Flow{
		Axis:      g.Axis,
		Alignment: g.Alignment,
		Num:       g.Num,
	}
	return flow.Layout(gtx, num, element)
}
