package image

import (
	"image"
	"image/draw"
	"math"
)

// FlipMode defines the direction for the flip operation.
type FlipMode int

const (
	FlipNone FlipMode = iota
	FlipVertically
	FlipHorizontally
	FlipBoth // Rotates the image 180 degrees
)

// Processor wraps an image.RGBA to provide chainable processing methods.
// It manipulates the Pix slice directly for high performance.
type Processor struct {
	img *image.RGBA
}

//	Init from any image.Image
//
// It converts the source image to RGBA format if necessary.
func (p *Processor) Init(src image.Image, deepCopyIfRGBA bool) *Processor {
	if rgba, ok := src.(*image.RGBA); ok {
		// If it's already an RGBA, make a deep copy to avoid modifying the original source.
		if deepCopyIfRGBA {
			dst := image.NewRGBA(rgba.Bounds())
			copy(dst.Pix, rgba.Pix)
			p.img = dst
			return p
		}

		p.img = rgba
		return p
	}

	// Convert other formats (e.g., JPEG YCbCr) to RGBA.
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	p.img = dst
	return p
}

func (p *Processor) GetImage() *image.RGBA {
	return p.img
}

// ---------------------------------------------------------------------
// Basic Color Adjustments
// ---------------------------------------------------------------------

// ToGray converts the image to grayscale using the weighted average method.
// Formula: 0.299R + 0.587G + 0.114B
func (p *Processor) ToGray() *Processor {
	bounds := p.img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			offset := p.img.PixOffset(x, y)

			r := float64(p.img.Pix[offset])
			g := float64(p.img.Pix[offset+1])
			b := float64(p.img.Pix[offset+2])

			gray := uint8(0.299*r + 0.587*g + 0.114*b)

			p.img.Pix[offset] = gray   // R
			p.img.Pix[offset+1] = gray // G
			p.img.Pix[offset+2] = gray // B
			// Alpha (offset+3) remains unchanged
		}
	}
	return p
}

// Invert inverts the colors (negative effect).
// Formula: 255 - current_value
func (p *Processor) Invert() *Processor {
	for i := 0; i < len(p.img.Pix); i += 4 {
		p.img.Pix[i] = 255 - p.img.Pix[i]     // R
		p.img.Pix[i+1] = 255 - p.img.Pix[i+1] // G
		p.img.Pix[i+2] = 255 - p.img.Pix[i+2] // B
	}
	return p
}

// AdjustBrightness adjusts the brightness by a delta value.
// delta range: typically -255 to 255.
func (p *Processor) AdjustBrightness(delta int) *Processor {
	// Pre-calculate lookup table for performance
	p.Adjust(func(lookup []uint8) {
		for i := 0; i < len(lookup); i++ {
			lookup[i] = clamp(i + delta)
		}
	})
	return p
}

// AdjustContrast adjusts the contrast.
// factor: 1.0 = original, >1.0 = increase contrast, <1.0 = decrease contrast.
func (p *Processor) AdjustContrast(factor float64) *Processor {
	p.Adjust(func(lookup []uint8) {
		for i := 0; i < len(lookup); i++ {
			val := ((float64(i) - 128.0) * factor) + 128.0
			lookup[i] = clamp(int(val))
		}
	})
	return p
}

func (p *Processor) Adjust(lookupValue func(lookup []uint8)) *Processor {
	lookup := getLookup()
	lookupValue(lookup)
	p.applyLookup(lookup)
	return p
}

func (p *Processor) applyLookup(lookup []uint8) *Processor {
	for i := 0; i < len(p.img.Pix); i += 4 {
		p.img.Pix[i] = lookup[p.img.Pix[i]]
		p.img.Pix[i+1] = lookup[p.img.Pix[i+1]]
		p.img.Pix[i+2] = lookup[p.img.Pix[i+2]]
	}
	return p
}

func getLookup() []uint8 {
	return make([]uint8, 256)
}

// ---------------------------------------------------------------------
// Transformations (Resize & Flip)
// ---------------------------------------------------------------------

