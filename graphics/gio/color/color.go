package color

import (
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func Fill(gtx layout.Context, col1, col2 color.NRGBA) {
	rect := image.Rectangle{Max: gtx.Constraints.Min}
	FillRect(gtx, rect, col1, col2)
}

func FillRectDynamicColor(gtx layout.Context, rect image.Rectangle, col1, col2 int) {
	FillRect(gtx, rect, DynamicColor(col1), DynamicColor(col2))
}

func FillRect(gtx layout.Context, rect image.Rectangle, col1, col2 color.NRGBA) {
	paint.FillShape(
		gtx.Ops,
		color.NRGBA{R: 0, G: 0, B: 0, A: 0xFF},
		clip.Rect(rect).Op(),
	)

	col2.R = byte(float32(col2.R))
	col2.G = byte(float32(col2.G))
	col2.B = byte(float32(col2.B))
	paint.LinearGradientOp{
		Stop1:  f32.Pt(float32(rect.Min.X), 0),
		Stop2:  f32.Pt(float32(rect.Max.X), 0),
		Color1: col1,
		Color2: col2,
	}.Add(gtx.Ops)
	defer clip.Rect(rect).Push(gtx.Ops).Pop()
	paint.PaintOp{}.Add(gtx.Ops)
}

func DynamicColor(i int) color.NRGBA {
	sn, cs := math.Sincos(float64(i) * math.Phi)
	return color.NRGBA{
		R: 0xA0 + byte(0x30*sn),
		G: 0xA0 + byte(0x30*cs),
		B: 0xD0,
		A: 0xFF,
	}
}
