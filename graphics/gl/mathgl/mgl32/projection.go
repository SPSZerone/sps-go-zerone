package mgl32

import "github.com/go-gl/mathgl/mgl32"

const FovAngleForRealisticView = 45

func PerspectiveFAWH(fovAngle, width, height, near, far float32) mgl32.Mat4 {
	fov := mgl32.DegToRad(fovAngle)
	aspect := width / height
	result := mgl32.Perspective(fov, aspect, near, far)
	return result
}
