package tab

import (
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

type Tabs struct {
	list     layout.List
	tabs     []Tab
	selected int
	slider   Slider
}

type Tab struct {
	btn   widget.Clickable
	Title string
}

func (t *Tabs) AddTab(title string) {
	t.tabs = append(t.tabs, Tab{Title: title})
}

func (t *Tabs) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.list.Layout(gtx, len(t.tabs), func(gtx layout.Context, tabIdx int) layout.Dimensions {
				tab := &t.tabs[tabIdx]
				if tab.btn.Clicked(gtx) {
					if t.selected < tabIdx {
						t.slider.PushLeft()
					} else if t.selected > tabIdx {
						t.slider.PushRight()
					}
					t.selected = tabIdx
				}
				var tabWidth int
				return layout.Stack{Alignment: layout.S}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := material.Clickable(gtx, &tab.btn, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(12)).Layout(gtx,
								material.H6(theme, tab.Title).Layout,
							)
						})
						tabWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						if t.selected != tabIdx {
							return layout.Dimensions{}
						}
						tabHeight := gtx.Dp(unit.Dp(4))
						tabRect := image.Rect(0, 0, tabWidth, tabHeight)
						paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(tabRect).Op())
						return layout.Dimensions{
							Size: image.Point{X: tabWidth, Y: tabHeight},
						}
					}),
				)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return t.slider.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				spscolor.Fill(gtx, spscolor.DynamicColor(t.selected), spscolor.DynamicColor(t.selected+1))
				return layout.Center.Layout(gtx,
					material.H1(theme, fmt.Sprintf("Tab content #%d", t.selected+1)).Layout,
				)
			})
		}),
	)
}
