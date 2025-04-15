package pref

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func decoratedCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	colIdx := TableColIdxValue
	if app.Pref.Settings.ValueInFront {
		colIdx = TableColIdxKey
	}
	switch row {
	case DecoratedRowIdxDecorated:
		switch col {
		case colIdx:
			return app.Pref.Settings.DecoratedLayoutSwitch(app.Window, app.Theme, gtx)
		default:
			labelStyle.Text = app.Pref.Settings.Decorated.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}
