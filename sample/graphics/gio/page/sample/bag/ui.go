package bag

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spseditor "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/editor"
	spsmenu "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/menu"
	spsslider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/slider"
)

func NewUI() UI {
	return UI{
		UseEditor: spseditor.NewEditor(spseditor.OptHint("use count")),
		UseSlider: spsslider.NewSlider(),
	}
}

type UI struct {
	MenuDetail []spsmenu.Menu

	UseEditor    spseditor.Editor
	UseEditorBtn widget.Clickable

	UseSlider    spsslider.Slider
	UseSliderBtn widget.Clickable
}

func (u *UI) GetMenuDetail(index int) *spsmenu.Menu {
	if index < 0 || index >= len(u.MenuDetail) {
		u.MenuDetail = append(u.MenuDetail, spsmenu.New())
		return &u.MenuDetail[len(u.MenuDetail)-1]
	}
	return &u.MenuDetail[index]
}

func (u *UI) LayoutUseEditor(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return u.UseEditor.Layout(
		theme,
		gtx,
		func(e *spseditor.Editor, s *material.EditorStyle) {
			s.Color = color.NRGBA{A: 0xFF, R: 0xFF, G: 0x00, B: 0x00}
		},
		func(e *spseditor.Editor) {
		},
	)
}

func (u *UI) LayoutUseEditorBtn(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return material.Button(theme, &u.UseEditorBtn, "Use").Layout(gtx)
}

func (u *UI) LayoutUseSlider(theme *material.Theme, gtx layout.Context, style spsslider.Style, onValueChanged spsslider.OnValueChanged) layout.Dimensions {
	return u.UseSlider.LayoutSlider(theme, gtx, style, onValueChanged)
}

func (u *UI) LayoutUseSliderBtn(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return material.Button(theme, &u.UseSliderBtn, "Use").Layout(gtx)
}
