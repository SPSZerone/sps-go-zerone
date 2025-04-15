package pref

import (
	"gioui.org/layout"
	"gioui.org/unit"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
)

const (
	TabIdxSettings = iota
	TabIdxDecorated
)

const (
	TabNameSettings  = "Settings"
	TabNameDecorated = "Decorated"
)

const (
	RatioKey   = 0.3
	RatioValue = 0.2
)

const (
	TableColIdxKey = iota
	TableColIdxValue
)

const (
	TableColWidthKey   = 220
	TableColWidthValue = 64
)

const (
	SettingsRowIdxNonModalDrawer = iota
	SettingsRowIdxValueInFront
	SettingsRowIdxBottomBar
	SettingsRowIdxTableStyle
	SettingsRowCount
)

const (
	DecoratedRowIdxDecorated = iota
	DecoratedRowCount
)

func dimension(app *spsgio.Application, gtx layout.Context, axis layout.Axis, index, constraint, minSize, height int) int {
	switch axis {
	case layout.Horizontal:
		var widthUnit int
		if app.Pref.Pref.ValueInFront {
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
