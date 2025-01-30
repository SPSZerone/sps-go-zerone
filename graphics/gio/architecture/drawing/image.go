package drawing

import (
	"image"
	"image/color"
	"math/rand"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/paint"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

// DrawImage
//
// paint.ImageOp is used to draw images.
// Like paint.ColorOp, it sets part of the drawing context (the “brush”) that’s used for subsequent paint.PaintOp.
// paint.ImageOp is used similarly to ColorOp.
//
// Note that image.NRGBA and image.Uniform images are efficient and treated specially.
// Other Image implementations will undergo a more expensive copy and conversion to the underlying image model.
func DrawImage(
	ops *op.Ops,
	img image.Image, filter paint.ImageFilter,
	scale f32.Point,
	position image.Point,
) {
	defer op.Offset(position).Push(ops).Pop()
	drawImage(ops, img, filter, scale)
}

func drawImage(
	ops *op.Ops,
	img image.Image, filter paint.ImageFilter,
	scale f32.Point,
) {
	imageOp := paint.NewImageOp(img)
	imageOp.Filter = filter
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), scale)).Add(ops)
	paint.PaintOp{}.Add(ops)
}

func NewImageNRGBARandColor(bounds image.Rectangle) *image.NRGBA {
	img := image.NewNRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			img.SetNRGBA(x, y, color.NRGBA{
				R: uint8(rand.Intn(0xFF)),
				G: uint8(rand.Intn(0xFF)),
				B: uint8(rand.Intn(0xFF)),
				A: 0xFF})
		}
	}
	return img
}

func NewImageNRGBADynamicColor(bounds image.Rectangle, xColor, yColor int) *image.NRGBA {
	img := image.NewNRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			img.SetNRGBA(x, y, spscolor.DynamicColor(x/xColor+y/yColor))
		}
	}
	return img
}
