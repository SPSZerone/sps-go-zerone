package mgl32

import "github.com/go-gl/mathgl/mgl32"

func Translate3D(x, y, z float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	result = result.Mul4(mgl32.Translate3D(x, y, z))
	return result
}
