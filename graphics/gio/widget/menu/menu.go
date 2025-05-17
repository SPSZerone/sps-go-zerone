package menu

import (
	"image"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func New() Menu {
	return Menu{}
}

type Menu struct {
	MenuState   component.MenuState
	ContextArea component.ContextArea
}

func (m *Menu) AddWidgets(widgets ...func(gtx layout.Context) layout.Dimensions) *Menu {
	m.MenuState.Options = append(m.MenuState.Options, widgets...)
	return m
}

func (m *Menu) SetWidgets(widgets ...func(gtx layout.Context) layout.Dimensions) *Menu {
	m.MenuState.Options = widgets
	return m
}

func (m *Menu) Layout(
	theme *material.Theme, gtx layout.Context,
	alignment layout.Direction,
	limitContextAreaSize image.Point,
	widget func(gtx layout.Context) layout.Dimensions,
) layout.Dimensions {
	return layout.Stack{Alignment: alignment}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			surface := spssurface.NewDefault()
			return surface.LayoutDefault(theme, gtx, func(gtx layout.Context) layout.Dimensions {
				if widget == nil {
					return layout.Dimensions{}
				}
				return widget(gtx)
			})
		}),
		m.LayoutExpandedContextArea(theme, limitContextAreaSize),
	)
}

func (m *Menu) LayoutContextArea(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return m.ContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min = image.Point{}
		return component.Menu(theme, &m.MenuState).Layout(gtx)
	})
}

func (m *Menu) LayoutExpandedContextArea(theme *material.Theme, limitSize image.Point) layout.StackChild {
	return layout.Expanded(func(gtx layout.Context) layout.Dimensions {
		if limitSize.X > 0 {
			gtx.Constraints.Max.X = limitSize.X
		}
		if limitSize.Y > 0 {
			gtx.Constraints.Max.Y = limitSize.Y
		}
		return m.LayoutContextArea(theme, gtx)
	})
}
