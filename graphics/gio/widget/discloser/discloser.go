package discloser

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func New(appear bool) Discloser {
	state := component.DiscloserState{}
	if appear {
		state.Click()
	}
	return Discloser{
		DiscloserState: state,
	}
}

type Discloser struct {
	DiscloserState component.DiscloserState
}

func (d *Discloser) Layout(
	theme *material.Theme, gtx layout.Context,
	controlInvisible, controlVisible string,
	summary, detail layout.Widget,
) layout.Dimensions {
	discloserState := &d.DiscloserState
	return component.Discloser(theme, discloserState).Layout(
		gtx,
		// control
		func(gtx layout.Context) layout.Dimensions {
			labelStyle := material.Body1(theme, controlInvisible)
			if discloserState.Visible() {
				labelStyle.Text = controlVisible
			}
			return layout.UniformInset(unit.Dp(4)).Layout(gtx, labelStyle.Layout)
		},
		// summary
		summary,
		// detail
		detail,
	)
}

func (d *Discloser) LayoutWithDefaultControl(
	theme *material.Theme, gtx layout.Context,
	summary, detail layout.Widget,
) layout.Dimensions {
	return d.Layout(theme, gtx, ` 󰞘 ` /* 󰞘 */, ` 󰞖 `, summary, detail)
}
