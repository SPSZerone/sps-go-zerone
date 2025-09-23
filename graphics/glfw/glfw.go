package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	spsgl "github.com/SPSZerone/sps-go-zerone/graphics/gl"
	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

func InitWindowHint() {
	glfw.WindowHint(glfw.SRGBCapable, glfw.True)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)
	glfw.WindowHint(glfw.CocoaRetinaFramebuffer, glfw.True)

	desktopGL := spsos.IsDarwin()
	if desktopGL {
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	} else {
		glfw.WindowHint(glfw.ContextCreationAPI, glfw.EGLContextAPI)
		glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 0)
	}
}

func ProcessInputDefaultKeyEscape(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

func FramebufferSizeCallback(win *glfw.Window, width int, height int) {
	spsgl.Viewport(0, 0, int32(width), int32(height))
}
