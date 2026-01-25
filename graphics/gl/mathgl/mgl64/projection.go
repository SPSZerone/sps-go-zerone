package mgl64

import "github.com/go-gl/mathgl/mgl64"

func NewPerspectiveFovDegWH(fovAngle, width, height, near, far float64) mgl64.Mat4 {
	aspect := width / height
	return mgl64.Perspective(mgl64.DegToRad(fovAngle), aspect, near, far)
}

func NewPerspectiveFovRadWH(fovAngle, width, height, near, far float64) mgl64.Mat4 {
	aspect := width / height
	return mgl64.Perspective(fovAngle, aspect, near, far)
}