// Resize resizes the image using Nearest Neighbor interpolation.
func (p *Processor) Resize(newWidth, newHeight int) *Processor {
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xRatio := float64(p.img.Bounds().Dx()) / float64(newWidth)
	yRatio := float64(p.img.Bounds().Dy()) / float64(newHeight)

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Find corresponding coordinate in source image
			srcX := int(math.Floor(float64(x) * xRatio))
			srcY := int(math.Floor(float64(y) * yRatio))

			srcOffset := p.img.PixOffset(srcX, srcY)
			dstOffset := dst.PixOffset(x, y)

			copy(dst.Pix[dstOffset:dstOffset+4], p.img.Pix[srcOffset:srcOffset+4])
		}
	}

	p.img = dst
	return p
}

// Flip performs a flip based on the provided mode.
func (p *Processor) Flip(mode FlipMode) *Processor {
	switch mode {
	case FlipVertically:
		return p.FlipVertically()
	case FlipHorizontally:
		return p.FlipHorizontally()
	case FlipBoth:
		return p.FlipVertically().FlipHorizontally()
	}
	return p
}

// FlipVertically flips the image upside down.
// Optimized by swapping entire rows of bytes.
func (p *Processor) FlipVertically() *Processor {
	bounds := p.img.Bounds()
	h := bounds.Dy()
	stride := p.img.Stride

	rowBuf := make([]uint8, stride)
	// Iterate only half the height to swap top rows with bottom rows
	for y := 0; y < h/2; y++ {
		topOffset := y * stride
		bottomOffset := (h - 1 - y) * stride

		// Swap rows
		copy(rowBuf, p.img.Pix[topOffset:topOffset+stride])
		copy(p.img.Pix[topOffset:topOffset+stride], p.img.Pix[bottomOffset:bottomOffset+stride])
		copy(p.img.Pix[bottomOffset:bottomOffset+stride], rowBuf)
	}
	return p
}

// FlipHorizontally flips the image left to right.
// Optimized by swapping pixels within the same row.
func (p *Processor) FlipHorizontally() *Processor {
	bounds := p.img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	stride := p.img.Stride

	for y := 0; y < h; y++ {
		rowStart := y * stride
		// Iterate only half the width
		for x := 0; x < w/2; x++ {
			i := rowStart + (x * 4)
			j := rowStart + ((w - 1 - x) * 4)

			// Swap R, G, B, A
			p.img.Pix[i], p.img.Pix[j] = p.img.Pix[j], p.img.Pix[i]
			p.img.Pix[i+1], p.img.Pix[j+1] = p.img.Pix[j+1], p.img.Pix[i+1]
			p.img.Pix[i+2], p.img.Pix[j+2] = p.img.Pix[j+2], p.img.Pix[i+2]
			p.img.Pix[i+3], p.img.Pix[j+3] = p.img.Pix[j+3], p.img.Pix[i+3]
		}
	}
	return p
}

// ---------------------------------------------------------------------
// Advanced Filters (Convolution & Blending)
// ---------------------------------------------------------------------

// GaussianBlur applies a 5x5 Gaussian Blur.
// NOTE: This operation is computationally expensive.
func (p *Processor) GaussianBlur() *Processor {
	bounds := p.img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(bounds)
	copy(dst.Pix, p.img.Pix) // Copy original to preserve Alpha and borders

	// Approximate 5x5 Gaussian Kernel (Sum = 256)
	kernel := [5][5]int{
		{1, 4, 6, 4, 1},
		{4, 16, 24, 16, 4},
		{6, 24, 36, 24, 6},
		{4, 16, 24, 16, 4},
		{1, 4, 6, 4, 1},
	}

	// Skip 2 pixels border to avoid boundary checks inside the loop
	for y := 2; y < h-2; y++ {
		for x := 2; x < w-2; x++ {
			var rSum, gSum, bSum int

			// Convolution
			for ky := -2; ky <= 2; ky++ {
				for kx := -2; kx <= 2; kx++ {
					offset := p.img.PixOffset(x+kx, y+ky)
					kVal := kernel[ky+2][kx+2]

					rSum += int(p.img.Pix[offset]) * kVal
					gSum += int(p.img.Pix[offset+1]) * kVal
					bSum += int(p.img.Pix[offset+2]) * kVal
				}
			}

			dstOffset := dst.PixOffset(x, y)
			// Bitwise shift >> 8 is equivalent to dividing by 256
			dst.Pix[dstOffset] = uint8(rSum >> 8)
			dst.Pix[dstOffset+1] = uint8(gSum >> 8)
			dst.Pix[dstOffset+2] = uint8(bSum >> 8)
		}
	}

	p.img = dst
	return p
}

