package drawing

import (
	"image"
	"image/color"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func DrawRectSC(ops *op.Ops, size image.Point, color color.NRGBA) {
	drawRectSC1(ops, size, color)
	//drawRectSC2(ops, size, color)
}

func drawRectSC1(ops *op.Ops, size image.Point, color color.NRGBA) {
	defer clip.Rect{Max: size}.Push(ops).Pop()
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func drawRectSC2(ops *op.Ops, size image.Point, color color.NRGBA) {
	stack := clip.Rect{Max: size}.Push(ops)
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
	stack.Pop()
}

func DrawRect(ops *op.Ops, size image.Point, color color.NRGBA, position image.Point) {
	defer op.Offset(position).Push(ops).Pop()
	DrawRectSC(ops, size, color)
}

func DrawRRectR(ops *op.Ops, size image.Point, color color.NRGBA, position image.Point, roundness int) {
	DrawRRect(ops, size, color, position, roundness, roundness, roundness, roundness)
}

func DrawRRect(ops *op.Ops, size image.Point, color color.NRGBA, position image.Point, se, sw, nw, ne int) {
	bounds := image.Rect(position.X, position.Y, position.X+size.X, position.Y+size.Y)
	defer clip.RRect{Rect: bounds, SE: se, SW: sw, NW: nw, NE: ne}.Push(ops).Pop()
	DrawRect(ops, size, color, position)
}

func DrawStrokeRectR(ops *op.Ops, size image.Point, color color.NRGBA, position image.Point, width float32, roundness int) {
	DrawStrokeRect(ops, size, color, position, width, roundness, roundness, roundness, roundness)
}

func DrawStrokeRect(ops *op.Ops, size image.Point, color color.NRGBA, position image.Point, width float32, se, sw, nw, ne int) {
	bounds := image.Rect(position.X, position.Y, position.X+size.X, position.Y+size.Y)
	rRect := clip.RRect{Rect: bounds, SE: se, SW: sw, NW: nw, NE: ne}
	stroke := clip.Stroke{Path: rRect.Path(ops), Width: width}.Op()
	paint.FillShape(ops, color, stroke)
}
