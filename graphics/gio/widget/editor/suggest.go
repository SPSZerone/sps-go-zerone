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

type SuggestEditor struct {
	Editor Editor
	List   spslist.List
}

func (e *SuggestEditor) LayoutSimple(
	theme *material.Theme, gtx layout.Context,
	suggestListHeight int,
	suggestHeight int,
	suggests []*Suggest,
) layout.Dimensions {
	return e.Layout(
		theme, gtx,
		nil, nil,
		suggestListHeight, suggestHeight,
		suggests,
		nil,
	)
}

func (e *SuggestEditor) Layout(
	theme *material.Theme, gtx layout.Context,
	style Style, onSubmit OnSubmit,
	suggestListHeight int,
	suggestHeight int,
	suggests []*Suggest,
	onSuggestClick func(suggest *Suggest),
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
			gtx.Constraints.Max.Y = suggestListHeight
			gtx.Constraints.Max.X = width
			size := image.Pt(width, suggestHeight)
			return e.List.Layout(theme, gtx, len(suggests), func(gtx layout.Context, index int) layout.Dimensions {
				suggest := suggests[index]
				return suggest.Layout(theme, gtx, size, func(suggest *Suggest) {
					if onSuggestClick != nil {
						onSuggestClick(suggest)
					}
					e.Editor.SetText(suggest.Content)
				})
			})
		}),
	)
}

func NewSuggest(content string, display layout.Widget, opts ...SuggestOption) *Suggest {
	s := &Suggest{
		Clickable: spsclickable.New(),
		Display:   display,
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

	Display layout.Widget
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
	onClick func(suggest *Suggest),
) layout.Dimensions {
	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// content background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max = size
			return spsbg.NewColorful(s.BGColor1, s.BGColor2).LayoutBG(theme, gtx, size)
		}),
		// click area
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			if s.Clickable.Clicked(gtx) {
				onClick(s)
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
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return s.Display(gtx)
		}),
	)
}
