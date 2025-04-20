package window

import (
	"log"
	"math"
	"time"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

const (
	DefaultNavRatio float32 = -0.8
)

func NewPages(win *Window) Pages {
	modalLayer := component.NewModal()

	navDrawer := component.NewNav("Navigation", "Enjoy!!")
	modalNavDrawer := component.ModalNavFrom(&navDrawer, modalLayer)

	appBar := component.NewAppBar(modalLayer)
	appBar.NavigationIcon = spsicon.NavigationMenu

	navAnim := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: time.Millisecond * 250,
	}
	split := spslayout.Split{
		Ratio: DefaultNavRatio,
	}
	return Pages{
		pages:          make(map[any]spsgio.Page),
		Window:         win,
		AppBar:         appBar,
		ModalLayer:     modalLayer,
		ModalNavDrawer: modalNavDrawer,
		NavAnim:        navAnim,
		Split:          split,
	}
}

type Pages struct {
	pages   map[any]spsgio.Page
	current any

	Window *Window

	AppBar         *component.AppBar
	ModalNavDrawer *component.ModalNavDrawer
	ModalLayer     *component.ModalLayer
	NavAnim        component.VisibilityAnimation
	Split          spslayout.Split
}

func (p *Pages) GetWindow() spsgio.Window {
	return p.Window
}

func (p *Pages) GetAppBar() *component.AppBar {
	return p.AppBar
}

func (p *Pages) GetModalNavDrawer() *component.ModalNavDrawer {
	return p.ModalNavDrawer
}

func (p *Pages) GetModalLayer() *component.ModalLayer {
	return p.ModalLayer
}

func (p *Pages) GetNavAnim() *component.VisibilityAnimation {
	return &p.NavAnim
}

func (p *Pages) Register(tag any, page spsgio.Page) {
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

func (p *Pages) SwitchTo(tag any) spsgio.Page {
	page, ok := p.pages[tag]
	if !ok {
		return nil
	}
	navItem := page.NavItem()
	p.current = tag
	p.AppBar.Title = navItem.Name
	p.AppBar.SetActions(page.Actions(), page.Overflow())
	return page
}

func (p *Pages) Current() spsgio.Page {
	return p.pages[p.current]
}

func (p *Pages) Count() int {
	return len(p.pages)
}

func (p *Pages) Start(tag any) spsgio.Page {
	page := p.SwitchTo(tag)

	timeNow := time.Now()
	if p.Window.Pref.Settings.ModalNavDrawer.Value {
		p.NavAnim.ToggleVisibility(timeNow)
	} else {
		p.ModalNavDrawer.Appear(timeNow)
		p.NavAnim.Disappear(timeNow)
	}
	p.ModalNavDrawer.SetNavDestination(tag)

	return page
}

func (p *Pages) OnEventPre(win spsgio.Window, evt event.Event, param any) {
	if p.current == nil {
		return
	}
	p.pages[p.current].OnEventPre(win, evt, param)
}

func (p *Pages) OnEventPost(win spsgio.Window, evt event.Event, param any) {
	if p.current == nil {
		return
	}
	p.pages[p.current].OnEventPost(win, evt, param)
}

func (p *Pages) Layout(win spsgio.Window, gtx layout.Context, param any, deco func() layout.FlexChild) layout.Dimensions {
	theme := win.GetTheme()
	pref := win.GetPref()

	totalWidth := gtx.Constraints.Max.X
	// => AppBar
	for _, evt := range p.AppBar.Events(gtx) {
		switch e := evt.(type) {
		case component.AppBarNavigationClicked:
			if pref.Settings.ModalNavDrawer.Value {
				p.NavAnim.ToggleVisibility(gtx.Now)
			} else {
				p.ModalNavDrawer.Appear(gtx.Now)
				p.NavAnim.Disappear(gtx.Now)
			}
		case component.AppBarContextMenuDismissed:
			log.Printf("Context menu dismissed: %v", e)
		case component.AppBarOverflowActionClicked:
			log.Printf("Overflow action selected: %v", e)
		}
	}

	// => ModalNav
	if p.ModalNavDrawer.NavDestinationChanged() {
		p.SwitchTo(p.ModalNavDrawer.CurrentNavDestination())
	}

	// => BG
	curIdx, ok := p.current.(int)
	if ok {
		spscolor.Fill(gtx, spscolor.DynamicColor(curIdx), spscolor.DynamicColor(curIdx+1))
	} else {
		paint.Fill(gtx.Ops, theme.Palette.Bg)
	}

	// => bar
	bar := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		thBar := *theme
		colorBar := spscolor.DynamicColor(3)
		thBar.ContrastBg = colorBar
		thBar.Palette.Bg = colorBar
		return p.AppBar.Layout(gtx, &thBar, "NavigationMenu", "Actions")
	})

	// => content
	content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		if p.Window.Pref.Settings.ModalNavDrawer.Value {
			if !p.NavAnim.Visible() {
				return p.pages[p.current].Layout(win, gtx, param)
			}
			return p.Split.Layout(
				gtx,
				func(gtx layout.Context) layout.Dimensions {
					if p.NavAnim.State == component.Disappearing {
						dimensions := p.ModalNavDrawer.NavDrawer.Layout(gtx, theme, &p.NavAnim)

						ratio := -(1 - float32(dimensions.Size.X)/float32(totalWidth>>1))
						p.Split.Ratio = ratio

						return dimensions
					}
					if p.NavAnim.State == component.Appearing {
						dimensions := p.ModalNavDrawer.NavDrawer.Layout(gtx, theme, &p.NavAnim)

						ratioValue := 1 - float32(math.Abs(float64(DefaultNavRatio)))
						gtxAnim := gtx
						gtxAnim.Constraints.Max.X = int(float32(totalWidth)*ratioValue) >> 1
						process := p.NavAnim.Revealed(gtxAnim) // 0.1 0.2 0.3 ...

						p.Split.Ratio = -(1 - ratioValue*process) // -0.95 -0.9 -0.85...
						return dimensions
					}

					return p.ModalNavDrawer.NavDrawer.Layout(gtx, theme, &p.NavAnim)
				},
				func(gtx layout.Context) layout.Dimensions {
					return p.pages[p.current].Layout(win, gtx, param)
				},
			)
		}

		children := []layout.FlexChild{
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.ModalNavDrawer.NavDrawer.Layout(gtx, theme, &p.NavAnim)
			}),
		}
		if p.current != nil {
			children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return p.pages[p.current].Layout(win, gtx, param)
			}))
		}
		return layout.Flex{}.Layout(gtx, children...)
	})

	// => Final
	flex := layout.Flex{Axis: layout.Vertical}

	if pref.Settings.Decorated.Value {
		if pref.Settings.BottomBar.Value {
			flex.Layout(gtx, content, bar)
		} else {
			flex.Layout(gtx, bar, content)
		}
	} else {
		decorations := deco()
		if pref.Settings.BottomBar.Value {
			flex.Layout(gtx, decorations, content, bar)
		} else {
			flex.Layout(gtx, decorations, bar, content)
		}
	}

	p.ModalLayer.Layout(gtx, theme)
	return layout.Dimensions{Size: gtx.Constraints.Max}
}
