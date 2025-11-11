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
