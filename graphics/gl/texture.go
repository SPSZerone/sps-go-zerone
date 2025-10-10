package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"

	spsimg "github.com/SPSZerone/sps-go-zerone/image"
)

func CreateTextureRGBA(file string, genMipmap bool, setParams func()) (texture uint32, err error) {
	rgba, errRGBA := spsimg.LoadAsRGBA(file, spsimg.OpFlipVertically)
	if errRGBA != nil {
		err = errRGBA
		return
	}

	size := rgba.Rect.Size()
	texture = NewTextureRGBA(rgba.Pix, int32(size.X), int32(size.Y), genMipmap, setParams)

	return
}

func NewTextureRGBA(pixels []uint8, width, height int32, genMipmap bool, setParams func()) (texture uint32) {
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	if setParams == nil {
		// set the texture wrapping parameters (default: gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_BORDER)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_BORDER)
		// set texture filtering parameters
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	} else {
		setParams()
	}

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

	if genMipmap {
		gl.GenerateMipmap(gl.TEXTURE_2D)
	}

	return
}
