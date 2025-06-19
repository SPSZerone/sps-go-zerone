package editor

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	spsbg "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bg"
	spsclickable "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/clickable"
	spslist "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/list"
)

func NewSuggestEditor(opts ...Option) SuggestEditor {
	return SuggestEditor{
		Editor: New(opts...),
		List:   spslist.New(),
	}
}

type (
	SuggestWidget func(gtx layout.Context, index int, size image.Point, onClick func(content string)) layout.Dimensions
)

type SuggestEditor struct {
	Editor Editor
	List   spslist.List
}

func (e *SuggestEditor) LayoutDefault(
	theme *material.Theme, gtx layout.Context,
	listHeight int,
	height int,
	length int,
	widget SuggestWidget,
) layout.Dimensions {
	return e.Layout(
		theme, gtx,
		nil, nil,
		listHeight, height,
		length, widget,
	)
}

func (e *SuggestEditor) Layout(
	theme *material.Theme, gtx layout.Context,
	style Style, onSubmit OnSubmit,
	listHeight int,
	height int,
	length int,
	widget SuggestWidget,
) layout.Dimensions {
	var width int
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					dim := e.Editor.Layout(theme, gtx, style, onSubmit)
					width = dim.Size.X
					return dim
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.Y = listHeight
			gtx.Constraints.Max.X = width
			size := image.Pt(width, height)
			return e.List.Layout(theme, gtx, length, func(gtx layout.Context, index int) layout.Dimensions {
				return widget(gtx, index, size, func(content string) {
					e.Editor.SetText(content)
				})
			})
		}),
	)
}

func NewSuggest(content string, display layout.Widget, opts ...SuggestOption) Suggest {
	s := Suggest{
		Clickable: spsclickable.New(),
		Widget:    display,
		Content:   content,
	}
	s.BGColor1, s.BGColor2 = spscolor.Rand2Color(0, 10)
	s.Update(opts...)
	return s
}

type Suggest struct {
	BGColor1 color.NRGBA
	BGColor2 color.NRGBA

	Clickable spsclickable.Clickable

	Widget  layout.Widget
	Content string
}

func (s *Suggest) Update(opts ...SuggestOption) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *Suggest) Layout(
	theme *material.Theme, gtx layout.Context,
	size image.Point,
	onClick func(),
) layout.Dimensions {
	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max = size
			return spsbg.NewColorful(s.BGColor1, s.BGColor2).LayoutBG(theme, gtx, size)
		}),
		// click area
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if s.Clickable.Clicked(gtx) {
				onClick()
			}
			return s.Clickable.Layout(
				gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{
						Size: size,
					}
				},
			)
		}),
		// content
		layout.Stacked(s.Widget),
	)
}
