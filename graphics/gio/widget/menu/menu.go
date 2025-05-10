package menu

import (
	"image"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type Menu struct {
	MenuState   component.MenuState
	ContextArea component.ContextArea
}

func (m *Menu) AddWidgets(widgets ...func(gtx layout.Context) layout.Dimensions) {
	m.MenuState.Options = append(m.MenuState.Options, widgets...)
}

func (m *Menu) LayoutContextArea(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return m.ContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min = image.Point{}
		return component.Menu(theme, &m.MenuState).Layout(gtx)
	})
}

func (m *Menu) LayoutExpandedContextArea(theme *material.Theme, gtx layout.Context) layout.StackChild {
	return layout.Expanded(func(gtx layout.Context) layout.Dimensions {
		return m.LayoutContextArea(theme, gtx)
	})
}
