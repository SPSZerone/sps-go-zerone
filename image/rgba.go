package image

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func LoadWithFlipAsRGBA(file string, flipMode FlipMode) (rgba *image.RGBA, err error) {
	img, err := LoadImage(file)
	if err != nil {
		return
	}

	var processor Processor
	processor.Init(img, false).Flip(flipMode)
	rgba = processor.GetImage()

	return
}
