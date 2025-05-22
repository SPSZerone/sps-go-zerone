package main

import (
	spsglfw "github.com/SPSZerone/sps-go-zerone/graphics/glfw"
	spsglfwin "github.com/SPSZerone/sps-go-zerone/graphics/glfw/window"
)

const (
	Name = "SPS GLFW Sample"
)

func main() {
	spsglfw.Run(
		func() spsglfw.Window {
			return spsglfwin.NewWindow(
				spsglfw.OptTitle(Name),
			)
		},
		func(win spsglfw.Window) {

		},
	)
}
