package gio

import (
	"log"
	"math"
	"time"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/x/component"

	"github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

const (
	DefaultNavRatio float32 = -0.8
)

func NewPages(app *Application) Pages {
	modalLayer := component.NewModal()

	navDrawer := component.NewNav("Navigation", "Enjoy!!")
	modalNavDrawer := component.ModalNavFrom(&navDrawer, modalLayer)

	appBar := component.NewAppBar(modalLayer)
	appBar.NavigationIcon = icon.NavigationMenu

	navAnim := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: time.Millisecond * 250,
	}
	split := spslayout.Split{
		Ratio: DefaultNavRatio,
	}
	return Pages{
		pages:          make(map[any]Page),
		App:            app,
		AppBar:         appBar,
		ModalLayer:     modalLayer,
		ModalNavDrawer: modalNavDrawer,
		NavAnim:        navAnim,
		split:          split,
	}
}

type Page interface {
	Actions() []component.AppBarAction
	Overflow() []component.OverflowAction
	NavItem() component.NavItem

	OnEventPre(app *Application, evt event.Event, param any)
	OnEventPost(app *Application, evt event.Event, param any)
	Layout(app *Application, gtx layout.Context, param any) layout.Dimensions
}

type Pages struct {
	pages   map[any]Page
	current any

	App *Application

	AppBar         *component.AppBar
	ModalNavDrawer *component.ModalNavDrawer
	ModalLayer     *component.ModalLayer
	NavAnim        component.VisibilityAnimation
	split          spslayout.Split
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

func (p *Pages) SwitchTo(tag any) Page {
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

func (p *Pages) Start(tag any) Page {
	page := p.SwitchTo(tag)

	timeNow := time.Now()
	if p.App.Pref.Settings.ModalNavDrawer.Value {
		p.NavAnim.ToggleVisibility(timeNow)
	} else {
		p.ModalNavDrawer.Appear(timeNow)
		p.NavAnim.Disappear(timeNow)
	}
	p.ModalNavDrawer.SetNavDestination(tag)

	return page
}

func (p *Pages) OnEventPre(app *Application, evt event.Event, param any) {
	if p.current == nil {
		return
	}
	p.pages[p.current].OnEventPre(app, evt, param)
}

func (p *Pages) OnEventPost(app *Application, evt event.Event, param any) {
	if p.current == nil {
		return
	}
	p.pages[p.current].OnEventPost(app, evt, param)
}

func (p *Pages) Layout(app *Application, gtx layout.Context, param any, deco func() layout.FlexChild) layout.Dimensions {
	totalWidth := gtx.Constraints.Max.X
	// => AppBar
	for _, evt := range p.AppBar.Events(gtx) {
		switch e := evt.(type) {
		case component.AppBarNavigationClicked:
			if app.Pref.Settings.ModalNavDrawer.Value {
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
		return p.AppBar.Layout(gtx, &thBar, "NavigationMenu", "Actions")
	})

	// => content
	content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		if p.App.Pref.Settings.ModalNavDrawer.Value {
			if !p.NavAnim.Visible() {
				return p.pages[p.current].Layout(app, gtx, param)
			}
			return p.split.Layout(
				gtx,
				func(gtx layout.Context) layout.Dimensions {
					if p.NavAnim.State == component.Disappearing {
						dimensions := p.ModalNavDrawer.NavDrawer.Layout(gtx, app.Theme, &p.NavAnim)

						ratio := -(1 - float32(dimensions.Size.X)/float32(totalWidth>>1))
						app.Logger.Info().Msgf("%v %v %v", ratio, totalWidth, dimensions)
						p.split.Ratio = ratio

						return dimensions
					}
					if p.NavAnim.State == component.Appearing {
						dimensions := p.ModalNavDrawer.NavDrawer.Layout(gtx, app.Theme, &p.NavAnim)

						ratioValue := 1 - float32(math.Abs(float64(DefaultNavRatio)))
						gtxAnim := gtx
						gtxAnim.Constraints.Max.X = int(float32(totalWidth)*ratioValue) >> 1
						process := p.NavAnim.Revealed(gtxAnim) // 0.1 0.2 0.3 ...

						p.split.Ratio = -(1 - ratioValue*process) // -0.95 -0.9 -0.85...
						return dimensions
					}

					return p.ModalNavDrawer.NavDrawer.Layout(gtx, app.Theme, &p.NavAnim)
				},
				func(gtx layout.Context) layout.Dimensions {
					return p.pages[p.current].Layout(app, gtx, param)
				},
			)
		}

		children := []layout.FlexChild{
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = app.GetNavigationWidth(gtx)
				return p.ModalNavDrawer.NavDrawer.Layout(gtx, app.Theme, &p.NavAnim)
			}),
		}
		if p.current != nil {
			children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return p.pages[p.current].Layout(app, gtx, param)
			}))
		}
		return layout.Flex{}.Layout(gtx, children...)
	})

	// => Final
	flex := layout.Flex{Axis: layout.Vertical}

	if app.Pref.Settings.Decorated.Value {
		if app.Pref.Settings.BottomBar.Value {
			flex.Layout(gtx, content, bar)
		} else {
			flex.Layout(gtx, bar, content)
		}
	} else {
		decorations := deco()
		if app.Pref.Settings.BottomBar.Value {
			flex.Layout(gtx, decorations, content, bar)
		} else {
			flex.Layout(gtx, decorations, bar, content)
		}
	}

	p.ModalLayer.Layout(gtx, app.Theme)
	return layout.Dimensions{Size: gtx.Constraints.Max}
}
