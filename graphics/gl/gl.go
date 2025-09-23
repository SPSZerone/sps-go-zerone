package gl

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v3.1/gles2"
	"github.com/go-gl/gl/v3.3-core/gl"

	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

var (
	Logger = spslog.NewLogger()
)

func Require() {
	// Required by the OpenGL threading model.
	runtime.LockOSThread()
}

func Init() {
	desktopGL := spsos.IsDarwin()

	var err error
	if desktopGL {
		err = gl.Init()
	} else {
		err = gles2.Init()
	}
	if err != nil {
		Logger.Fatal().Msgf("gl.Init failed: %v", err)
	}

	if desktopGL {
		// Enable sRGB.
		gl.Enable(gl.FRAMEBUFFER_SRGB)
		// Set up default VBA, required for the forward-compatible core profile.
		var defVBA uint32
		gl.GenVertexArrays(1, &defVBA)
		gl.BindVertexArray(defVBA)
	}
}

func Clear(red, green, blue, alpha float32, desktopGL bool) {
	if desktopGL {
		gl.ClearColor(red, green, blue, alpha)
		gl.Clear(gl.COLOR_BUFFER_BIT)
	} else {
		gles2.ClearColor(red, green, blue, alpha)
		gles2.Clear(gl.COLOR_BUFFER_BIT)
	}
}

func Viewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func CreateProgram(vertexShaderSource, fragmentShaderSource string) (shaderProgram uint32, err error) {
	// vertex shader ==================================================
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	defer gl.DeleteShader(vertexShader)

	vertexShaderSourceCStrs, free := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, vertexShaderSourceCStrs, nil)
	free()
	gl.CompileShader(vertexShader)

	// check for shader compile errors
	var status int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vertexShader, logLength, nil, gl.Str(log))

		err = fmt.Errorf("CreateProgram: failed to compile vertex shader: %v", log)
		return
	}

	// fragment shader ==================================================
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	defer gl.DeleteShader(fragmentShader)

	fragmentShaderSourceCStrs, free := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, fragmentShaderSourceCStrs, nil)
	free()
	gl.CompileShader(fragmentShader)

	// check for shader compile errors
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fragmentShader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fragmentShader, logLength, nil, gl.Str(log))

		err = fmt.Errorf("CreateProgram: failed to compile fragment shader: %v", log)
		return
	}

	// link shaders ==================================================
	shaderProgram = gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)

	gl.LinkProgram(shaderProgram)
	// check for linking errors
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))

		err = fmt.Errorf("CreateProgram: failed to link shader program: %v", log)
		return
	}

	return
}
