package gl

import (
	"fmt"
	"runtime"

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

		// Enable depth test
		gl.Enable(gl.DEPTH_TEST)
	}
}

func ClearDefault() {
	Clear(0.2, 0.3, 0.3, 1, gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT, true)
}

func Clear(red, green, blue, alpha float32, mask uint32, desktopGL bool) {
	if desktopGL {
		gl.ClearColor(red, green, blue, alpha)
		gl.Clear(mask)
	} else {
		gles2.ClearColor(red, green, blue, alpha)
		gles2.Clear(mask)
	}
}

func Viewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func Str(name string) *uint8 {
	return gl.Str(fmt.Sprintf("%s\x00", name))
}
