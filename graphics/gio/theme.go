package gio

import (
	"gioui.org/text"
	"gioui.org/widget/material"

	spsfont "github.com/SPSZerone/sps-go-zerone/graphics/gio/font"
	spsnerdfont "github.com/SPSZerone/sps-go-zerone/graphics/gio/font/nerdfont"
)

func NewTheme() *material.Theme {
	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(spsfont.Collection()))
	theme.Face = spsnerdfont.MesloLGSNerdFontMono
	return theme
}
