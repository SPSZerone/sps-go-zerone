package image

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func LoadAsRGBA(file string, op Op) (rgba *image.RGBA, err error) {
	// open file
	imgFile, errFile := os.Open(file)
	if errFile != nil {
		err = fmt.Errorf("file %v open failed %v", file, errFile)
		return
	}
	defer func() {
		errClose := imgFile.Close()
		if errClose != nil {
			fmt.Println(fmt.Sprintf("file %v close failed %v", file, errClose))
		}
	}()

	// decode image
	img, _, errDecode := image.Decode(imgFile)
	if errDecode != nil {
		err = fmt.Errorf("image decode err: %v", errDecode)
		return
	}

	imgBounds := img.Bounds()
	rgba = image.NewRGBA(imgBounds)
	rgbaSize := rgba.Rect.Size()
	if rgba.Stride != rgbaSize.X*4 {
		err = fmt.Errorf("unsupported stride")
		return
	}

	switch op {
	case OpFlipVertically:
		for x := imgBounds.Min.X; x < imgBounds.Max.X; x++ {
			for y := imgBounds.Min.Y; y < imgBounds.Max.Y; y++ {
				color := img.At(x, y)
				rgba.Set(x, imgBounds.Max.Y-y-1, color)
			}
		}
	case OpFlipHorizontally:
		for x := imgBounds.Min.X; x < imgBounds.Max.X; x++ {
			for y := imgBounds.Min.Y; y < imgBounds.Max.Y; y++ {
				color := img.At(x, y)
				rgba.Set(imgBounds.Max.X-x-1, y, color)
			}
		}
	case OpFlipVerticallyHorizontally:
		for x := imgBounds.Min.X; x < imgBounds.Max.X; x++ {
			for y := imgBounds.Min.Y; y < imgBounds.Max.Y; y++ {
				color := img.At(x, y)
				rgba.Set(imgBounds.Max.X-x-1, imgBounds.Max.Y-y-1, color)
			}
		}
	case OpOriginally:
		fallthrough
	default:
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
	}

	return
}
