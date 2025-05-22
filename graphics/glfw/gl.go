package glfw

import (
	"runtime"

	"github.com/go-gl/gl/v3.1/gles2"
	"github.com/go-gl/gl/v3.3-core/gl"

	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

func require() {
	// Required by the OpenGL threading model.
	runtime.LockOSThread()
}

func glInit() {
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
