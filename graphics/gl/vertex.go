package gl

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type SimpleVertex[T float32 | float64] struct {
	VAO uint32
	VBO uint32
	EBO uint32

	Vertices  []T
	AttrSizes []int32 // e.g. -> { 3 /* position */, 3 /* color */, 2 /* texture */ }
	Indices   []uint32

	vertexCount int32
}

func (v *SimpleVertex[T]) VertexDataInfo() (typeSize int32, glType uint32) {
	var zero T
	switch any(zero).(type) {
	case float32:
		typeSize = 4
		glType = gl.FLOAT
	case float64:
		typeSize = 8
		glType = gl.DOUBLE
	}
	return
}

// SetUpAndConfigure
// set up vertex data (and buffer(s)) and configure vertex attributes
func (v *SimpleVertex[T]) SetUpAndConfigure(usage uint32, wireframeMode bool, configure func() error) (err error) {
	if usage == 0 {
		usage = gl.STATIC_DRAW
	}

	verticesLen := len(v.Vertices)
	if verticesLen == 0 {
		err = fmt.Errorf("vertices is empty")
		return
	}

	vertexTypeSize, vertexGLType := v.VertexDataInfo()

	indicesLen := len(v.Indices)
	eboEnable := indicesLen > 0

	// = Generate Buffers/Arrays
	// ----------------------------------------------------------------------------------------------------
	gl.GenVertexArrays(1, &v.VAO)
	gl.GenBuffers(1, &v.VBO)
	if eboEnable {
		gl.GenBuffers(1, &v.EBO)
	}

	// = 1. bind the Vertex Array Object first
	// ----------------------------------------------------------------------------------------------------
	gl.BindVertexArray(v.VAO)

	// = 2. then bind and set vertex buffer(s)
	// ----------------------------------------------------------------------------------------------------
	gl.BindBuffer(gl.ARRAY_BUFFER, v.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, verticesLen*int(vertexTypeSize), gl.Ptr(v.Vertices), usage)

	if eboEnable {
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.EBO)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indicesLen*4 /* 4: bytes of uint32 */, gl.Ptr(v.Indices), usage)
	}

	// = 3. and then configure vertex attributes(s)
	// ----------------------------------------------------------------------------------------------------
	const positionSize = 3
	if configure != nil {
		err = configure()
		if err != nil {
			return
		}
	} else if eboEnable && indicesLen%positionSize == 0 { // default: Config as Position{XYZ(float32)} Attributes
		// position attribute
		index := uint32(0)

		// Linking Vertex Attributes
		// The vertex shader allows us to specify any input we want in the form of vertex attributes and while this allows for great flexibility,
		// it does mean we have to manually specify what part of our input data goes to which vertex attribute in the vertex shader.
		// This means we have to specify how OpenGL should interpret the vertex data before rendering.
		// With this knowledge we can tell OpenGL how it should interpret the vertex data (per vertex attribute) using gl.VertexAttribPointerWithOffset
		gl.VertexAttribPointerWithOffset(index,
			positionSize, vertexGLType, false,
			positionSize*vertexTypeSize, 0)

		// Now that we specified how OpenGL should interpret the vertex data we should also
		// enable the vertex attribute with gl.EnableVertexAttribArray giving the vertex attribute location as its argument;
		// vertex attributes are disabled by default.
		gl.EnableVertexAttribArray(index)
	}

	// = Unbind
	// ----------------------------------------------------------------------------------------------------
	// UnbindVBO
	// note that this is allowed,
	// the call to gl.VertexAttribPointerWithOffset
	// registered VBO as the vertex attribute's bound vertex buffer object
	// so afterwards we can safely unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	// UnbindEBO
	// remember: do NOT unbind the EBO while a VAO is active as the bound element buffer object IS stored in the VAO;
	// keep the EBO bound.
	//gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

	// UnbindVAO
	// You can unbind the VAO afterwards so other VAO calls won'v accidentally modify this VAO, but this rarely happens.
	// Modifying other VAOs requires a call to gl.BindVertexArray anyways
	// so we generally don'v unbind VAOs (nor VBOs) when it's not directly necessary.
	gl.BindVertexArray(0)

	// = wireframe mode
	// ----------------------------------------------------------------------------------------------------
	if wireframeMode {
		v.WireframePolygons()
	}

	return
}

func (v *SimpleVertex[T]) SetUpAttrib(normalized bool) {
	vertexTypeSize, vertexGLType := v.VertexDataInfo()
	attrLen := len(v.AttrSizes)

	totalSize := int32(0)
	for i := 0; i < attrLen; i++ {
		totalSize += v.AttrSizes[i]
	}
	stride := totalSize * vertexTypeSize

	offset := uintptr(0)
	for index := uint32(0); index < uint32(attrLen); index++ {
		attrSize := v.AttrSizes[index]

		gl.VertexAttribPointerWithOffset(index, attrSize, vertexGLType, normalized, stride, offset)
		gl.EnableVertexAttribArray(index)

		offset += uintptr(attrSize * vertexTypeSize)
	}
}

// WireframePolygons
// draw in wireframe polygons.
func (v *SimpleVertex[T]) WireframePolygons() {
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE /* default: gl.FILL */)
}

func (v *SimpleVertex[T]) BindVAO() {
	gl.BindVertexArray(v.VAO)
}

func (v *SimpleVertex[T]) UnBindVAO() {
	gl.BindVertexArray(0)
}

func (v *SimpleVertex[T]) DrawAsTriangles() {
	// seeing as we only have a single VAO there's no need to bind it every time,
	// but we'll do so to keep things a bit more organized
	v.BindVAO()

	// Draw
	v.DrawTrianglesByElements()

	// no need to unbind it every time
	//v.UnBindVAO()
}

func (v *SimpleVertex[T]) DrawAsTrianglesByArrays(first, count int32) {
	v.BindVAO()
	v.DrawTrianglesByArrays(first, count)
}

func (v *SimpleVertex[T]) DrawTrianglesByElements() {
	gl.DrawElementsWithOffset(gl.TRIANGLES, int32(len(v.Indices)), gl.UNSIGNED_INT, 0)
}

func (v *SimpleVertex[T]) VertexCount() int32 {
	if v.vertexCount == 0 {
		var attrSize int32
		for _, attr := range v.AttrSizes {
			attrSize += attr
		}
		verticesLen := int32(len(v.Vertices))
		v.vertexCount = verticesLen / attrSize
	}
	return v.vertexCount
}

func (v *SimpleVertex[T]) DrawTrianglesByArraysAll() {
	v.DrawTrianglesByArrays(0, v.VertexCount())
}

func (v *SimpleVertex[T]) DrawTrianglesByArrays(first, count int32) {
	gl.DrawArrays(gl.TRIANGLES, first, count)
}

func (v *SimpleVertex[T]) Delete() {
	if v.VAO != 0 {
		gl.DeleteVertexArrays(1, &v.VAO)
		v.VAO = 0
	}
	if v.VBO != 0 {
		gl.DeleteBuffers(1, &v.VBO)
		v.VBO = 0
	}
	if v.EBO != 0 {
		gl.DeleteBuffers(1, &v.EBO)
		v.EBO = 0
	}
}
