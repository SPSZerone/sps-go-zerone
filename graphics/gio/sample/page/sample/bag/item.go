package bag

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func newItem(data Data) Item {
	return Item{
		UI:   NewUI(),
		Data: data,
	}
}

type Item struct {
	spsitem.Item
	UI
	Data
}

func (i *Item) Init(item spsitem.Item, theme *material.Theme) {
	i.Item = item
	i.InitMenu(theme)
}

func (i *Item) InitMenu(theme *material.Theme) {
	item := &i.Item
	item.UI.MenuItems = []widget.Clickable{
		{},
	}
	item.UI.Menu = component.MenuState{
		Options: []func(gtx layout.Context) layout.Dimensions{
			func(gtx layout.Context) layout.Dimensions {
				return spssurface.NewSurface().Layout(theme, gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H6(theme, "Item Info").Layout(gtx)
				})
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{}.Layout(theme, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: unit.Dp(16), Right: unit.Dp(16)}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return i.LayoutContentProperty(theme, gtx, "Id", fmt.Sprintf(" %v", i.Id))
					},
				)
			},
			func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: unit.Dp(16), Right: unit.Dp(16)}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return i.LayoutContentProperty(theme, gtx, "Name", fmt.Sprintf(" %v", i.Name))
					},
				)
			},
			func(gtx layout.Context) layout.Dimensions {
				return spsdivider.Divider{Subheading: "Action"}.Layout(theme, gtx)
			},
			func(gtx layout.Context) layout.Dimensions {
				return component.MenuItem(theme, &item.UI.MenuItems[0], "Use").Layout(gtx)
			},
		},
	}
}

func (i *Item) LayoutContent(
	theme *material.Theme, gtx layout.Context,
	item *spsitem.Item, layoutCtx spsitem.LayoutContext,
) layout.Dimensions {
	gtx.Constraints.Max = layoutCtx.Size

	if item.UI.MenuItems[0].Clicked(gtx) {
	}
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, fmt.Sprintf("%v", i.Id)).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, i.Name).Layout(gtx)
		}),
	)
}

func (i *Item) LayoutDetail(
	theme *material.Theme, gtx layout.Context,
	item *spsitem.Item,
) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H5(theme, "Item Detail")
			baseInfo.Font.Style = font.Italic
			baseInfo.Font.Weight = font.Bold
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spssurface.NewSurface().Layout(theme, gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H6(theme, "Use ...").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spsdivider.Divider{}.Layout(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutUseEditor(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutUseSlider(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H5(theme, "Info").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spsdivider.Divider{}.Layout(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, "Id", fmt.Sprintf("%v", i.Id))
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, "Name", fmt.Sprintf("%v", i.Name))
		}),
	)
}

func (i *Item) LayoutUseEditor(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	spacer := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return spsspacer.NewWidthSpacer(8).Layout(gtx)
	})
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		spacer,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseEditor(theme, gtx)
		}),
		spacer,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseEditorBtn(theme, gtx)
		}),
		spacer,
	)
}

func (i *Item) LayoutUseSlider(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	spacer := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return spsspacer.NewWidthSpacer(8).Layout(gtx)
	})
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		spacer,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseSlider(theme, gtx, func() {

			})
		}),
		spacer,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseSliderBtn(theme, gtx)
		}),
		spacer,
	)
}

func (i *Item) LayoutContentProperty(theme *material.Theme, gtx layout.Context, name, value string) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, name).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, value).Layout(gtx)
		}),
	)
}

func (i *Item) LayoutDetailProperty(theme *material.Theme, gtx layout.Context, name, value string) layout.Dimensions {
	return spslayout.FlexInset{
		Ratio: 0.2,
	}.LayoutABWidget(
		gtx,
		material.H6(theme, name).Layout,
		material.Body1(theme, value).Layout,
	)
}
