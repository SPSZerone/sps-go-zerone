package main

import (
	spsgioglfw "github.com/SPSZerone/sps-go-zerone/graphics/gio/glfw"
	spsgl "github.com/SPSZerone/sps-go-zerone/graphics/gl"
	spsglfw "github.com/SPSZerone/sps-go-zerone/graphics/glfw"
	spsglfwwin "github.com/SPSZerone/sps-go-zerone/graphics/glfw/window"
)

const (
	Name = "SPS GLFW & GIO Sample"
)

func main() {
	var ctx spsgioglfw.Context

	spsglfw.Run(
		spsglfw.OptOnGLFWInit(func() (err error) {
			return
		}),
		spsglfw.OptNewWindow(func() spsglfw.Window {
			return spsglfwwin.NewWindow(
				spsglfwwin.OptWidth(1920),
				spsglfwwin.OptHeight(1080),
				spsglfwwin.OptTitle(Name),
			)
		}),
		spsglfw.OptOnGLInit(func() (err error) {
			ctx = spsgioglfw.NewDefaultContext()
			return
		}),
		spsglfw.OptOnStop(func() {
			ctx.Release()
		}),
		spsglfw.OptOnProcessInput(func(win spsglfw.Window) {
			spsglfw.ProcessInputDefaultKeyEscape(win.GetWindow())
		}),
		spsglfw.OptOnRender(func(win spsglfw.Window) {
			gtx, size := ctx.LayoutContext(win.GetWindow())
			spsgl.ClearDefault()

			ctx.Frame(gtx, size)
		}),
	)
}
