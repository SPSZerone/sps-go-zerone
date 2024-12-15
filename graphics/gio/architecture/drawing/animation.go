package drawing

import (
	"image"
	"image/color"
	"time"

	"gioui.org/io/input"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func DrawProgressBar(
	ops *op.Ops, source input.Source,
	color color.NRGBA,
	maxWidth float64,
	y int,
	elapsed, duration time.Duration,
) {
	progress := elapsed.Seconds() / duration.Seconds()
	if progress < 1 {
		// The progress bar hasn’t yet finished animating.
		source.Execute(op.InvalidateCmd{})
	} else {
		progress = 1
	}

	width := maxWidth * progress
	defer clip.Rect{Max: image.Pt(int(width), y)}.Push(ops).Pop()
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
