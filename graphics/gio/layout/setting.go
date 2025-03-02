package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

var DefaultInset = layout.UniformInset(unit.Dp(8))

type Setting struct {
	Key float32
	layout.Inset
}

func (d Setting) Layout(gtx layout.Context, key, value layout.Widget) layout.Dimensions {
	if d.Key == 0 {
		d.Key = 0.333
	}
	if d.Inset == (layout.Inset{}) {
		d.Inset = DefaultInset
	}
	return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Flexed(d.Key, func(gtx layout.Context) layout.Dimensions {
			return d.Inset.Layout(gtx, key)
		}),
		layout.Flexed(1-d.Key, func(gtx layout.Context) layout.Dimensions {
			return d.Inset.Layout(gtx, value)
		}),
	)
}
