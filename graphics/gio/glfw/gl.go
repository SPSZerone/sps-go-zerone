package glfw

import (
	"github.com/go-gl/gl/v3.1/gles2"
	"github.com/go-gl/gl/v3.3-core/gl"

	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

func Clear(red, green, blue, alpha float32) {
	desktopGL := spsos.IsDarwin()
	if desktopGL {
		gl.ClearColor(red, green, blue, alpha)
		gl.Clear(gl.COLOR_BUFFER_BIT)
	} else {
		gles2.ClearColor(red, green, blue, alpha)
		gles2.Clear(gl.COLOR_BUFFER_BIT)
	}
}
