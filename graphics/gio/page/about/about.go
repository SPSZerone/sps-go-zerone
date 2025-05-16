package about

import (
	"image/color"

	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsicon "github.com/SPSZerone/sps-go-zerone/graphics/gio/icon"
	spscopy "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/copy"
	spslist "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/list"
	spstreenode "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/treenode"
)

const (
	Name     = "SPS Gio Framework base on Gio"
	Author   = "SPSZerone"
	License  = "GPLv3"
	HomePage = "https://github.com/SPSZerone"
)

func New(pages spsgio.Pages) *Page {
	p := &Page{
		Pages:    pages,
		TreeNode: TestTreeNode(),
	}
	pages.RegisterAppBarEvent(p, p.OnAppBarEvent)
	return p
}

func TestTreeNode() spstreenode.TreeNode {
	return spstreenode.TreeNode{
		Text: "Expand Me",
		Children: []spstreenode.TreeNode{
			{
				Text: "Disclosers can be (expand me)...",
				Children: []spstreenode.TreeNode{
					{
						Text: "...nested to arbitrary depths.",
					},
					{
						Text: "There are also types available to customize the look and feel of the discloser:",
						Children: []spstreenode.TreeNode{
							{
								Text: "• DiscloserStyle lets you provide your own control instead of the default triangle used here.",
							},
							{
								Text: "• DiscloserArrowStyle lets you alter the presentation of the triangle used here, like changing its color, size, left/right anchoring, or margin.",
							},
						},
					},
				},
			},
		},
	}
}

var _ spsgio.Page = (*Page)(nil)

type Contextual struct {
	Clickable widget.Clickable

	AppBarActionFavorite widget.Clickable
	OverflowActionA      widget.Clickable
	OverflowActionB      widget.Clickable
}

type Page struct {
	spslist.List
	spsgio.Pages

	CopyName     spscopy.Copy
	CopyAuthor   spscopy.Copy
	CopyLicense  spscopy.Copy
	CopyHomePage spscopy.Copy

	Contextual Contextual
	TreeNode   spstreenode.TreeNode
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "About",
		Icon: spsicon.ActionHelp,
	}
}

func (p *Page) OnAppBarEvent(pages spsgio.Pages, tag any, event component.AppBarEvent) {
	p.Pages.GetWindow().GetLogger().Info().Msgf("%v | Page:About | AppBarEvent | pages:%v tag:%v | event:%v",
		p.Pages.GetWindow().LogPrefix(), pages == p.Pages, tag == p, event)

	switch evt := event.(type) {
	case component.AppBarOverflowActionClicked:
		p.Pages.GetWindow().GetLogger().Info().Msgf("%v | Page:About | AppBarOverflowActionClicked | OverflowActionA:%v OverflowActionB:%v",
			p.Pages.GetWindow().LogPrefix(),
			evt.Tag == &p.Contextual.OverflowActionA,
			evt.Tag == &p.Contextual.OverflowActionB,
		)
	}
}

func (p *Page) OnEventPre(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) OnEventPost(win spsgio.Window, evt event.Event, param any) {

}

func (p *Page) Layout(win spsgio.Window, gtx layout.Context, param any) layout.Dimensions {
	p.List.Axis = layout.Vertical
	theme := win.GetTheme()
	return p.List.Layout(theme, gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "Name", Name, &p.CopyName)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "Author", Author, &p.CopyAuthor)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "License", License, &p.CopyLicense)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.LayoutInfo(theme, gtx, "HomePage", HomePage, &p.CopyHomePage)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if p.Contextual.Clickable.Clicked(gtx) {
					// AppBarActions
					appBarActions := []component.AppBarAction{
						{
							OverflowAction: component.OverflowAction{
								Name: "Favorite",
								Tag:  &p.Contextual.AppBarActionFavorite,
							},
							Layout: func(gtx layout.Context, bg, fg color.NRGBA) layout.Dimensions {
								if p.Contextual.AppBarActionFavorite.Clicked(gtx) {
									p.Pages.GetWindow().GetLogger().Info().Msg("Favorite")
								}
								btn := component.SimpleIconButton(bg, fg, &p.Contextual.AppBarActionFavorite, spsicon.ActionFavorite)
								return btn.Layout(gtx)
							},
						},
					}

					// OverflowActions
					overflowActions := []component.OverflowAction{
						{
							Name: "OverflowAction A",
							Tag:  &p.Contextual.OverflowActionA,
						},
						{
							Name: "OverflowAction B",
							Tag:  &p.Contextual.OverflowActionB,
						},
					}

					p.Pages.GetAppBar().SetContextualActions(appBarActions, overflowActions)
					p.Pages.GetAppBar().ToggleContextual(gtx.Now, "Contextual Title")
				}
				return material.Button(theme, &p.Contextual.Clickable, "Contextual").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.TreeNode.LayoutTreeNodeDefault(theme, gtx, &p.TreeNode)
			}),
		)
	})
}

func (p *Page) LayoutInfo(theme *material.Theme, gtx layout.Context, name, value string, copy *spscopy.Copy) layout.Dimensions {
	return copy.LayoutCopyFlexChild(
		theme,
		gtx,
		func() string {
			return value
		},
		layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:   unit.Dp(8),
				Right:  unit.Dp(8),
				Top:    unit.Dp(4),
				Bottom: unit.Dp(4),
			}.Layout(gtx, material.H6(theme, name).Layout)
		}),
		layout.Flexed(0.8, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:   unit.Dp(8),
				Right:  unit.Dp(8),
				Top:    unit.Dp(4),
				Bottom: unit.Dp(4),
			}.Layout(gtx, material.Body1(theme, value).Layout)
		}),
	)
}
