package tab

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/SPSZerone/sps-go-zerone/generic/slice"
	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

func NewTabsByNames(names ...string) Tabs {
	t := Tabs{}
	t.AddTabByNames(names...)
	return t
}

func NewTabs(tabs ...Tab) Tabs {
	t := Tabs{}
	t.AddTab(tabs...)
	return t
}

type Tabs struct {
	list     layout.List
	tabs     []Tab
	selected int
	slider   Slider
}

func (t *Tabs) AddTabByNames(names ...string) {
	for _, title := range names {
		t.tabs = append(t.tabs, Tab{Name: title})
	}
}

func (t *Tabs) AddTab(tabs ...Tab) {
	t.tabs = append(t.tabs, tabs...)
}

func (t *Tabs) UpdateTabs(cb func(idx int, tab *Tab) (stop bool)) {
	if cb == nil {
		return
	}
	for i := 0; i < len(t.tabs); i++ {
		if cb(i, &t.tabs[i]) {
			break
		}
	}
}

func (t *Tabs) GetTab(cb func(idx int, tab Tab) (ok bool)) *Tab {
	for i := 0; i < len(t.tabs); i++ {
		if cb(i, t.tabs[i]) {
			return &t.tabs[i]
		}
	}
	return nil
}

func (t *Tabs) GetTabByIdx(index int) *Tab {
	if index < 0 || index >= len(t.tabs) {
		return nil
	}
	return &t.tabs[index]
}

func (t *Tabs) DelTab(cb func(idx int, tab Tab) (del bool)) {
	delIdx := -1
	for i, tab := range t.tabs {
		if cb(i, tab) {
			delIdx = i
			break
		}
	}
	t.DelTabByIndex(delIdx)
}

func (t *Tabs) DelTabByIndex(index int) {
	t.tabs = slice.RemoveFast(t.tabs, index)
	if t.selected >= len(t.tabs) {
		t.SetFirstSelected()
	}
}

func (t *Tabs) UpdateTabName(index int, name string) {
	if index < 0 || index >= len(t.tabs) {
		return
	}
	t.tabs[index].Name = name
}

func (t *Tabs) Count() int {
	return len(t.tabs)
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

func (t *Tabs) SetFirstSelected() int {
	t.selected = 0
	return t.selected
}

func (t *Tabs) SetLastSelected() int {
	count := t.Count()
	if count <= 0 {
		return t.selected
	}
	t.selected = count - 1
	return t.selected
}

func (t *Tabs) Layout(
	application *spsgio.Application, gtx layout.Context, param any,
	content func(gtx layout.Context, selected int) layout.Dimensions,
) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.list.Layout(gtx, t.Count(), func(gtx layout.Context, tabIdx int) layout.Dimensions {
				tab := t.GetTabByIdx(tabIdx)
				if tab.Btn.Clicked(gtx) {
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
						dims := material.Clickable(gtx, &tab.Btn, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(12)).Layout(gtx,
								material.H6(application.Theme, tab.Name).Layout,
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
