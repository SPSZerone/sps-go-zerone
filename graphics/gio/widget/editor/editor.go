package editor

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func NewEditor(opts ...Option) Editor {
	return Editor{
		Editor: widget.Editor{
			SingleLine: true,
			Submit:     true,
		},
		Border: widget.Border{
			Color:        color.NRGBA{A: 0xFF, R: 0xFF, G: 0xFF, B: 0xFF},
			CornerRadius: unit.Dp(8),
			Width:        unit.Dp(2),
		},
	}
}

type OnSubmit func(e *Editor)
type Style func(e *Editor, s *material.EditorStyle)

type Editor struct {
	widget.Editor
	widget.Border

	Hint string
}

func (e *Editor) Update(opts ...Option) {
	for _, o := range opts {
		o(e)
	}
}

func (e *Editor) Layout(
	theme *material.Theme, gtx layout.Context,
	style Style, onSubmit OnSubmit,
) layout.Dimensions {
	e.LayoutEvent(gtx, onSubmit)
	return e.LayoutEditor(theme, gtx, style)
}

func (e *Editor) LayoutEvent(gtx layout.Context, onSubmit OnSubmit) {
	for {
		editorEvent, ok := e.Editor.Update(gtx)
		if !ok {
			break
		}
		if _, ok := editorEvent.(widget.SubmitEvent); ok {
			if onSubmit != nil {
				onSubmit(e)
			}
		}
	}
}

func (e *Editor) LayoutEditor(theme *material.Theme, gtx layout.Context, style Style) layout.Dimensions {
	editorStyle := material.Editor(theme, &e.Editor, e.Hint)
	editorStyle.Font.Style = font.Italic
	editorStyle.Font.Weight = font.Bold
	if style != nil {
		style(e, &editorStyle)
	}
	return e.Border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(8)).Layout(gtx, editorStyle.Layout)
	})
}
