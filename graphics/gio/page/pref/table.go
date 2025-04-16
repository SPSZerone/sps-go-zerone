package pref

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func tableDimension(app *spsgio.Application, gtx layout.Context, axis layout.Axis, index, constraint, minSize, height int) int {
	switch axis {
	case layout.Horizontal:
		var widthUnit int
		if app.Pref.Settings.ValueInFront.Value {
			switch index {
			case TableColIdxKey:
				widthUnit = gtx.Dp(unit.Dp(TableColWidthValue))
			case TableColIdxValue:
				widthUnit = gtx.Dp(unit.Dp(TableColWidthKey))
			}
		} else {
			switch index {
			case TableColIdxKey:
				widthUnit = gtx.Dp(unit.Dp(TableColWidthKey))
			case TableColIdxValue:
				widthUnit = gtx.Dp(unit.Dp(TableColWidthValue))
			}
		}
		return widthUnit
	default:
		return height
	}
}

func settingsCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	colIdx := TableColIdxValue
	if app.Pref.Settings.ValueInFront.Value {
		colIdx = TableColIdxKey
	}
	switch row {
	case SettingsRowIdxNonModalDrawer:
		switch col {
		case colIdx:
			return p.NonModalDrawer(app, gtx)
		default:
			labelStyle.Text = app.Pref.Settings.ModalNavDrawer.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxTabAxis:
		switch col {
		case colIdx:
			return p.TabAxis(app, gtx)
		default:
			labelStyle.Text = app.Pref.Settings.TabAxis.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxValueInFront:
		switch col {
		case colIdx:
			return p.ValueInFront(app, gtx)
		default:
			labelStyle.Text = app.Pref.Settings.ValueInFront.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxBottomBar:
		switch col {
		case colIdx:
			return p.BottomBar(app, gtx)
		default:
			labelStyle.Text = app.Pref.Settings.BottomBar.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxDecorated:
		switch col {
		case colIdx:
			return p.Decorated(app, gtx)
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

func prefCell(p *Page, app *spsgio.Application, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	colIdx := TableColIdxValue
	if app.Pref.Settings.ValueInFront.Value {
		colIdx = TableColIdxKey
	}
	switch row {
	case PreferencesRowIdxTableStyle:
		switch col {
		case colIdx:
			return app.Pref.Pref.TableStyle.LayoutSwitch(app.Window, app.Theme, gtx)
		default:
			labelStyle.Text = app.Pref.Pref.TableStyle.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}
