package page

import (
	"log"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/pref"
)

type Page interface {
	Actions() []component.AppBarAction
	Overflow() []component.OverflowAction
	Layout(gtx layout.Context, w *app.Window, th *material.Theme) layout.Dimensions
	NavItem() component.NavItem
}

type Pages struct {
	pages   map[any]Page
	current any

	*component.ModalNavDrawer
	NavAnim component.VisibilityAnimation
	*component.AppBar
	*component.ModalLayer

	Pref *pref.Preferences
}

func NewPages(pref *pref.Preferences) Pages {
	modal := component.NewModal()

	nav := component.NewNav("Navigation", "Enjoy!!")
	modalNav := component.ModalNavFrom(&nav, modal)

	bar := component.NewAppBar(modal)
	bar.NavigationIcon = icon.NavigationMenu

	na := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: time.Millisecond * 250,
	}
	return Pages{
		pages:          make(map[any]Page),
		ModalLayer:     modal,
		ModalNavDrawer: modalNav,
		AppBar:         bar,
		NavAnim:        na,
		Pref:           pref,
	}
}

func (p *Pages) Register(tag any, page Page) {
	p.pages[tag] = page
	navItem := page.NavItem()
	navItem.Tag = tag
	if p.current == nil {
		p.current = tag
		p.AppBar.Title = navItem.Name
		p.AppBar.SetActions(page.Actions(), page.Overflow())
	}
	p.ModalNavDrawer.AddNavItem(navItem)
}

func (p *Pages) SwitchTo(tag any) {
	page, ok := p.pages[tag]
	if !ok {
		return
	}
	navItem := page.NavItem()
	p.current = tag
	p.AppBar.Title = navItem.Name
	p.AppBar.SetActions(page.Actions(), page.Overflow())
}

func (p *Pages) Layout(gtx layout.Context, w *app.Window, th *material.Theme, deco func() layout.FlexChild) layout.Dimensions {
	// => AppBar
	for _, event := range p.AppBar.Events(gtx) {
		switch event := event.(type) {
		case component.AppBarNavigationClicked:
			if p.Pref.Settings.NonModalDrawer {
				p.NavAnim.ToggleVisibility(gtx.Now)
			} else {
				p.ModalNavDrawer.Appear(gtx.Now)
				p.NavAnim.Disappear(gtx.Now)
			}
		case component.AppBarContextMenuDismissed:
			log.Printf("Context menu dismissed: %v", event)
		case component.AppBarOverflowActionClicked:
			log.Printf("Overflow action selected: %v", event)
		}
	}

	// => ModalNav
	if p.ModalNavDrawer.NavDestinationChanged() {
		p.SwitchTo(p.ModalNavDrawer.CurrentNavDestination())
	}

	// => BG
	curIdx, ok := p.current.(int)
	if ok {
		color.Fill(gtx, color.DynamicColor(curIdx), color.DynamicColor(curIdx+1))
	} else {
		paint.Fill(gtx.Ops, th.Palette.Bg)
	}

	// => bar
	bar := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		thBar := *th
		colorBar := color.DynamicColor(3)
		thBar.ContrastBg = colorBar
		thBar.Palette.Bg = colorBar
		return p.AppBar.Layout(gtx, &thBar, "NavigationMenu", "Actions")
	})

	// => content
	content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		children := []layout.FlexChild{
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X /= 5
				return p.NavDrawer.Layout(gtx, th, &p.NavAnim)
			}),
		}
		if p.current != nil {
			children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return p.pages[p.current].Layout(gtx, w, th)
			}))
		}
		return layout.Flex{}.Layout(gtx, children...)
	})

	// => Final
	flex := layout.Flex{Axis: layout.Vertical}

	if p.Pref.Settings.Decorated {
		if p.Pref.Settings.BottomBar {
			flex.Layout(gtx, content, bar)
		} else {
			flex.Layout(gtx, bar, content)
		}
	} else {
		decorations := deco()
		if p.Pref.Settings.BottomBar {
			flex.Layout(gtx, decorations, content, bar)
		} else {
			flex.Layout(gtx, decorations, bar, content)
		}
	}

	p.ModalLayer.Layout(gtx, th)
	return layout.Dimensions{Size: gtx.Constraints.Max}
}
