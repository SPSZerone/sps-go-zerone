package gl

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"

	spsio "github.com/SPSZerone/sps-go-zerone/io"
)

type Shader interface {
	ProgramID() uint32
	CreateProgram() error
	UseProgram()
	DeleteProgram()

	Uniform1i(name string, value int32)
	Uniform1ui(name string, value uint32)
	Uniform1f(name string, value float32)
	Uniform1d(name string, value float64)

	Uniform2i(name string, v0, v1 int32)
	Uniform2ui(name string, v0, v1 uint32)
	Uniform2f(name string, v0, v1 float32)
	Uniform2d(name string, v0, v1 float64)

	Uniform3i(name string, v0, v1, v2 int32)
	Uniform3ui(name string, v0, v1, v2 uint32)
	Uniform3f(name string, v0, v1, v2 float32)
	Uniform3d(name string, v0, v1, v2 float64)

	Uniform4i(name string, v0, v1, v2, v3 int32)
	Uniform4ui(name string, v0, v1, v2, v3 uint32)
	Uniform4f(name string, v0, v1, v2, v3 float32)
	Uniform4d(name string, v0, v1, v2, v3 float64)
}

type SourceType byte

const (
	SourceTypeCode SourceType = iota
	SourceTypeFile
)

type SimpleShader struct {
	SourceType SourceType
	Vertex     string
	Fragment   string

	programID uint32
}

func (s *SimpleShader) ProgramID() uint32 {
	return s.programID
}

func (s *SimpleShader) CreateProgram() (err error) {
	switch s.SourceType {
	case SourceTypeCode:
		err = s.createProgramAsCode()
	case SourceTypeFile:
		err = s.createProgramAsFile()
	default:
		err = fmt.Errorf("SimpleShader unknown SourceType: %v", s.SourceType)
	}
	return
}

func (s *SimpleShader) createProgramAsCode() (err error) {
	if s.programID != 0 {
		return
	}
	s.programID, err = CreateProgram(s.Vertex, s.Fragment)
	return
}

func (s *SimpleShader) createProgramAsFile() (err error) {
	if s.programID != 0 {
		return
	}
	const bufSize = 0x800 // 2KB
	vertexSource, vertexErr := spsio.ReadFileBytes(s.Vertex, bufSize)
	if vertexErr != nil {
		return vertexErr
	}
	fragmentSource, fragmentErr := spsio.ReadFileBytes(s.Fragment, bufSize)
	if fragmentErr != nil {
		return fragmentErr
	}
	s.programID, err = CreateProgram(string(vertexSource), string(fragmentSource))
	return
}

func (s *SimpleShader) UseProgram() {
	gl.UseProgram(s.programID)
}

func (s *SimpleShader) DeleteProgram() {
	gl.DeleteProgram(s.programID)
}

func (s *SimpleShader) Uniform1i(name string, value int32) {
	gl.Uniform1i(GetUniformLocation(s.programID, name), value)
}

func (s *SimpleShader) Uniform1ui(name string, value uint32) {
	gl.Uniform1ui(GetUniformLocation(s.programID, name), value)
}

func (s *SimpleShader) Uniform1f(name string, value float32) {
	gl.Uniform1f(GetUniformLocation(s.programID, name), value)
}

func (s *SimpleShader) Uniform1d(name string, value float64) {
	gl.Uniform1d(GetUniformLocation(s.programID, name), value)
}

func (s *SimpleShader) Uniform2i(name string, v0, v1 int32) {
	gl.Uniform2i(GetUniformLocation(s.programID, name), v0, v1)
}

func (s *SimpleShader) Uniform2ui(name string, v0, v1 uint32) {
	gl.Uniform2ui(GetUniformLocation(s.programID, name), v0, v1)
}

func (s *SimpleShader) Uniform2f(name string, v0, v1 float32) {
	gl.Uniform2f(GetUniformLocation(s.programID, name), v0, v1)
}

func (s *SimpleShader) Uniform2d(name string, v0, v1 float64) {
	gl.Uniform2d(GetUniformLocation(s.programID, name), v0, v1)
}

func (s *SimpleShader) Uniform3i(name string, v0, v1, v2 int32) {
	gl.Uniform3i(GetUniformLocation(s.programID, name), v0, v1, v2)
}

func (s *SimpleShader) Uniform3ui(name string, v0, v1, v2 uint32) {
	gl.Uniform3ui(GetUniformLocation(s.programID, name), v0, v1, v2)
}

func (s *SimpleShader) Uniform3f(name string, v0, v1, v2 float32) {
	gl.Uniform3f(GetUniformLocation(s.programID, name), v0, v1, v2)
}

func (s *SimpleShader) Uniform3d(name string, v0, v1, v2 float64) {
	gl.Uniform3d(GetUniformLocation(s.programID, name), v0, v1, v2)
}

func (s *SimpleShader) Uniform4i(name string, v0, v1, v2, v3 int32) {
	gl.Uniform4i(GetUniformLocation(s.programID, name), v0, v1, v2, v3)
}

func (s *SimpleShader) Uniform4ui(name string, v0, v1, v2, v3 uint32) {
	gl.Uniform4ui(GetUniformLocation(s.programID, name), v0, v1, v2, v3)
}

func (s *SimpleShader) Uniform4f(name string, v0, v1, v2, v3 float32) {
	gl.Uniform4f(GetUniformLocation(s.programID, name), v0, v1, v2, v3)
}

func (s *SimpleShader) Uniform4d(name string, v0, v1, v2, v3 float64) {
	gl.Uniform4d(GetUniformLocation(s.programID, name), v0, v1, v2, v3)
}

func CreateProgram(vertexShaderSource, fragmentShaderSource string) (shaderProgram uint32, err error) {
	var vertexShader, fragmentShader uint32

	// = vertex shader
	// ----------------------------------------------------------------------------------------------------
	vertexShader, err = NewShader(gl.VERTEX_SHADER, vertexShaderSource)
	defer gl.DeleteShader(vertexShader)
	if err != nil {
		return
	}

	// = fragment shader
	// ----------------------------------------------------------------------------------------------------
	fragmentShader, err = NewShader(gl.FRAGMENT_SHADER, fragmentShaderSource)
	defer gl.DeleteShader(fragmentShader)
	if err != nil {
		return
	}

	// = link shaders
	// ----------------------------------------------------------------------------------------------------
	shaderProgram, err = NewProgram(vertexShader, fragmentShader)

	return
}

func NewShader(shaderType uint32, source string) (shader uint32, err error) {
	// create shader
	shader = gl.CreateShader(shaderType)

	// shader source
	sourceCStrs, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, sourceCStrs, nil)
	free()

	// compile shader
	gl.CompileShader(shader)

	// check for shader compile errors
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		err = fmt.Errorf("NewShader | failed to compile shader | %v | %v", shaderType, log)
	}

	return
}

func NewProgram(vertexShader, fragmentShader uint32) (shaderProgram uint32, err error) {
	// create program
	shaderProgram = gl.CreateProgram()

	// attach shader
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)

	// link program
	gl.LinkProgram(shaderProgram)

	// check for linking errors
	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))

		err = fmt.Errorf("NewProgram | failed to link shader program | %v", log)
	}

	return
}

func GetUniformLocation(programID uint32, name string) int32 {
	return gl.GetUniformLocation(programID, Str(name))
}
