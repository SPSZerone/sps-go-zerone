package mgl32

import "github.com/go-gl/mathgl/mgl32"

func NewMat4IdentityMatrix() mgl32.Mat4 {
	return mgl32.Ident4()
}

func NewMat4Translate3DVec3(v mgl32.Vec3) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.Translate3D(v.X(), v.Y(), v.Z()))
}

func NewMat4Translate3DXYZ(x, y, z float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.Translate3D(x, y, z))
}

func NewMat4HomogeneousRotate3DDegree(angle float32, axis mgl32.Vec3) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(angle), axis))
}

func NewMat4HomogeneousRotate3DRadian(angle float32, axis mgl32.Vec3) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3D(angle, axis))
}

func NewMat4HomogeneousRotate3DXDegree(angle float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3DX(mgl32.DegToRad(angle)))
}

func NewMat4HomogeneousRotate3DXRadian(angle float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3DX(angle))
}

func NewMat4HomogeneousRotate3DYDegree(angle float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3DY(mgl32.DegToRad(angle)))
}

func NewMat4HomogeneousRotate3DYRadian(angle float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3DY(angle))
}

func NewMat4HomogeneousRotate3DZDegree(angle float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(angle)))
}

func NewMat4HomogeneousRotate3DZRadian(angle float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	return result.Mul4(mgl32.HomogRotate3DZ(angle))
}

func Mat4HomogeneousRotate3DDegree(mat4 mgl32.Mat4, angle float32, axis mgl32.Vec3) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(angle), axis))
}

func Mat4HomogeneousRotate3DRadian(mat4 mgl32.Mat4, angle float32, axis mgl32.Vec3) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3D(angle, axis))
}

func Mat4HomogeneousRotate3DXDegree(mat4 mgl32.Mat4, angle float32) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3DX(mgl32.DegToRad(angle)))
}

func Mat4HomogeneousRotate3DXRadian(mat4 mgl32.Mat4, angle float32) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3DX(angle))
}

func Mat4HomogeneousRotate3DYDegree(mat4 mgl32.Mat4, angle float32) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3DY(mgl32.DegToRad(angle)))
}

func Mat4HomogeneousRotate3DYRadian(mat4 mgl32.Mat4, angle float32) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3DY(angle))
}

func Mat4HomogeneousRotate3DZDegree(mat4 mgl32.Mat4, angle float32) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(angle)))
}

func Mat4HomogeneousRotate3DZRadian(mat4 mgl32.Mat4, angle float32) mgl32.Mat4 {
	return mat4.Mul4(mgl32.HomogRotate3DZ(angle))
}
