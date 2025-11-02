package mgl32

import "github.com/go-gl/mathgl/mgl32"

const FovDegreeForRealisticView = 45

func NewPerspectiveFovDegWH(fovAngle, width, height, near, far float32) mgl32.Mat4 {
	aspect := width / height
	return mgl32.Perspective(mgl32.DegToRad(fovAngle), aspect, near, far)
}

func NewPerspectiveFovRadWH(fovAngle, width, height, near, far float32) mgl32.Mat4 {
	aspect := width / height
	return mgl32.Perspective(fovAngle, aspect, near, far)
}
