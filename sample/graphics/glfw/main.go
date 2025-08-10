package main

import (
	spsgioglfw "github.com/SPSZerone/sps-go-zerone/graphics/gio/glfw"

	spsglfw "github.com/SPSZerone/sps-go-zerone/graphics/glfw"
	spsglfwwin "github.com/SPSZerone/sps-go-zerone/graphics/glfw/window"
)

const (
	Name = "SPS GLFW & GIO Sample"
)

func main() {
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
			gtx, size := ctx.LayoutContext(win.GetWindow())
			spsglfw.Clear(0, 0, 0, 1)
			ctx.Frame(gtx, size)
		}),
	)
}
