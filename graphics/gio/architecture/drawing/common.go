package drawing

import (
	"image"

	"gioui.org/f32"
)

func F32Pt2ImagePt(pt f32.Point) image.Point {
	return image.Point{X: int(pt.X), Y: int(pt.Y)}
}

func ImagePt2F32Pt(pt image.Point) f32.Point {
	return f32.Pt(float32(pt.X), float32(pt.Y))
}