// SobelEdgeDetection applies the Sobel operator to detect edges.
// It converts the image to grayscale internally first.
func (p *Processor) SobelEdgeDetection() *Processor {
	// Edge detection works best on grayscale
	p.ToGray()

	bounds := p.img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(bounds)

	// Sobel Kernels
	gxKernel := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	gyKernel := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	// Skip 1 pixel border
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			var sumX, sumY int

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					offset := p.img.PixOffset(x+kx, y+ky)
					val := int(p.img.Pix[offset]) // Use R channel (already gray)

					sumX += val * gxKernel[ky+1][kx+1]
					sumY += val * gyKernel[ky+1][kx+1]
				}
			}

			// Magnitude = sqrt(Gx^2 + Gy^2)
			magnitude := math.Sqrt(float64(sumX*sumX + sumY*sumY))
			finalVal := clamp(int(magnitude))

			dstOffset := dst.PixOffset(x, y)
			dst.Pix[dstOffset] = finalVal   // R
			dst.Pix[dstOffset+1] = finalVal // G
			dst.Pix[dstOffset+2] = finalVal // B
			dst.Pix[dstOffset+3] = 255      // A (Opaque)
		}
	}

	p.img = dst
	return p
}

// AddWatermark overlays another image onto the current image.
// opacity: 0.0 (transparent) to 1.0 (opaque).
func (p *Processor) AddWatermark(watermark image.Image, startX, startY int, opacity float64) *Processor {
	// Ensure watermark is RGBA
	wmRGBA, ok := watermark.(*image.RGBA)
	if !ok {
		b := watermark.Bounds()
		wmRGBA = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(wmRGBA, wmRGBA.Bounds(), watermark, b.Min, draw.Src)
	}

	wmBounds := wmRGBA.Bounds()
	wmW, wmH := wmBounds.Dx(), wmBounds.Dy()
	srcW, srcH := p.img.Bounds().Dx(), p.img.Bounds().Dy()

	// Clamp opacity
	if opacity < 0 {
		opacity = 0
	}
	if opacity > 1 {
		opacity = 1
	}
	invOpacity := 1.0 - opacity

	for wy := 0; wy < wmH; wy++ {
		for wx := 0; wx < wmW; wx++ {
			srcX := startX + wx
			srcY := startY + wy

			// Boundary check
			if srcX < 0 || srcX >= srcW || srcY < 0 || srcY >= srcH {
				continue
			}

			srcOffset := p.img.PixOffset(srcX, srcY)
			wmOffset := wmRGBA.PixOffset(wx, wy)

			// Simple Alpha Blending: Result = Src * (1-opacity) + Watermark * opacity
			r := float64(p.img.Pix[srcOffset])*invOpacity + float64(wmRGBA.Pix[wmOffset])*opacity
			g := float64(p.img.Pix[srcOffset+1])*invOpacity + float64(wmRGBA.Pix[wmOffset+1])*opacity
			b := float64(p.img.Pix[srcOffset+2])*invOpacity + float64(wmRGBA.Pix[wmOffset+2])*opacity

			p.img.Pix[srcOffset] = uint8(r)
			p.img.Pix[srcOffset+1] = uint8(g)
			p.img.Pix[srcOffset+2] = uint8(b)
		}
	}
	return p
}

// ---------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------

// clamp restricts an integer to the range [0, 255].
func clamp(v int) uint8 {
	if v > 255 {
		return 255
	}
	if v < 0 {
		return 0
	}
	return uint8(v)
}
