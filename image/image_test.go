package image

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestProcessor(t *testing.T) {
	// 1. Load an image (Ensure you have 'input.png' in the directory)
	srcImg, err := LoadImage("temp/input.png")
	if err != nil {
		t.Logf("Error loading image('input.png') [%v]. using dummy image instead ...", err)
		// Create a dummy image for demonstration if file not found
		srcImg = createDummyImage(512, 512)
	}

	t.Log("Processing images...")

	var processor Processor

	// Example 1: Basic Filters + Resize
	processor.Init(srcImg, true).ToGray().AdjustContrast(1.2).Resize(256, 256)
	SaveImage("temp/output_basic.png", processor.GetImage())

	// Example 2: Flip + Invert
	processor.Init(srcImg, true).Flip(FlipVertically).Invert()
	SaveImage("temp/output_flip_invert.png", processor.GetImage())

	// Example 3: Edge Detection
	processor.Init(srcImg, true).SobelEdgeDetection()
	SaveImage("temp/output_sobel.png", processor.GetImage())

	// Example 4: Blur + Watermark
	// Create a simple red watermark
	wm, err := LoadImage("temp/logo.png")
	if err != nil {
		wm = image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(wm.(*image.RGBA), wm.Bounds(), &image.Uniform{C: color.RGBA{R: 255, A: 255}}, image.Point{}, draw.Src)
	}
	processor.Init(srcImg, true).GaussianBlur().AddWatermark(wm, 10, 10, 0.7)
	SaveImage("temp/output_blur_watermark.png", processor.GetImage())

	t.Log("Done! Check output files.")
}

func createDummyImage(w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	// Draw a gradient
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x % 255), G: uint8(y % 255), B: 100, A: 255})
		}
	}
	return img
}
