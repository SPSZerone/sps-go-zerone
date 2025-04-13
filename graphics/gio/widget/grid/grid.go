package grid

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func NewGrid() Grid {
	return Grid{
		list: widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},

		Axis:      layout.Horizontal,
		Alignment: layout.Start,
		Num:       5,
	}
}

type Grid struct {
	list widget.List

	Axis      layout.Axis
	Alignment layout.Alignment
	Wrap      bool
	Num       int
}

func (g *Grid) Layout(app *spsgio.Application, gtx layout.Context, num int, element outlay.FlowElement) layout.Dimensions {
	return material.List(app.Theme, &g.list).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return g.LayoutContent(app, gtx, num, element)
	})
}

func (g *Grid) LayoutContent(app *spsgio.Application, gtx layout.Context, num int, element outlay.FlowElement) layout.Dimensions {
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
