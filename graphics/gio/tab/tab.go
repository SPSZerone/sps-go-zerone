package tab

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

func NewTabs(titles ...string) Tabs {
	t := Tabs{}
	t.AddTab(titles...)
	return t
}

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

func (t *Tabs) AddTab(titles ...string) {
	for _, title := range titles {
		t.tabs = append(t.tabs, Tab{Title: title})
	}
}

func (t *Tabs) UpdateTabTitle(index int, title string) {
	if index < 0 || index >= len(t.tabs) {
		return
	}
	t.tabs[index].Title = title
}

func (t *Tabs) Selected() int {
	return t.selected
}

func (t *Tabs) SetSelected(index int) {
	if index < 0 || index >= len(t.tabs) {
		return
	}
	t.selected = index
}

func (t *Tabs) Layout(
	application *spsgio.Application, gtx layout.Context, param any,
	content func(gtx layout.Context, selected int) layout.Dimensions,
) layout.Dimensions {
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
								material.H6(application.Theme, tab.Title).Layout,
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
						paint.FillShape(gtx.Ops, application.Theme.Palette.ContrastBg, clip.Rect(tabRect).Op())
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
				return content(gtx, t.selected)
			})
		}),
	)
}
