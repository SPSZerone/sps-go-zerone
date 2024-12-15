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
func DrawBezierQuadratic(ops *op.Ops, opts BezierOptions) {
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(opts.Move.Add(opts.Pos))                           // P0
	path.QuadTo(opts.Ctrl[0].Add(opts.Pos), opts.To.Add(opts.Pos)) // P1, P2
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	DrawRect(ops, opts.Size, opts.Color, F32Pt2ImagePt(opts.Pos))
}

// DrawBezierCubic
//
// see [WIKI-EN] [WIKI-CN]
//
// [WIKI-EN]: https://en.wikipedia.org/wiki/B%C3%A9zier_curve#Cubic_B%C3%A9zier_curves
// [WIKI-CN]: https://zh.wikipedia.org/zh-cn/%E8%B2%9D%E8%8C%B2%E6%9B%B2%E7%B7%9A#%E9%AB%98%E9%9A%8E%E6%9B%B2%E7%B7%9A
func DrawBezierCubic(ops *op.Ops, opts BezierOptions) {
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(opts.Move.Add(opts.Pos))                                                       // P0
	path.CubeTo(opts.Ctrl[0].Add(opts.Pos), opts.Ctrl[1].Add(opts.Pos), opts.To.Add(opts.Pos)) // P1, P2, P3
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	DrawRect(ops, opts.Size, opts.Color, F32Pt2ImagePt(opts.Pos))
}

type BezierOption func(*BezierOptions)

type BezierOptions struct {
	Size  image.Point
	Color color.NRGBA
	Pos   f32.Point
	Move  f32.Point
	To    f32.Point
	Ctrl  []f32.Point
}

func BZOptSize(value image.Point) BezierOption {
	return func(opts *BezierOptions) {
		opts.Size = value
	}
}

func BZOptColor(value color.NRGBA) BezierOption {
	return func(opts *BezierOptions) {
		opts.Color = value
	}
}

func BZOptPos(value f32.Point) BezierOption {
	return func(opts *BezierOptions) {
		opts.Pos = value
	}
}

func BZOptBPoints(move, to f32.Point, ctrl ...f32.Point) BezierOption {
	return func(opts *BezierOptions) {
		opts.Move = move
		opts.To = to
		opts.Ctrl = ctrl
	}
}

func BZOptMove(value f32.Point) BezierOption {
	return func(opts *BezierOptions) {
		opts.Move = value
	}
}

func BZOptTo(value f32.Point) BezierOption {
	return func(opts *BezierOptions) {
		opts.To = value
	}
}

func BZOptCtrl(value ...f32.Point) BezierOption {
	return func(opts *BezierOptions) {
		opts.Ctrl = value
	}
}
