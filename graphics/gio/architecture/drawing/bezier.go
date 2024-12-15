package drawing

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
)

// DrawBezierQuadratic
//
// see [WIKI-EN] [WIKI-CN]
//
// [WIKI-EN]: https://en.wikipedia.org/wiki/B%C3%A9zier_curve#Quadratic_B%C3%A9zier_curves
// [WIKI-CN]: https://zh.wikipedia.org/zh-cn/%E8%B2%9D%E8%8C%B2%E6%9B%B2%E7%B7%9A#%E4%BA%8C%E6%AC%A1%E6%9B%B2%E7%B7%9A
func DrawBezierQuadratic(
	ops *op.Ops,
	size image.Point, color color.NRGBA, position image.Point,
	move, ctrl, to f32.Point,
) {
	pos := f32.Pt(float32(position.X), float32(position.Y))
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(move.Add(pos))              // P0
	path.QuadTo(ctrl.Add(pos), to.Add(pos)) // P1, P2
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	DrawRect(ops, size, color, position)
}

// DrawBezierCubic
//
// see [WIKI-EN] [WIKI-CN]
//
// [WIKI-EN]: https://en.wikipedia.org/wiki/B%C3%A9zier_curve#Cubic_B%C3%A9zier_curves
// [WIKI-CN]: https://zh.wikipedia.org/zh-cn/%E8%B2%9D%E8%8C%B2%E6%9B%B2%E7%B7%9A#%E9%AB%98%E9%9A%8E%E6%9B%B2%E7%B7%9A
func DrawBezierCubic(
	ops *op.Ops,
	size image.Point, color color.NRGBA, position image.Point,
	move, ctrl1, ctrl2, to f32.Point,
) {
	pos := f32.Pt(float32(position.X), float32(position.Y))
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(move.Add(pos))                               // P0
	path.CubeTo(ctrl1.Add(pos), ctrl2.Add(pos), to.Add(pos)) // P1, P2, P3
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	DrawRect(ops, size, color, position)
}
