package tab

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsbg "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bg"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spsmenu "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/menu"
)

func NewUI() UI {
	ui := UI{
		FlexInset: spslayout.New(),
		Menu:      spsmenu.New(),
	}
	ui.FlexInset.Flex.Axis = layout.Horizontal
	ui.BGColor1, ui.BGColor2 = spscolor.Rand2Color(0, 10)
	//ui.BGColor2 = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	return ui
}

type UI struct {
	FlexInset spslayout.FlexInset

	Clickable      widget.Clickable
	CloseClickable widget.Clickable
	Menu           spsmenu.Menu

	BGColor1, BGColor2 color.NRGBA
	size               image.Point
}

func (u *UI) LayoutContent(
	theme *material.Theme, gtx layout.Context,
	isVertical bool,
	widthLimit int,
	highlightThickness int,
	colorfulBG bool,
	closeMode CloseMode,
	nameWidget layout.Widget,
) layout.Dimensions {
	if widthLimit > 0 {
		if isVertical {
			gtx.Constraints.Max.X = widthLimit
		}
	}

	if !colorfulBG {
		return u.doLayoutContent(theme, gtx, isVertical, widthLimit, closeMode, nameWidget)
	}

	return layout.Stack{
		Alignment: layout.Center,
	}.Layout(
		gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max = u.size
			return spsbg.NewColorful(u.BGColor1, u.BGColor2).LayoutBG(theme, gtx, u.size)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			dim := u.doLayoutContent(theme, gtx, isVertical, widthLimit, closeMode, nameWidget)
			u.size = dim.Size
			return dim
		}),
	)
}

func (u *UI) doLayoutContent(
	theme *material.Theme, gtx layout.Context,
	isVertical bool,
	widthLimit int,
	closeMode CloseMode,
	nameWidget layout.Widget,
) layout.Dimensions {
	// final name widget
	var finalNameWidget layout.Widget
	if closeMode == CloseModeNormal {
		finalNameWidget = func(gtx layout.Context) layout.Dimensions {
			var nameChild layout.FlexChild
			if isVertical {
				nameChild = layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return u.Clickable.Layout(gtx, nameWidget)
				})
			} else {
				nameChild = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return u.Clickable.Layout(gtx, nameWidget)
				})
			}
			return layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}.Layout(gtx,
				nameChild,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return u.LayoutClose(theme, gtx, &u.CloseClickable)
				}),
			)
		}
	} else {
		finalNameWidget = func(gtx layout.Context) layout.Dimensions {
			dims := nameWidget(gtx)
			size := dims.Size
			if widthLimit > 0 {
				if isVertical {
					size.X = widthLimit
				}
			}
			return layout.Dimensions{Size: size}
		}
	}

	// menu widgets
	menuWidgets := []func(gtx layout.Context) layout.Dimensions{
		func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H6(theme, fmt.Sprintf("Tab Menu")).Layout(gtx)
			})
		},
		func(gtx layout.Context) layout.Dimensions {
			return spsdivider.Divider{
				Subheading: "Actions",
			}.Layout(theme, gtx)
		},
	}
	if closeMode != CloseModeNone {
		menuWidgets = append(menuWidgets,
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return u.LayoutClose(theme, gtx, &u.CloseClickable)
				}))
			})
	}
	u.Menu.SetWidgets(menuWidgets...)

	// final
	return layout.Stack{
		Alignment: layout.Center,
	}.Layout(
		gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return u.Clickable.Layout(gtx, finalNameWidget)
		}),
		// menu
		u.Menu.LayoutExpandedContextArea(theme, image.Point{}),
	)
}

func (u *UI) LayoutHighlight(
	theme *material.Theme, gtx layout.Context,
	highlightThickness int, highlight bool,
	isVertical bool,
	size image.Point,
) layout.Dimensions {
	if !highlight {
		return layout.Dimensions{}
	}

	var rect image.Rectangle
	if isVertical {
		// on right
		startX := size.X - highlightThickness
		rect = image.Rect(startX, 0, startX+highlightThickness, size.Y)
	} else {
		// on bottom
		rect = image.Rect(0, size.Y-highlightThickness, size.X, size.Y)
	}

	paint.FillShape(gtx.Ops, theme.Palette.ContrastBg, clip.Rect(rect).Op())
	return layout.Dimensions{Size: size}
}

func (u *UI) LayoutClose(
	theme *material.Theme, gtx layout.Context,
	close *widget.Clickable,
) layout.Dimensions {
	return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//   󰅖    󰅗 󰅙 󰅜 󰛉 󱎘 󰖭  󰅘 󰅚 󰅝
		return material.Button(theme, close, ``).Layout(gtx)
		//return material.IconButton(theme, close, spsicon.NavigationClose, "Close").Layout(gtx)
	})
}
