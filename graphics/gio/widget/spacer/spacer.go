package spacer

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

func NewWidthSpacer(value int) Spacer {
	return Spacer{
		layout.Spacer{Width: unit.Dp(value)},
	}
}

func NewHeightSpacer(value int) Spacer {
	return Spacer{
		layout.Spacer{Height: unit.Dp(value)},
	}
}

type Spacer struct {
	layout.Spacer
}

func (s Spacer) Layout(gtx layout.Context) layout.Dimensions {
	return s.Spacer.Layout(gtx)
}
