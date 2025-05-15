package table

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"
)

func New(opts ...Option) Table {
	t := Table{
		MinSize: 200,
		Border: widget.Border{
			Color: color.NRGBA{A: 255},
			Width: unit.Dp(1),
		},
		Inset: layout.UniformInset(unit.Dp(2)),
	}
	t.Update(opts...)
	return t
}

type (
	LabelStyle  func(theme *material.Theme) material.LabelStyle
	Dimensioner func(axis layout.Axis, index, constraint, minSize, height int) int
	Cell        func(gtx layout.Context, row, col int, labelStyle material.LabelStyle) layout.Dimensions
)

type Table struct {
	Headers []Header

	MinSize float32
	Border  widget.Border
	Inset   layout.Inset

	HeaderLabelStyle LabelStyle
	DataLabelStyle   LabelStyle

	gridState component.GridState
}

func (t *Table) Update(opts ...Option) {
	for _, opt := range opts {
		opt(t)
	}
}

func (t *Table) Layout(
	theme *material.Theme, gtx layout.Context,
	rows int, dimensioner Dimensioner, cell Cell,
) layout.Dimensions {
	minSize := gtx.Dp(unit.Dp(t.MinSize))
	border := t.Border
	inset := t.Inset

	var headingLabel material.LabelStyle
	if t.HeaderLabelStyle != nil {
		headingLabel = t.HeaderLabelStyle(theme)
	} else {
		headingLabel = DefaultHeaderLabelStyle(theme)
	}

	var dataLabel material.LabelStyle
	if t.DataLabelStyle != nil {
		dataLabel = t.DataLabelStyle(theme)
	} else {
		dataLabel = DefaultDataLabelStyle(theme)
	}

	orig := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	macro := op.Record(gtx.Ops)
	dims := inset.Layout(gtx, headingLabel.Layout)
	_ = macro.Stop()
	gtx.Constraints = orig

	var finalDimensioner outlay.Dimensioner
	if dimensioner != nil {
		finalDimensioner = func(axis layout.Axis, index, constraint int) int {
			return dimensioner(axis, index, constraint, minSize, dims.Size.Y)
		}
	} else {
		finalDimensioner = func(axis layout.Axis, index, constraint int) int {
			widthUnit := max(int(float32(constraint)/3), minSize)
			switch axis {
			case layout.Horizontal:
				return widthUnit
			default:
				return dims.Size.Y
			}
		}
	}

	finalHeading := func(gtx layout.Context, col int) layout.Dimensions {
		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				headingLabel.Text = t.Headers[col].Text
				return headingLabel.Layout(gtx)
			})
		})
	}

	finalCell := func(gtx layout.Context, row, col int) layout.Dimensions {
		return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return cell(gtx, row, col, dataLabel)
		})
	}

	return component.Table(theme, &t.gridState).Layout(
		gtx,
		rows, len(t.Headers),
		finalDimensioner, finalHeading, finalCell,
	)
}

func DefaultHeaderLabelStyle(th *material.Theme) material.LabelStyle {
	result := material.Body1(th, "")
	result.Font.Weight = font.Bold
	result.Alignment = text.Middle
	result.MaxLines = 1
	return result
}

func DefaultDataLabelStyle(th *material.Theme) material.LabelStyle {
	result := material.Body1(th, "")
	result.Font.Typeface = "Go Mono"
	result.MaxLines = 1
	result.Alignment = text.End
	return result
}
