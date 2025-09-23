package main

import (
	spsos "github.com/SPSZerone/sps-go-zerone/os"

	spsgl "github.com/SPSZerone/sps-go-zerone/graphics/gl"
	spsglfw "github.com/SPSZerone/sps-go-zerone/graphics/glfw"
	spsglfwwin "github.com/SPSZerone/sps-go-zerone/graphics/glfw/window"

	spsgioglfw "github.com/SPSZerone/sps-go-zerone/graphics/gio/glfw"
)

const (
	Name = "SPS GLFW & GIO Sample"
)

func main() {
	desktopGL := spsos.IsDarwin()
	var ctx spsgioglfw.Context

	spsglfw.Run(
		spsglfw.OptOnGLFWInit(func() {
		}),
		spsglfw.OptNewWindow(func() spsglfw.Window {
			return spsglfwwin.NewWindow(
				spsglfwwin.OptWidth(1920),
				spsglfwwin.OptHeight(1080),
				spsglfwwin.OptTitle(Name),
			)
		}),
		spsglfw.OptOnGLInit(func() {
			ctx = spsgioglfw.NewDefaultContext()
		}),
		spsglfw.OptOnStop(func() {
			ctx.Release()
		}),
		spsglfw.OptOnLoop(func(win spsglfw.Window) {
			spsglfw.ProcessInputDefaultKeyEscape(win.GetWindow())

			gtx, size := ctx.LayoutContext(win.GetWindow())
			spsgl.Clear(0, 0, 0, 1, desktopGL)

			ctx.Frame(gtx, size)
		}),
	)
}
