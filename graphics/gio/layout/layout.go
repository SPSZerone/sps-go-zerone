package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Setting struct {
	Key float32
	layout.Inset
}

var DefaultInset = layout.UniformInset(unit.Dp(8))

func (d Setting) Layout(gtx C, key, value layout.Widget) D {
	if d.Key == 0 {
		d.Key = 0.333
	}
	if d.Inset == (layout.Inset{}) {
		d.Inset = DefaultInset
	}
	return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Flexed(d.Key, func(gtx C) D {
			return d.Inset.Layout(gtx, key)
		}),
		layout.Flexed(1-d.Key, func(gtx C) D {
			return d.Inset.Layout(gtx, value)
		}),
	)
}
