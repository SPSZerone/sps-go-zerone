package tab

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spsslice "github.com/SPSZerone/sps-go-zerone/generic/slice"
	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

func NewTabsByNames(names []string, opts ...Option) Tabs {
	t := NewTabs(opts...)
	t.AddTabByNames(names...)
	return t
}

func NewTabsByTabs(tabs []Tab, opts ...Option) Tabs {
	t := NewTabs(opts...)
	t.AddTab(tabs...)
	return t
}

func NewTabs(opts ...Option) Tabs {
	t := newTabs()
	t.Update(opts...)
	return t
}

func newTabs() Tabs {
	return Tabs{
		Opts: NewOptions(),
		List: widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

type Content func(gtx layout.Context, selected int) layout.Dimensions

type Tabs struct {
	Opts Options

	List widget.List
	Tabs []Tab

	selected int
	Slider   Slider
}

func (t *Tabs) Update(opts ...Option) {
	for _, opt := range opts {
		opt(&t.Opts)
	}
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

func (t *Tabs) GetSelected() int {
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

func (t *Tabs) GetLayoutAxis() layout.Axis {
	if t.Opts.AxisSetting != nil {
		return t.Opts.AxisSetting.GetAxis()
	}
	return t.Opts.Axis
}

func (t *Tabs) IsHorizontal() bool {
	return t.GetLayoutAxis() == layout.Horizontal
}

func (t *Tabs) IsVertical() bool {
	return t.GetLayoutAxis() == layout.Vertical
}

func (t *Tabs) Layout(theme *material.Theme, gtx layout.Context, param any, content Content) layout.Dimensions {
	axis := layout.Horizontal
	if t.IsHorizontal() {
		axis = layout.Vertical
	}
	return layout.Flex{Axis: axis}.Layout(gtx,
		// Tabs
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.LayoutTabs(theme, gtx, param)
		}),
		// LayoutContent
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return t.doLayoutContent(theme, gtx, param, content)
		}),
	)
}

func (t *Tabs) doLayoutContent(theme *material.Theme, gtx layout.Context, param any, content Content) layout.Dimensions {
	return t.Slider.Layout(t.GetLayoutAxis(), gtx, func(gtx layout.Context) layout.Dimensions {
		spscolor.Fill(gtx, spscolor.DynamicColor(t.selected), spscolor.DynamicColor(t.selected+1))
		return content(gtx, t.selected)
	})
}

func (t *Tabs) LayoutTabs(theme *material.Theme, gtx layout.Context, param any) layout.Dimensions {
	// Axis
	t.List.Axis = t.GetLayoutAxis()

	// WidthWhenVertical
	var width int
	if t.Opts.WidthWhenVertical > 0 {
		if t.IsVertical() {
			width = gtx.Dp(unit.Dp(t.Opts.WidthWhenVertical))
		}
	}

	return material.List(theme, &t.List).Layout(gtx, t.Count(), func(gtx layout.Context, tabIdx int) layout.Dimensions {
		tab := t.GetTabByIdx(tabIdx)
		if tab.Clickable.Clicked(gtx) {
			if t.selected < tabIdx {
				t.Slider.PushLeft()
			} else if t.selected > tabIdx {
				t.Slider.PushRight()
			}
			t.selected = tabIdx
		}

		var tabSize image.Point
		return layout.Stack{Alignment: layout.S}.Layout(gtx,
			// click area
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				if width > 0 {
					if t.IsVertical() {
						gtx.Constraints.Max.X = width
					}
				}

				return material.Clickable(gtx, &tab.Clickable, func(gtx layout.Context) layout.Dimensions {
					dims := layout.UniformInset(unit.Dp(12)).Layout(
						gtx,
						func(gtx layout.Context) layout.Dimensions {
							name := material.H6(theme, tab.Name)
							return name.Layout(gtx)
						},
					)
					tabSize = dims.Size
					if width > 0 {
						if t.IsVertical() {
							tabSize.X = width
						}
					}
					return layout.Dimensions{Size: tabSize}
				})
			}),
			// highlight
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				if t.selected != tabIdx {
					return layout.Dimensions{}
				}

				highlightThickness := gtx.Dp(unit.Dp(4))
				var highlightRect image.Rectangle
				var size image.Point
				if t.GetLayoutAxis() == layout.Vertical {
					size = image.Pt(highlightThickness, tabSize.Y)
					// right
					startX := tabSize.X>>1 - highlightThickness>>1
					// left
					//startX = -startX
					highlightRect = image.Rect(startX, 0, startX+highlightThickness, tabSize.Y)
				} else {
					size = image.Pt(tabSize.X, highlightThickness)
					highlightRect = image.Rect(0, 0, tabSize.X, highlightThickness)
				}

				paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(highlightRect).Op())
				return layout.Dimensions{Size: size}
			}),
		)
	})
}
