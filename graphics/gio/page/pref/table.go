package pref

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

func tableDimension(win spsgio.Window, gtx layout.Context, axis layout.Axis, index, constraint, minSize, height int) int {
	pref := win.GetPref()

	switch axis {
	case layout.Horizontal:
		var widthUnit int
		if pref.Settings.ValueInFront.Value {
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

func settingsCell(p *Page, win spsgio.Window, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	pref := win.GetPref()

	colIdx := TableColIdxValue
	if pref.Settings.ValueInFront.Value {
		colIdx = TableColIdxKey
	}
	switch row {
	case SettingsRowIdxNonModalDrawer:
		switch col {
		case colIdx:
			return p.NonModalDrawer(win, gtx)
		default:
			labelStyle.Text = pref.Settings.ModalNavDrawer.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxTabAxis:
		switch col {
		case colIdx:
			return p.TabAxis(win, gtx)
		default:
			labelStyle.Text = pref.Settings.TabAxis.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxValueInFront:
		switch col {
		case colIdx:
			return p.ValueInFront(win, gtx)
		default:
			labelStyle.Text = pref.Settings.ValueInFront.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxBottomBar:
		switch col {
		case colIdx:
			return p.BottomBar(win, gtx)
		default:
			labelStyle.Text = pref.Settings.BottomBar.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxDecorated:
		switch col {
		case colIdx:
			return p.Decorated(win, gtx)
		default:
			labelStyle.Text = pref.Settings.Decorated.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}

func prefCell(p *Page, win spsgio.Window, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	window := win.GetWindow()
	theme := win.GetTheme()
	pref := win.GetPref()

	colIdx := TableColIdxValue
	if pref.Settings.ValueInFront.Value {
		colIdx = TableColIdxKey
	}
	switch row {
	case PreferencesRowIdxTableStyle:
		switch col {
		case colIdx:
			return pref.Pref.TableStyle.LayoutSwitch(window, theme, gtx)
		default:
			labelStyle.Text = pref.Pref.TableStyle.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}
