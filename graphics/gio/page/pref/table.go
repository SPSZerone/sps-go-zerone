package pref

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spswin "github.com/SPSZerone/sps-go-zerone/graphics/gio/window"
)

func tableDimension(win *spswin.Window, gtx layout.Context, axis layout.Axis, index, constraint, minSize, height int) int {
	switch axis {
	case layout.Horizontal:
		var widthUnit int
		if win.Pref.Settings.ValueInFront.Value {
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

func settingsCell(p *Page, win *spswin.Window, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	colIdx := TableColIdxValue
	if win.Pref.Settings.ValueInFront.Value {
		colIdx = TableColIdxKey
	}
	switch row {
	case SettingsRowIdxNonModalDrawer:
		switch col {
		case colIdx:
			return p.NonModalDrawer(win, gtx)
		default:
			labelStyle.Text = win.Pref.Settings.ModalNavDrawer.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxTabAxis:
		switch col {
		case colIdx:
			return p.TabAxis(win, gtx)
		default:
			labelStyle.Text = win.Pref.Settings.TabAxis.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxValueInFront:
		switch col {
		case colIdx:
			return p.ValueInFront(win, gtx)
		default:
			labelStyle.Text = win.Pref.Settings.ValueInFront.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxBottomBar:
		switch col {
		case colIdx:
			return p.BottomBar(win, gtx)
		default:
			labelStyle.Text = win.Pref.Settings.BottomBar.Name
		}
		return labelStyle.Layout(gtx)
	case SettingsRowIdxDecorated:
		switch col {
		case colIdx:
			return p.Decorated(win, gtx)
		default:
			labelStyle.Text = win.Pref.Settings.Decorated.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}

func prefCell(p *Page, win *spswin.Window, gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions {
	colIdx := TableColIdxValue
	if win.Pref.Settings.ValueInFront.Value {
		colIdx = TableColIdxKey
	}
	switch row {
	case PreferencesRowIdxTableStyle:
		switch col {
		case colIdx:
			return win.Pref.Pref.TableStyle.LayoutSwitch(win.Window, win.Theme, gtx)
		default:
			labelStyle.Text = win.Pref.Pref.TableStyle.Name
		}
		return labelStyle.Layout(gtx)
	default:
		labelStyle.Text = "Unknown"
		labelStyle.Color = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
		return labelStyle.Layout(gtx)
	}
}
