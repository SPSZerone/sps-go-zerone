package tab

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spsslice "github.com/SPSZerone/sps-go-zerone/generic/slice"
	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	spslist "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/list"
)

func NewTabsByNames(names []string, opts ...TabsOption) Tabs {
	t := NewTabs(opts...)
	t.AddTabByNames(names...)
	return t
}

func NewTabsByTabs(tabs []Tab, opts ...TabsOption) Tabs {
	t := NewTabs(opts...)
	t.AddTab(tabs...)
	return t
}

func NewTabs(opts ...TabsOption) Tabs {
	t := newTabs()
	t.Update(opts...)
	return t
}

func newTabs() Tabs {
	return Tabs{
		Opts: NewTabsOptions(),
		List: spslist.New(),
	}
}

type Content func(gtx layout.Context, selected int) layout.Dimensions

type Tabs struct {
	Opts TabsOptions

	List spslist.List
	Tabs []Tab

	selected int
	Slider   Slider
}

func (t *Tabs) Update(opts ...TabsOption) {
	for _, opt := range opts {
		opt(&t.Opts)
	}
}

func (t *Tabs) AddTabByNames(names ...string) {
	for _, name := range names {
		t.addTab(New(name, nil))
	}
}

func (t *Tabs) AddTab(tabs ...Tab) {
	for _, tab := range tabs {
		t.addTab(tab)
	}
}

func (t *Tabs) addTab(tab Tab) {
	if tab.Opts.CloseModeInheritTabs {
		tab.Opts.CloseMode = t.Opts.CloseMode
	}
	t.Tabs = append(t.Tabs, tab)
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

func (t *Tabs) GetTab(cb func(idx int, tab *Tab) (ok bool)) *Tab {
	for i := 0; i < len(t.Tabs); i++ {
		if cb(i, &t.Tabs[i]) {
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

func (t *Tabs) DelTab(cb func(idx int, tab *Tab) (del bool)) {
	if len(t.Tabs) == 0 {
		return
	}

	delIndexes := make([]int, 0, len(t.Tabs))
	for i, tab := range t.Tabs {
		if cb(i, &tab) {
			delIndexes = append(delIndexes, i)
		}
	}

	for i := len(delIndexes) - 1; i >= 0; i-- {
		idx := delIndexes[i]
		t.Tabs = append(t.Tabs[:idx], t.Tabs[idx+1:]...)
	}
	if t.selected >= len(t.Tabs) {
		t.SetFirstSelected()
	}
}

func (t *Tabs) DelTabByIndex(index int) {
	t.Tabs = spsslice.RemoveByKeepOrder(t.Tabs, index)
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
		// Content
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

	// WidthLimitWhenVertical
	var widthLimit int
	if t.Opts.WidthLimitWhenVertical > 0 {
		if t.IsVertical() {
			widthLimit = gtx.Dp(unit.Dp(t.Opts.WidthLimitWhenVertical))
		}
	}

	t.DelTab(func(idx int, tab *Tab) (del bool) {
		return tab.IsClose()
	})
	return t.List.Layout(theme, gtx, t.Count(), func(gtx layout.Context, tabIdx int) layout.Dimensions {
		tab := t.GetTabByIdx(tabIdx)
		highlight := t.selected == tabIdx
		d, clicked := tab.LayoutDefault(
			theme, gtx,
			highlight,
			t.GetLayoutAxis(),
			widthLimit,
			t.Opts.ColorfulBG,
		)
		if clicked {
			if t.selected < tabIdx {
				t.Slider.PushLeft()
			} else if t.selected > tabIdx {
				t.Slider.PushRight()
			}
			t.selected = tabIdx
		}
		return d
	})
}
