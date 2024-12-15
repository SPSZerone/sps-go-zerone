package drawing

import (
	"image"
	"image/color"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func DrawRecordRect(ops *op.Ops, size image.Point, color color.NRGBA, position image.Point) {
	macroOp := op.Record(ops)
	DrawRectSC(ops, size, color)
	callOp := macroOp.Stop()

	pos := position
	for i := 0; i < 5; i++ {
		stack := op.Offset(pos).Push(ops)
		callOp.Add(ops)
		stack.Pop()
		pos = pos.Add(image.Pt(size.X>>1, size.Y>>3))
	}
}

func DrawRecordCache(ops *op.Ops) {
	cache := new(op.Ops)
	macroOp := op.Record(cache)

	cl := clip.Rect{Max: image.Pt(100, 100)}.Push(cache)
	paint.ColorOp{Color: color.NRGBA{G: 0x80, A: 0xFF}}.Add(cache)
	paint.PaintOp{}.Add(cache)
	cl.Pop()

	callOp := macroOp.Stop()

	callOp.Add(ops)
}
