package mgl64

import "github.com/go-gl/mathgl/mgl64"

func NewMat4IdentityMatrix() mgl64.Mat4 {
	return mgl64.Ident4()
}

func NewMat4Translate3DVec3(v mgl64.Vec3) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.Translate3D(v.X(), v.Y(), v.Z()))
}

func NewMat4Translate3DXYZ(x, y, z float64) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.Translate3D(x, y, z))
}

func NewMat4HomogeneousRotate3DDegree(angle float64, axis mgl64.Vec3) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3D(mgl64.DegToRad(angle), axis))
}

func NewMat4HomogeneousRotate3DRadian(angle float64, axis mgl64.Vec3) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3D(angle, axis))
}

func NewMat4HomogeneousRotate3DXDegree(angle float64) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3DX(mgl64.DegToRad(angle)))
}

func NewMat4HomogeneousRotate3DXRadian(angle float64) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3DX(angle))
}

func NewMat4HomogeneousRotate3DYDegree(angle float64) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3DY(mgl64.DegToRad(angle)))
}

func NewMat4HomogeneousRotate3DYRadian(angle float64) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3DY(angle))
}

func NewMat4HomogeneousRotate3DZDegree(angle float64) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3DZ(mgl64.DegToRad(angle)))
}

func NewMat4HomogeneousRotate3DZRadian(angle float64) mgl64.Mat4 {
	result := mgl64.Ident4()
	return result.Mul4(mgl64.HomogRotate3DZ(angle))
}

func Mat4HomogeneousRotate3DDegree(mat4 mgl64.Mat4, angle float64, axis mgl64.Vec3) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3D(mgl64.DegToRad(angle), axis))
}

func Mat4HomogeneousRotate3DRadian(mat4 mgl64.Mat4, angle float64, axis mgl64.Vec3) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3D(angle, axis))
}

func Mat4HomogeneousRotate3DXDegree(mat4 mgl64.Mat4, angle float64) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3DX(mgl64.DegToRad(angle)))
}

func Mat4HomogeneousRotate3DXRadian(mat4 mgl64.Mat4, angle float64) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3DX(angle))
}

func Mat4HomogeneousRotate3DYDegree(mat4 mgl64.Mat4, angle float64) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3DY(mgl64.DegToRad(angle)))
}

func Mat4HomogeneousRotate3DYRadian(mat4 mgl64.Mat4, angle float64) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3DY(angle))
}

func Mat4HomogeneousRotate3DZDegree(mat4 mgl64.Mat4, angle float64) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3DZ(mgl64.DegToRad(angle)))
}

func Mat4HomogeneousRotate3DZRadian(mat4 mgl64.Mat4, angle float64) mgl64.Mat4 {
	return mat4.Mul4(mgl64.HomogRotate3DZ(angle))
}
