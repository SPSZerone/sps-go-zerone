package spacer

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

func NewSpacer() Spacer {
	return Spacer{
		layout.Spacer{Height: unit.Dp(8)},
	}
}

func NewWidthSpacer(value int) Spacer {
	if value <= 0 {
		value = 8
	}
	return Spacer{
		layout.Spacer{Width: unit.Dp(value)},
	}
}

func NewHeightSpacer(value int) Spacer {
	if value <= 0 {
		value = 8
	}
	return Spacer{
		layout.Spacer{Height: unit.Dp(value)},
	}
}

type Spacer struct {
	layout.Spacer
}

func (s Spacer) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Spacer{Height: unit.Dp(8)}.Layout(gtx)
}
