package drawing

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func DrawTriangle(
	ops *op.Ops,
	color color.NRGBA,
	point1 f32.Point, point2 f32.Point, point3 f32.Point,
) {
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(point1)
	path.LineTo(point2)
	path.LineTo(point3)
	path.Close()

	shape := clip.Outline{Path: path.End()}.Op()
	paint.FillShape(ops, color, shape)
}

func DrawStrokeTriangle(
	ops *op.Ops,
	color color.NRGBA,
	point1 f32.Point, point2 f32.Point, point3 f32.Point,
	width float32,
) {
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(point1)
	path.LineTo(point2)
	path.LineTo(point3)
	path.Close()

	stroke := clip.Stroke{Path: path.End(), Width: width}.Op()
	paint.FillShape(ops, color, stroke)
}
