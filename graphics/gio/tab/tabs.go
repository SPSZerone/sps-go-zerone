package tab

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spsslice "github.com/SPSZerone/sps-go-zerone/generic/slice"
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
	list layout.List
	Tabs []Tab

	selected int
	slider   Slider
}

func (t *Tabs) AddTabByNames(names ...string) {
	for _, title := range names {
		t.Tabs = append(t.Tabs, Tab{Name: title})
	}
}

func (t *Tabs) AddTab(tabs ...Tab) {
	t.Tabs = append(t.Tabs, tabs...)
}

func (t *Tabs) UpdateTabs(cb func(idx int, tab *Tab) (stop bool)) {
	if cb == nil {
		return
	}
	for i := 0; i < len(t.Tabs); i++ {
		if cb(i, &t.Tabs[i]) {
			break
		}
	}
}

func (t *Tabs) GetTab(cb func(idx int, tab Tab) (ok bool)) *Tab {
	for i := 0; i < len(t.Tabs); i++ {
		if cb(i, t.Tabs[i]) {
			return &t.Tabs[i]
		}
	}
	return nil
}

func (t *Tabs) GetTabs() []Tab {
	return t.Tabs
}

func (t *Tabs) GetTabByIdx(index int) *Tab {
	if index < 0 || index >= len(t.Tabs) {
		return nil
	}
	return &t.Tabs[index]
}

func (t *Tabs) DelTab(cb func(idx int, tab Tab) (del bool)) {
	if len(t.Tabs) == 0 {
		return
	}

	delIndexes := make([]int, len(t.Tabs))
	for i, tab := range t.Tabs {
		if cb(i, tab) {
			delIndexes = append(delIndexes, i)
		}
	}

	for _, index := range delIndexes {
		t.DelTabByIndex(index)
	}
}

func (t *Tabs) DelTabByIndex(index int) {
	t.Tabs = spsslice.RemoveFast(t.Tabs, index)
	if t.selected >= len(t.Tabs) {
		t.SetFirstSelected()
	}
}

func (t *Tabs) UpdateTabName(index int, name string) {
	if index < 0 || index >= len(t.Tabs) {
		return
	}
	t.Tabs[index].Name = name
}

func (t *Tabs) UpdateTabData(index int, data any) {
	if index < 0 || index >= len(t.Tabs) {
		return
	}
	t.Tabs[index].Data = data
}

func (t *Tabs) Count() int {
	return len(t.Tabs)
}

func (t *Tabs) Selected() int {
	return t.selected
}

func (t *Tabs) SetSelected(index int) {
	if index < 0 || index >= len(t.Tabs) {
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
	app *spsgio.Application, gtx layout.Context, param any,
	content func(gtx layout.Context, selected int) layout.Dimensions,
) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Tabs
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
					// click area
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := material.Clickable(gtx, &tab.Btn, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(12)).Layout(gtx,
								material.H6(app.Theme, tab.Name).Layout,
							)
						})
						tabWidth = dims.Size.X
						return dims
					}),
					// highlight
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						if t.selected != tabIdx {
							return layout.Dimensions{}
						}

						highlightHeight := gtx.Dp(unit.Dp(4))
						highlightRect := image.Rect(0, 0, tabWidth, highlightHeight)
						paint.FillShape(gtx.Ops, app.Theme.Palette.ContrastBg, clip.Rect(highlightRect).Op())
						return layout.Dimensions{
							Size: image.Point{X: tabWidth, Y: highlightHeight},
						}
					}),
				)
			})
		}),
		// content
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return t.slider.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				spscolor.Fill(gtx, spscolor.DynamicColor(t.selected), spscolor.DynamicColor(t.selected+1))
				return content(gtx, t.selected)
			})
		}),
	)
}
