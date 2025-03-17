package gio

import (
	"log"
	"time"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/x/component"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
)

type Tab interface {
	Actions() []component.AppBarAction
	Overflow() []component.OverflowAction
	NavItem() component.NavItem

	OnEventPre(app *Application, evt event.Event, param any)
	OnEventPost(app *Application, evt event.Event, param any)
	Layout(app *Application, gtx layout.Context, param any) layout.Dimensions
}

type Tabs struct {
	tabs    map[any]Tab
	current any

	*component.ModalNavDrawer
	NavAnim component.VisibilityAnimation
	*component.AppBar
	*component.ModalLayer
}

func NewTabs() Tabs {
	modal := component.NewModal()

	nav := component.NewNav("Navigation", "Enjoy!!")
	modalNav := component.ModalNavFrom(&nav, modal)

	bar := component.NewAppBar(modal)
	bar.NavigationIcon = icon.NavigationMenu

	na := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: time.Millisecond * 250,
	}
	return Tabs{
		tabs:           make(map[any]Tab),
		ModalLayer:     modal,
		ModalNavDrawer: modalNav,
		AppBar:         bar,
		NavAnim:        na,
	}
}

func (t *Tabs) Register(tag any, tab Tab) {
	t.tabs[tag] = tab
	navItem := tab.NavItem()
	navItem.Tag = tag
	if t.current == nil {
		t.current = tag
		t.AppBar.Title = navItem.Name
		t.AppBar.SetActions(tab.Actions(), tab.Overflow())
	}
	t.ModalNavDrawer.AddNavItem(navItem)
}

func (t *Tabs) SwitchTo(tag any) {
	page, ok := t.tabs[tag]
	if !ok {
		return
	}
	navItem := page.NavItem()
	t.current = tag
	t.AppBar.Title = navItem.Name
	t.AppBar.SetActions(page.Actions(), page.Overflow())
}

func (t *Tabs) OnEventPre(app *Application, evt event.Event, param any) {
	if t.current == nil {
		return
	}
	t.tabs[t.current].OnEventPre(app, evt, param)
}

func (t *Tabs) OnEventPost(app *Application, evt event.Event, param any) {
	if t.current == nil {
		return
	}
	t.tabs[t.current].OnEventPost(app, evt, param)
}

func (t *Tabs) Layout(app *Application, gtx layout.Context, param any, deco func() layout.FlexChild) layout.Dimensions {
	// => AppBar
	for _, evt := range t.AppBar.Events(gtx) {
		switch e := evt.(type) {
		case component.AppBarNavigationClicked:
			if app.Pref.Settings.NonModalDrawer {
				t.NavAnim.ToggleVisibility(gtx.Now)
			} else {
				t.ModalNavDrawer.Appear(gtx.Now)
				t.NavAnim.Disappear(gtx.Now)
			}
		case component.AppBarContextMenuDismissed:
			log.Printf("Context menu dismissed: %v", e)
		case component.AppBarOverflowActionClicked:
			log.Printf("Overflow action selected: %v", e)
		}
	}

	// => ModalNav
	if t.ModalNavDrawer.NavDestinationChanged() {
		t.SwitchTo(t.ModalNavDrawer.CurrentNavDestination())
	}

	// => BG
	curIdx, ok := t.current.(int)
	if ok {
		color.Fill(gtx, color.DynamicColor(curIdx), color.DynamicColor(curIdx+1))
	} else {
		paint.Fill(gtx.Ops, app.Theme.Palette.Bg)
	}

	// => bar
	bar := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		thBar := *app.Theme
		colorBar := color.DynamicColor(3)
		thBar.ContrastBg = colorBar
		thBar.Palette.Bg = colorBar
		return t.AppBar.Layout(gtx, &thBar, "NavigationMenu", "Actions")
	})

	// => content
	content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		children := []layout.FlexChild{
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X /= 5
				return t.NavDrawer.Layout(gtx, app.Theme, &t.NavAnim)
			}),
		}
		if t.current != nil {
			children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return t.tabs[t.current].Layout(app, gtx, param)
			}))
		}
		return layout.Flex{}.Layout(gtx, children...)
	})

	// => Final
	flex := layout.Flex{Axis: layout.Vertical}

	if app.Pref.Settings.Decorated {
		if app.Pref.Settings.BottomBar {
			flex.Layout(gtx, content, bar)
		} else {
			flex.Layout(gtx, bar, content)
		}
	} else {
		decorations := deco()
		if app.Pref.Settings.BottomBar {
			flex.Layout(gtx, decorations, content, bar)
		} else {
			flex.Layout(gtx, decorations, bar, content)
		}
	}

	t.ModalLayer.Layout(gtx, app.Theme)
	return layout.Dimensions{Size: gtx.Constraints.Max}
}
