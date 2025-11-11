package image

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func LoadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error closing file %s err: %+v\n", filePath, err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	if _, ok := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}); ok {
		return img, nil
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba, nil
}

func SaveImage(filePath string, img image.Image) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error closing file %s err: %+v\n", filePath, err)
		}
	}()

	return png.Encode(file, img)
}

func Rotate90CCW(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, h, w))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dst.Set(y, w-1-x, src.At(x+bounds.Min.X, y+bounds.Min.Y))
		}
	}
	return dst
}
