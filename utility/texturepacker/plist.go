package texturepacker

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"

	"howett.net/plist"
)

func ReadPlist(fileName string) (result Plist, err error) {
	file, errOpen := os.Open(fileName)
	if errOpen != nil {
		err = fmt.Errorf("error opening file: %v", errOpen)
		return
	}
	defer func() {
		errClose := file.Close()
		if errClose != nil {
			fmt.Printf("error closing file %s err: %+v\n", fileName, errClose)
		}
	}()

	decoder := plist.NewDecoder(file)
	err = decoder.Decode(&result)
	return
}

// ParseFrameRectString parse '{{x,y},{w,h}}' to FrameRect
func ParseFrameRectString(frameStr string) (frameRect FrameRect, err error) {
	re := regexp.MustCompile(`[0-9]+`)
	matches := re.FindAllString(frameStr, -1)
	const dataCount = 4

	if len(matches) != dataCount {
		err = fmt.Errorf("frame data invalid: %s. should be '{{x,y},{w,h}}'", frameStr)
		return
	}

	nums := make([]int, dataCount)
	for i, match := range matches {
		n, errConvert := strconv.Atoi(match)
		if errConvert != nil {
			err = fmt.Errorf("convert str '%s' to num err: %v", match, errConvert)
			return
		}
		nums[i] = n
	}

	frameRect.X = nums[0]
	frameRect.Y = nums[1]
	frameRect.W = nums[2]
	frameRect.H = nums[3]
	return
}

type Plist struct {
	Metadata Metadata             `plist:"metadata"`
	Frames   map[string]FrameInfo `plist:"frames"`
}

type Metadata struct {
	Format              int    `plist:"format"`
	PixelFormat         string `plist:"pixelFormat"`
	PremultiplyAlpha    bool   `plist:"premultiplyAlpha"`
	RealTextureFileName string `plist:"realTextureFileName"`
	Size                string `plist:"size"`
	TextureFileName     string `plist:"textureFileName"`
}

type FrameInfo struct {
	Frame           string `plist:"frame"`
	Offset          string `plist:"offset"`
	Rotated         bool   `plist:"rotated"`
	SourceColorRect string `plist:"sourceColorRect"`
	SourceSize      string `plist:"sourceSize"`
}

type FrameRect struct {
	X int
	Y int
	W int
	H int
}

func (r *FrameRect) ToImageRect(rotate90 bool) image.Rectangle {
	if rotate90 {
		return image.Rect(r.X, r.Y, r.X+r.H, r.Y+r.W)
	}
	return image.Rect(r.X, r.Y, r.X+r.W, r.Y+r.H)
}
